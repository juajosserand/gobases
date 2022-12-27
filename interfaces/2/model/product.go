package model

const (
	SmallProduct  = "S"
	MediumProduct = "M"
	LargeProduct  = "L"
)

type product struct {
	Cost, MaintenanceCost, ShippingCost float64
}

func (p *product) Price() float64 {
	return p.Cost + p.MaintenanceCost + p.ShippingCost
}
