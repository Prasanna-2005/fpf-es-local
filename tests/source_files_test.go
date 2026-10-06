package tests

import (
	"testing"

	"quantrium-tech/form_pf_extractor_service/executors"
	curio "quantrium-tech/form_pf_extractor_service/executors/FEG/Curio"
	invesmentservices "quantrium-tech/form_pf_extractor_service/executors/FEG/InvesmentServices"
	privateinvestors "quantrium-tech/form_pf_extractor_service/executors/FEG/PrivateInvestors"
)

func TestQ17TermAndLineKinds(t *testing.T) {
	if executors.Q17Term("17_(d)_gross_performance") != "first" {
		t.Fatal("expected first")
	}
	if !executors.IsLiquidityLine("15") || !executors.IsLiquidityLine("16_a") {
		t.Fatal("liquidity lines")
	}
	if !executors.IsBalanceLine("8", false) || executors.IsBalanceLine("14_assets_level1", false) {
		t.Fatal("private/curio balance")
	}
	if !executors.IsBalanceLine("14_assets_level1", true) {
		t.Fatal("investment services level1 is balance")
	}
}

func TestUniqueFiles(t *testing.T) {
	got := executors.UniqueFiles(" a.xlsx ", "", "a.xlsx", "b.xlsx")
	if len(got) != 2 || got[0] != "a.xlsx" || got[1] != "b.xlsx" {
		t.Fatalf("got %v", got)
	}
}

func TestPrivateInvestorSourceFiles(t *testing.T) {
	fund := privateinvestors.NewPrivateFund("FEGPOF 2", "A", "FEG Private Opportunities Fund 2")
	fund.SetHolder(1, "/p/q1.xlsx", "q1.xlsx")
	fund.SetHolder(2, "/p/q2.xlsx", "q2.xlsx")
	fund.SetHolder(3, "/p/q3.xlsx", "q3.xlsx")
	fund.SetSharedSourceFiles("liquidity.xlsx", "")

	if got := fund.SourceFiles("8"); len(got) != 1 || got[0] != "q3.xlsx" {
		t.Fatalf("balance = %v", got)
	}
	if got := fund.SourceFiles("15"); len(got) != 1 || got[0] != "liquidity.xlsx" {
		t.Fatalf("liquidity = %v", got)
	}
	if got := fund.SourceFiles("17_(d)_gross_performance"); len(got) != 1 || got[0] != "q1.xlsx" {
		t.Fatalf("q17 first = %v", got)
	}
	if got := fund.SourceFiles("5_a"); len(got) != 0 {
		t.Fatalf("identity = %v", got)
	}
}

func TestInvestmentServicesSourceFiles(t *testing.T) {
	fund := invesmentservices.NewFund("FEGMP")
	fund.SetHolder("nav_tb", "/p/tb.xlsx", "tb.xlsx")
	fund.SetHolder("nav_report", "/p/nav.xlsx", "nav.xlsx")
	fund.SetSharedSourceFiles("liq.xlsx", "unfunded.xlsx")

	got := fund.SourceFiles("8")
	if len(got) != 2 || got[0] != "tb.xlsx" || got[1] != "unfunded.xlsx" {
		t.Fatalf("balance = %v", got)
	}
	if got := fund.SourceFiles("15"); len(got) != 1 || got[0] != "liq.xlsx" {
		t.Fatalf("liquidity = %v", got)
	}
	if got := fund.SourceFiles("17_(h)_net_performance"); len(got) != 1 || got[0] != "nav.xlsx" {
		t.Fatalf("q17 = %v", got)
	}
}

func TestCurioSourceFiles(t *testing.T) {
	fund := curio.NewFund("CURIV")
	sep := "september"
	third := "third"
	fund.SetHolder("nav_tb", "/p/tb.xlsx", nil, "tb.xlsx")
	fund.SetHolder("capital_register", "/p/sep.xlsx", &sep, "sep.xlsx")
	fund.SetHolder("capital_register", "/p/third.xlsx", &third, "third.xlsx")
	fund.SetSharedSourceFiles("liq.xlsx", "")

	got := fund.SourceFiles("8")
	if len(got) != 3 {
		t.Fatalf("balance = %v", got)
	}
	if got := fund.SourceFiles("17_(k)_gross_performance"); len(got) != 1 || got[0] != "sep.xlsx" {
		t.Fatalf("september q17 = %v", got)
	}
	if got := fund.SourceFiles("16_a"); len(got) != 1 || got[0] != "liq.xlsx" {
		t.Fatalf("liquidity = %v", got)
	}
}
