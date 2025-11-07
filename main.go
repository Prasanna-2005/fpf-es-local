package main

import (
	"fmt"
	"os"
	"os/signal"
	"quantrium-tech/form_pf_extractor_service/config"
	"quantrium-tech/form_pf_extractor_service/db"
	"quantrium-tech/form_pf_extractor_service/executors"
	privateinvestors "quantrium-tech/form_pf_extractor_service/executors/FEG/PrivateInvestors"
	"quantrium-tech/form_pf_extractor_service/ipc"
	"quantrium-tech/form_pf_extractor_service/logging"
	"reflect"
	"syscall"
)

func init() {

	logging.InitialiseLogger()
	logging.Info("Logger Initialised")

	config.InitialiseConfig()
	logging.Info("Config Initialised")

	err := db.InitialseDB()
	if err != nil {
		panic(err)
	}
	logging.Info("Database Connection Initialised")

	err = ipc.InitialiseDaemon()
	if err != nil {
		panic(err)
	}
	logging.Info("IPC Daemon Initialised")

	err = db.SetFundInfo()
	if err != nil {
		panic(err)
	}
	logging.Info("Fund Information Initialised")

	// start the active listener
	go ipc.Listen()

}

// handleSignal waits/blocks the routine till expected signal arrives and proceed with the cleanup.
func handleSignal(c chan os.Signal) {
	<-c
	logging.Warn("Got a Kill/Interupt Signal")
	if err := ipc.TearDownSocket(); err != nil {
		logging.Error("Error in Tearing Down: " + err.Error() + " Sock File May Prevail. Manual Cleanup is Mandatory")
	}
	logging.Info("Teardown of Sock file is completed")
	os.Exit(1)

}

func main() {

	// Create the signal handler and register with a channel
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go handleSignal(c)

	// Main Forwarder/Handler from the Listener
	for {

		select {
		case m := <-ipc.InboundChannel:
			logging.Info("Recieved processing signal for the filing with ID: " + m.FilingID)
			go execute(m.FilingID, m.UserID)
		}

	}

}

// prepareFilingPrivateInvestors takes the filingID and prepares all the funds for the execution.
// It retrieves the mappings from the db and map them to internal fund structures
func prepareFilingPrivateInvestors(filingID string) (map[string][]*privateinvestors.PrivateFund, string, error) {

	fundsMapper := make(map[string][]*privateinvestors.PrivateFund)
	var investorLiquidityPath string

	mappings, err := db.PullFiling(filingID)
	if err != nil {
		return fundsMapper, investorLiquidityPath, fmt.Errorf("pulling filing details from DB: " + err.Error())
	}

	for _, mapping := range *mappings {

		if mapping.FundShortCode == "Q15_Q16" {
			investorLiquidityPath = os.Getenv("STORAGE_PATH") + mapping.FilingPath
			continue
		}

		converted := privateinvestors.ShortCodeConvertor[mapping.FundShortCode]

		if converted == "" {
			logging.Warn(mapping.FundShortCode + "->" + "Unknown Short Code Encountered. Conversion Failed")
			continue
		}

		mapping.FundShortCode = converted
		if _, ok := fundsMapper[mapping.FundShortCode]; !ok {
			fundsMapper[mapping.FundShortCode] = privateinvestors.CreateAllFundSeries(mapping.FundShortCode)
		}

		fundsList := fundsMapper[mapping.FundShortCode]

		for _, fund := range fundsList {
			fund.SetHolder(mapping.QuarterNo, mapping.FilingPath)
		}

	}

	return fundsMapper, investorLiquidityPath, nil
}

// setStaticDataPV sets the static information to the form1bData
func setStaticDataPV(data *executors.Form1bData, fundId string) {

	info := db.FundStorer[fundId]

	data.FundID = info.FundID
	data.FundName = info.FundName
	data.FundLEI = info.FundLEI.String
	data.FundNFA = info.FundNFA.String
}

// execute acts as the core execution and triggering point for both extraction and computation.
func execute(filingID string, userID string) {

	logging.Info("Execution Started for the filing with ID: " + filingID)
	mapper, investorLiquidityPath, err := prepareFilingPrivateInvestors(filingID)
	if err != nil {
		logging.Error("Error in Preparing for the Filing with Id: " + filingID + ": " + err.Error())
		return
	}

	if len(mapper) == 0 {
		logging.Error("No Funds Mapped. Execution cannot be Started")
		return
	}

	existing, err := db.GetExisitingFunds(filingID)
	if err != nil {
		logging.Error("Error in Getting the Existing Funds in Form Data: " + err.Error())
		return
	}

	liqudityHolder := make(privateinvestors.LiquidityDataHolder)

	if investorLiquidityPath != "" {
		liqudityHolder, err = privateinvestors.ExtractInvestorLiquidity(investorLiquidityPath)
		if err != nil {
			logging.Error("Error in Extracting Investor Liquidity: " + err.Error())
			return
		}
	}

	logging.Info("Extraction of Investor Liquidity Completed")

	logging.Info(fmt.Sprintf("Total number of funds mapped: %d", len(mapper)))
	for _, list := range mapper {
		for _, fund := range list {

			if _, ok := existing[fund.GetFundID()]; ok {
				logging.Info(fund.GetFundID() + " Already Populated in the Form Data. Skipping")
				continue
			}

			answersMapper := make(map[[2]string]string)

			if holder, ok := liqudityHolder[fund.GetFundID()]; ok {
				fund.SetLiquidityData(holder)
			} else {
				logging.Warn("Liquidity Data not available for the Fund: " + fund.GetFundID())
			}

			logging.Info("Extraction Started for the Fund: " + fund.GetFundID())
			err := fund.ExtractWorkPaperData()
			if err != nil {
				logging.Error("Error in extracting for the Fund: " + fund.GetFundID() + ": " + err.Error())
				continue
			}
			logging.Info("Extraction Completed for the Fund: " + fund.GetFundID())

			extractionData := fund.ExportWorkPaperData()
			err = db.PushEntities(filingID, fund.GetFundID(), userID, extractionData)
			if err != nil {
				logging.Error("Error in Pushing to Extraction Table: " + err.Error())
				continue
			}
			logging.Info("Entraction Entities are Pushed to the DB")

			logging.Info("Computation Started for the Fund: " + fund.GetFundID())
			formData := fund.Compute()
			logging.Info("Computation Completed for the Fund: " + fund.GetFundID())

			setStaticDataPV(formData, fund.GetFundID())

			executors.CollectAnswers(reflect.ValueOf(*formData), answersMapper)
			err = db.PushFormData(filingID, fund.GetFundID(), userID, answersMapper)
			if err != nil {
				logging.Error("Error in Pushing to Form Data Table: " + err.Error())
			}

			logging.Info("Form Data Entries are Pushed to the DB")

			logging.Info("Execution Completed for the Fund: " + fund.GetFundID())

		}

	}
	logging.Info("Execution Completed for the filing with ID: " + filingID)
}
