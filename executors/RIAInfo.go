package executors

import (
	"strings"
)

type RegInvAdv struct {
	clientCode string
	series     string
	itemA      ItemA
	itemB      ItemB
}

type ItemA struct {
	legalName           string
	crdNumber           string
	sec801Number        string
	nfaId               string
	largeTraderId       string
	largeTraderIdSuffix string
	reportingType       string
	firmSign            RIAInfo
	firmDependents      RIADependents
}

type ItemB struct {
	Hedge          AUM
	Liquidity      AUM
	Private        AUM
	RealEstate     AUM
	SecuredAssets  AUM
	VentureCapital AUM
	Other          AUM
	NOTA           AUM
}

type AUM struct {
	NetAssest  float64
	GrossAsset float64
}

type RIADependents struct {
	dependentName string
	firmInfo      RIAInfo
}

type RIAInfo struct {
	name  string
	sign  string
	title string
	email string
	phno  string
}

func NewRegInvAdv(clientCode, series string) *RegInvAdv {
	return &RegInvAdv{clientCode: clientCode, series: series}
}

func (firm *RegInvAdv) NewfirmSign() RIAInfo {
	switch firm.clientCode {
	case "FEG":
		return RIAInfo{
			name:  "Julie T. Thomas",
			sign:  "Julie T. Thomas",
			title: "Chief Compliance Officer",
			email: "jthomas@feg.com",
			phno:  "5139774400",
		}
	default:
		return RIAInfo{}

	}

}

func (firm *RegInvAdv) CalcAUM(fundType string, nav float64, gav float64) {
	switch strings.ToLower(fundType) {
	case "hedge":
		firm.itemB.Hedge.NetAssest += nav
		firm.itemB.Hedge.GrossAsset += gav
	case "liquidity":
		firm.itemB.Liquidity.NetAssest += nav
		firm.itemB.Liquidity.GrossAsset += gav
	case "private":
		firm.itemB.Private.NetAssest += nav
		firm.itemB.Private.GrossAsset += gav
	case "realestate":
		firm.itemB.RealEstate.NetAssest += nav
		firm.itemB.RealEstate.GrossAsset += gav
	case "securedassets":
		firm.itemB.SecuredAssets.NetAssest += nav
		firm.itemB.SecuredAssets.GrossAsset += gav
	case "venturecapital":
		firm.itemB.VentureCapital.NetAssest += nav
		firm.itemB.VentureCapital.GrossAsset += gav
	case "other":
		firm.itemB.Other.NetAssest += nav
		firm.itemB.Other.GrossAsset += gav
	case "nota":
		firm.itemB.NOTA.NetAssest += nav
		firm.itemB.NOTA.GrossAsset += gav
	}
}

func (firm *RegInvAdv) Sec1aData() {
	switch firm.clientCode + ":" + firm.series {
	case "FEG:Investment Services":
		firm.itemA.legalName = "FEG INVESTMENT SERVICES, LLC"
		firm.itemA.crdNumber = "281516"
		firm.itemA.sec801Number = "801-108644"
		firm.itemA.nfaId = "0518605"
		firm.itemA.largeTraderId = "NA"
		firm.itemA.largeTraderIdSuffix = "NA"
		firm.itemA.firmSign = firm.NewfirmSign()

	case "FEG:Curio":
		firm.itemA.legalName = "FEG CURIO INVESTMENT PARTNERS, LLC"
		firm.itemA.crdNumber = "323497"
		firm.itemA.sec801Number = "801-126960"
		firm.itemA.nfaId = ""
		firm.itemA.largeTraderId = "NA"
		firm.itemA.largeTraderIdSuffix = "NA"
		firm.itemA.firmSign = firm.NewfirmSign()

	case "FEG:Private Investors":
		firm.itemA.legalName = "FEG PRIVATE INVESTORS, LLC"
		firm.itemA.crdNumber = "156458"
		firm.itemA.sec801Number = "801-108644"
		firm.itemA.nfaId = "0469730"
		firm.itemA.largeTraderId = "NA"
		firm.itemA.largeTraderIdSuffix = "NA"
		firm.itemA.firmSign = firm.NewfirmSign()
	}

}

func (firm *RegInvAdv) Compute() *Form1aData {
	form1a := NewForm1aData()

	// Populate ItemA (Firm Info)
	form1a.LegalName = firm.itemA.legalName
	form1a.CrdNumber = firm.itemA.crdNumber
	form1a.Sec801Number = firm.itemA.sec801Number
	form1a.NfaId = firm.itemA.nfaId
	form1a.LargeTraderId = firm.itemA.largeTraderId
	form1a.LargeTraderIdSuffix = firm.itemA.largeTraderIdSuffix
	form1a.FirmIndividualName = firm.itemA.firmSign.name
	form1a.FirmIndividualSignature = firm.itemA.firmSign.sign
	form1a.FirmIndividualTitle = firm.itemA.firmSign.title
	form1a.FirmIndividualEmail = firm.itemA.firmSign.email
	form1a.FirmIndividualPhno = firm.itemA.firmSign.phno

	// Populate ItemB (AUM Data) - Hedge
	form1a.HedgeFundsGAV = firm.itemB.Hedge.GrossAsset
	form1a.HedgeFundsNAV = firm.itemB.Hedge.NetAssest

	// Populate ItemB (AUM Data) - Liquidity
	form1a.LiquidityGAV = firm.itemB.Liquidity.GrossAsset
	form1a.LiquidityNAV = firm.itemB.Liquidity.NetAssest

	// Populate ItemB (AUM Data) - Private
	form1a.PrivateGAV = firm.itemB.Private.GrossAsset
	form1a.PrivateNAV = firm.itemB.Private.NetAssest

	// Populate ItemB (AUM Data) - Real Estate
	form1a.RealEstateGAV = firm.itemB.RealEstate.GrossAsset
	form1a.RealEstateNAV = firm.itemB.RealEstate.NetAssest

	// Populate ItemB (AUM Data) - Secured Assets
	form1a.SecuredAssetsGAV = firm.itemB.SecuredAssets.GrossAsset
	form1a.SecuredAssetsNAV = firm.itemB.SecuredAssets.NetAssest

	// Populate ItemB (AUM Data) - Venture Capital
	form1a.VentureCapitalGAV = firm.itemB.VentureCapital.GrossAsset
	form1a.VentureCapitalNAV = firm.itemB.VentureCapital.NetAssest

	// Populate ItemB (AUM Data) - Other
	form1a.OtherGAV = firm.itemB.Other.GrossAsset
	form1a.OtherNAV = firm.itemB.Other.NetAssest

	// Populate ItemB (AUM Data) - NOTA
	form1a.NOTAGAV = firm.itemB.NOTA.GrossAsset
	form1a.NOTANAV = firm.itemB.NOTA.NetAssest

	return form1a
}
