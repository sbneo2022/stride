package v18

import sdk "github.com/cosmos/cosmos-sdk/types"

var (
	UpgradeName = "v18"

	// Redemption rate bounds updated to give ~3 months of slack on outer bounds
	RedemptionRateOuterMinAdjustment = sdk.MustNewDecFromStr("0.05")
	RedemptionRateOuterMaxAdjustment = sdk.MustNewDecFromStr("0.10")

	// Osmosis will have a slighly larger buffer with the redemption rate
	// since their yield is less predictable
	OsmosisChainId              = "osmosis-1"
	OsmosisRedemptionRateBuffer = sdk.MustNewDecFromStr("0.02")

	// Terra chain ID for delegation changes in progress
	TerraChainId = "phoenix-1"

	// Get Initial Redemption Rates for Unbonding Records Migration
	RedemptionRatesAtTimeOfProp = map[string]sdk.Dec{
		"comdex-1":     sdk.MustNewDecFromStr("1.204404927372203376"),
		"cosmoshub-4":  sdk.MustNewDecFromStr("1.299315098715274953"),
		"evmos_9001-2": sdk.MustNewDecFromStr("1.492396096716486696"),
		"injective-1":  sdk.MustNewDecFromStr("1.215553256473652866"),
		"juno-1":       sdk.MustNewDecFromStr("1.418210972076073590"),
		"osmosis-1":    sdk.MustNewDecFromStr("1.201353579705385297"),
		"phoenix-1":    sdk.MustNewDecFromStr("1.178171857075037002"),
		"sommelier-3":  sdk.MustNewDecFromStr("1.025900883208774724"),
		"stargaze-1":   sdk.MustNewDecFromStr("1.429976684963222047"),
		"umee-1":       sdk.MustNewDecFromStr("1.128473850654652585"),
	}

	// Get Amount Unbonded for each HostZone for Unbonding Records Migration
	StartingEstimateEpoch     = uint64(509)
	RedemptionRatesBeforeProp = map[string]map[uint64]sdk.Dec{
		"juno-1": {
			480: sdk.MustNewDecFromStr("1.4053501787364933"),
			481: sdk.MustNewDecFromStr("1.4053501787364933"),
			482: sdk.MustNewDecFromStr("1.4053501787364933"),
			484: sdk.MustNewDecFromStr("1.4053501787364933"),
			487: sdk.MustNewDecFromStr("1.4087287741035914"),
			488: sdk.MustNewDecFromStr("1.4087287741035914"),
			489: sdk.MustNewDecFromStr("1.4087287741035914"),
			490: sdk.MustNewDecFromStr("1.4099190036765492"),
			491: sdk.MustNewDecFromStr("1.4099190036765492"),
			493: sdk.MustNewDecFromStr("1.4099190036765492"),
			494: sdk.MustNewDecFromStr("1.4099190036765492"),
			495: sdk.MustNewDecFromStr("1.4138284392049727"),
			496: sdk.MustNewDecFromStr("1.4138284392049727"),
			497: sdk.MustNewDecFromStr("1.4138284392049727"),
			500: sdk.MustNewDecFromStr("1.4161992241945822"),
			501: sdk.MustNewDecFromStr("1.4161992241945822"),
			503: sdk.MustNewDecFromStr("1.4161992241945822"),
			504: sdk.MustNewDecFromStr("1.4161992241945822"),
			505: sdk.MustNewDecFromStr("1.417724248601981"),
			507: sdk.MustNewDecFromStr("1.417724248601981"),
			508: sdk.MustNewDecFromStr("1.417724248601981"),
		},
		"phoenix-1": {
			480: sdk.MustNewDecFromStr("1.1617706976314004"),
			481: sdk.MustNewDecFromStr("1.1617706976314004"),
			482: sdk.MustNewDecFromStr("1.1617706976314004"),
			483: sdk.MustNewDecFromStr("1.1617706976314004"),
			486: sdk.MustNewDecFromStr("1.1617706976314004"),
			487: sdk.MustNewDecFromStr("1.1617706976314004"),
			488: sdk.MustNewDecFromStr("1.1617706976314004"),
			489: sdk.MustNewDecFromStr("1.1617706976314004"),
			490: sdk.MustNewDecFromStr("1.1617706976314004"),
			491: sdk.MustNewDecFromStr("1.1617706976314004"),
			492: sdk.MustNewDecFromStr("1.1642615263477947"),
			493: sdk.MustNewDecFromStr("1.1642615263477947"),
			495: sdk.MustNewDecFromStr("1.1642615263477947"),
			496: sdk.MustNewDecFromStr("1.1740941880190199"),
			498: sdk.MustNewDecFromStr("1.1740941880190199"),
			499: sdk.MustNewDecFromStr("1.1740941880190199"),
			500: sdk.MustNewDecFromStr("1.175864139818246"),
			503: sdk.MustNewDecFromStr("1.175864139818246"),
			504: sdk.MustNewDecFromStr("1.1776233988812017"),
			505: sdk.MustNewDecFromStr("1.1776233988812017"),
			506: sdk.MustNewDecFromStr("1.1776233988812017"),
			507: sdk.MustNewDecFromStr("1.1776233988812017"),
			508: sdk.MustNewDecFromStr("1.177732197355669"),
		},
		"sommelier-3": {
			482: sdk.MustNewDecFromStr("1.0221309385498634"),
			483: sdk.MustNewDecFromStr("1.0221309385498634"),
			484: sdk.MustNewDecFromStr("1.0221309385498634"),
			486: sdk.MustNewDecFromStr("1.022666920867979"),
			488: sdk.MustNewDecFromStr("1.022666920867979"),
			489: sdk.MustNewDecFromStr("1.022666920867979"),
			490: sdk.MustNewDecFromStr("1.023749963371928"),
			493: sdk.MustNewDecFromStr("1.023749963371928"),
			495: sdk.MustNewDecFromStr("1.024808400658687"),
			496: sdk.MustNewDecFromStr("1.024808400658687"),
			497: sdk.MustNewDecFromStr("1.024808400658687"),
			499: sdk.MustNewDecFromStr("1.024808400658687"),
			501: sdk.MustNewDecFromStr("1.0256449837573494"),
			502: sdk.MustNewDecFromStr("1.0256449837573494"),
			503: sdk.MustNewDecFromStr("1.0256449837573494"),
			504: sdk.MustNewDecFromStr("1.0256449837573494"),
			505: sdk.MustNewDecFromStr("1.0259008616418266"),
			507: sdk.MustNewDecFromStr("1.0259008616418266"),
			508: sdk.MustNewDecFromStr("1.0259008616418266"),
		},
		"cosmoshub-4": {
			484: sdk.MustNewDecFromStr("1.2875942696916478"),
			485: sdk.MustNewDecFromStr("1.2875942696916478"),
			486: sdk.MustNewDecFromStr("1.2875942696916478"),
			487: sdk.MustNewDecFromStr("1.2875942696916478"),
			488: sdk.MustNewDecFromStr("1.2901383615073008"),
			489: sdk.MustNewDecFromStr("1.2901383615073008"),
			490: sdk.MustNewDecFromStr("1.2901383615073008"),
			491: sdk.MustNewDecFromStr("1.2901383615073008"),
			492: sdk.MustNewDecFromStr("1.2914303773961855"),
			493: sdk.MustNewDecFromStr("1.2914303773961855"),
			494: sdk.MustNewDecFromStr("1.2914303773961855"),
			495: sdk.MustNewDecFromStr("1.2914303773961855"),
			496: sdk.MustNewDecFromStr("1.2946765429135605"),
			497: sdk.MustNewDecFromStr("1.2946765429135605"),
			498: sdk.MustNewDecFromStr("1.2946765429135605"),
			499: sdk.MustNewDecFromStr("1.2946765429135605"),
			500: sdk.MustNewDecFromStr("1.2966353732132307"),
			501: sdk.MustNewDecFromStr("1.2966353732132307"),
			502: sdk.MustNewDecFromStr("1.2966353732132307"),
			503: sdk.MustNewDecFromStr("1.2966353732132307"),
			504: sdk.MustNewDecFromStr("1.2986468613477709"),
			505: sdk.MustNewDecFromStr("1.2986468613477709"),
			506: sdk.MustNewDecFromStr("1.2986468613477709"),
			507: sdk.MustNewDecFromStr("1.2986468613477709"),
			508: sdk.MustNewDecFromStr("1.2990760366941585"),
			509: sdk.MustNewDecFromStr("1.2990760366941585"),
		},
		"comdex-1": {
			485: sdk.MustNewDecFromStr("1.1865251410281585"),
			486: sdk.MustNewDecFromStr("1.1865251410281585"),
			488: sdk.MustNewDecFromStr("1.1891618527186743"),
			489: sdk.MustNewDecFromStr("1.1891618527186743"),
			490: sdk.MustNewDecFromStr("1.1891618527186743"),
			491: sdk.MustNewDecFromStr("1.1891618527186743"),
			492: sdk.MustNewDecFromStr("1.1926667940346496"),
			493: sdk.MustNewDecFromStr("1.1926667940346496"),
			494: sdk.MustNewDecFromStr("1.1926667940346496"),
			495: sdk.MustNewDecFromStr("1.1926667940346496"),
			496: sdk.MustNewDecFromStr("1.1973725182109523"),
			497: sdk.MustNewDecFromStr("1.1973725182109523"),
			498: sdk.MustNewDecFromStr("1.1973725182109523"),
			499: sdk.MustNewDecFromStr("1.1973725182109523"),
			500: sdk.MustNewDecFromStr("1.2004214235326311"),
			501: sdk.MustNewDecFromStr("1.2004214235326311"),
			502: sdk.MustNewDecFromStr("1.2004214235326311"),
			503: sdk.MustNewDecFromStr("1.2004214235326311"),
			504: sdk.MustNewDecFromStr("1.2034589430292977"),
			505: sdk.MustNewDecFromStr("1.2034589430292977"),
			506: sdk.MustNewDecFromStr("1.2034589430292977"),
			507: sdk.MustNewDecFromStr("1.2034589430292977"),
			508: sdk.MustNewDecFromStr("1.2040062434963579"),
		},
		"injective-1": {
			464: sdk.MustNewDecFromStr("1.1500807256744128"),
			465: sdk.MustNewDecFromStr("1.1500807256744128"),
			466: sdk.MustNewDecFromStr("1.1500807256744128"),
			467: sdk.MustNewDecFromStr("1.1500807256744128"),
			468: sdk.MustNewDecFromStr("1.1500807256744128"),
			469: sdk.MustNewDecFromStr("1.1500807256744128"),
			470: sdk.MustNewDecFromStr("1.1500807256744128"),
			471: sdk.MustNewDecFromStr("1.1500807256744128"),
			472: sdk.MustNewDecFromStr("1.1500807256744128"),
			473: sdk.MustNewDecFromStr("1.1500807256744128"),
			474: sdk.MustNewDecFromStr("1.1500807256744128"),
			475: sdk.MustNewDecFromStr("1.1500807256744128"),
			476: sdk.MustNewDecFromStr("1.1500807256744128"),
			477: sdk.MustNewDecFromStr("1.1500807256744128"),
			478: sdk.MustNewDecFromStr("1.1500807256744128"),
			479: sdk.MustNewDecFromStr("1.1500807256744128"),
			480: sdk.MustNewDecFromStr("1.1500807256744128"),
			481: sdk.MustNewDecFromStr("1.1500807256744128"),
			482: sdk.MustNewDecFromStr("1.1500807256744128"),
			483: sdk.MustNewDecFromStr("1.1500807256744128"),
			484: sdk.MustNewDecFromStr("1.1500807256744128"),
			485: sdk.MustNewDecFromStr("1.1500807256744128"),
			486: sdk.MustNewDecFromStr("1.1500807256744128"),
			487: sdk.MustNewDecFromStr("1.1500807256744128"),
			488: sdk.MustNewDecFromStr("1.1500807256744128"),
			489: sdk.MustNewDecFromStr("1.1500807256744128"),
			498: sdk.MustNewDecFromStr("1.1500807256744128"),
			499: sdk.MustNewDecFromStr("1.1500807256744128"),
			500: sdk.MustNewDecFromStr("1.212374830290344"),
			501: sdk.MustNewDecFromStr("1.212374830290344"),
			502: sdk.MustNewDecFromStr("1.212374830290344"),
			503: sdk.MustNewDecFromStr("1.212374830290344"),
			504: sdk.MustNewDecFromStr("1.2147869431614962"),
			505: sdk.MustNewDecFromStr("1.2147869431614962"),
			507: sdk.MustNewDecFromStr("1.2147869431614962"),
			508: sdk.MustNewDecFromStr("1.214906161115688"),
		},
		"evmos_9001-2": {
			492: sdk.MustNewDecFromStr("1.4846653184918148"),
			493: sdk.MustNewDecFromStr("1.4846653184918148"),
			494: sdk.MustNewDecFromStr("1.4846653184918148"),
			495: sdk.MustNewDecFromStr("1.4877949604847145"),
			496: sdk.MustNewDecFromStr("1.4877949604847145"),
			497: sdk.MustNewDecFromStr("1.4877949604847145"),
			499: sdk.MustNewDecFromStr("1.4895996973002357"),
			500: sdk.MustNewDecFromStr("1.4895996973002357"),
			501: sdk.MustNewDecFromStr("1.490678848400448"),
			502: sdk.MustNewDecFromStr("1.490678848400448"),
			503: sdk.MustNewDecFromStr("1.490678848400448"),
			504: sdk.MustNewDecFromStr("1.4917210211376328"),
			505: sdk.MustNewDecFromStr("1.4917210211376328"),
			507: sdk.MustNewDecFromStr("1.4918520366929944"),
			508: sdk.MustNewDecFromStr("1.4918520366929944"),
		},
		"osmosis-1": {
			492: sdk.MustNewDecFromStr("1.1963688867243563"),
			493: sdk.MustNewDecFromStr("1.1963688867243563"),
			494: sdk.MustNewDecFromStr("1.1963688867243563"),
			495: sdk.MustNewDecFromStr("1.1970692695006457"),
			496: sdk.MustNewDecFromStr("1.1970692695006457"),
			497: sdk.MustNewDecFromStr("1.1970692695006457"),
			498: sdk.MustNewDecFromStr("1.1987054330176354"),
			499: sdk.MustNewDecFromStr("1.1987054330176354"),
			500: sdk.MustNewDecFromStr("1.1987054330176354"),
			501: sdk.MustNewDecFromStr("1.1997005791572002"),
			502: sdk.MustNewDecFromStr("1.1997005791572002"),
			503: sdk.MustNewDecFromStr("1.1997005791572002"),
			504: sdk.MustNewDecFromStr("1.200668177509732"),
			505: sdk.MustNewDecFromStr("1.200668177509732"),
			506: sdk.MustNewDecFromStr("1.200668177509732"),
			507: sdk.MustNewDecFromStr("1.2008572054271187"),
			508: sdk.MustNewDecFromStr("1.2008572054271187"),
			509: sdk.MustNewDecFromStr("1.2008572054271187"),
		},
		"stargaze-1": {
			492: sdk.MustNewDecFromStr("1.4221879666345822"),
			493: sdk.MustNewDecFromStr("1.4221879666345822"),
			494: sdk.MustNewDecFromStr("1.4221879666345822"),
			495: sdk.MustNewDecFromStr("1.4234096730900296"),
			496: sdk.MustNewDecFromStr("1.4234096730900296"),
			498: sdk.MustNewDecFromStr("1.425950601619048"),
			499: sdk.MustNewDecFromStr("1.425950601619048"),
			500: sdk.MustNewDecFromStr("1.425950601619048"),
			501: sdk.MustNewDecFromStr("1.4273024812101673"),
			502: sdk.MustNewDecFromStr("1.4273024812101673"),
			503: sdk.MustNewDecFromStr("1.4273024812101673"),
			504: sdk.MustNewDecFromStr("1.4288169571251907"),
			505: sdk.MustNewDecFromStr("1.4288169571251907"),
			506: sdk.MustNewDecFromStr("1.4288169571251907"),
			507: sdk.MustNewDecFromStr("1.429363817000515"),
			508: sdk.MustNewDecFromStr("1.429363817000515"),
		},
		"umee-1": {
			505: sdk.MustNewDecFromStr("1.1275274920184462"),
		},
	}
)