package main

import (
	"fmt"
	"os"
	"os/signal"
	"quantrium-tech/form_pf_extractor_service/config"
	"quantrium-tech/form_pf_extractor_service/db"
	privateinvestors "quantrium-tech/form_pf_extractor_service/executors/FEG/PrivateInvestors"
	"quantrium-tech/form_pf_extractor_service/ipc"
	"quantrium-tech/form_pf_extractor_service/logging"
	"strconv"
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
			go execute(m.FilingID)
		}

	}

}

// prepareFilingPrivateInvestors takes the filingID and prepares all the funds for the execution.
// It retrieves the mappings from the db and map them to internal fund structures
func prepareFilingPrivateInvestors(filingID string) (map[string][]*privateinvestors.PrivateFund, string, error) {

	fundsMapper := make(map[string][]*privateinvestors.PrivateFund)
	var investorLiquidityPath string

	Id, _ := strconv.Atoi(filingID)
	mappings, err := db.PullFiling(Id)
	if err != nil {
		return fundsMapper, investorLiquidityPath, fmt.Errorf("pulling filing details from DB: " + err.Error())
	}

	for _, mapping := range *mappings {

		if mapping.FundShortCode == "Q15_Q16" {
			investorLiquidityPath = mapping.FilingPath
			continue
		}

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

// execute acts as the core execution and triggering point for both extraction and computation.
func execute(filingID string) {

	logging.Info("Execution Started for the filing with ID: " + filingID)
	mapper, investorLiquidityPath, err := prepareFilingPrivateInvestors(filingID)
	if err != nil {
		logging.Error("Error in Preparing for the Filing with Id: " + filingID + ": " + err.Error())
		return
	}

	liqudityHolder, err := privateinvestors.ExtractInvestorLiquidity(investorLiquidityPath)
	if err != nil {
		logging.Error("Error in Extracting Investor Liquidity: " + err.Error())
		return
	}

	deferedHolder := *(liqudityHolder)
	for _, list := range mapper {
		for _, fund := range list {
			fund.SetLiquidityData(deferedHolder[fund.GetFundID()])
			logging.Info("Extraction Started for the Fund: " + fund.GetFundID())
			err := fund.ExtractWorkPaperData()
			if err != nil {
				logging.Error("Error in extracting for the Fund: " + fund.GetFundID() + ": " + err.Error())
				continue
			}
			logging.Info("Extraction Completed for the Fund: " + fund.GetFundID())

			logging.Info("Computation Started for the Fund: " + fund.GetFundID())
			fund.Compute()
			logging.Info("Computation Completed for the Fund: " + fund.GetFundID())
		}

	}

}
