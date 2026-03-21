package executors

// CalculateGrossAssets calculates the Gross Assets based on the given parameters.
func CalculateGrossAssets(totalAsset float64, unfundedCommitments float64) float64 {

	if unfundedCommitments > 0 {
		return totalAsset + unfundedCommitments
	}

	return totalAsset
}

// CalculateNetAsset calculates the Net Assets based on the given parameters.
func CalculateNetAsset(partnersCapital float64, unfundedCommitments float64) float64 {

	if unfundedCommitments > 0 {
		return partnersCapital + unfundedCommitments
	}

	return partnersCapital

}

func CalcCBLiabilities(gav float64, nav float64, fairValueLiabilites float64) float64 {
	return gav - nav - fairValueLiabilites
}
