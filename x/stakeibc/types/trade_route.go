package types

func (t TradeRoute) GetKey() []byte {
	return TradeRouteKeyFromDenoms(t.RewardDenomOnHostZone, t.TargetDenomOnHostZone)
}