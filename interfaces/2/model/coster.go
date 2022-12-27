package model

import "errors"

var (
	errInvalidProductType = errors.New("invalid product type")
)

// Coster represents anything that can be priced
// based on its individual costs.
type Coster interface {
	Price() float64
}

// NewCoster returns a Coster based on the provided type and price.
func NewCoster(t string, price float64) (Coster, error) {
	switch {
	case t == SmallProduct:
		return &product{
			Cost: price,
		}, nil
	case t == MediumProduct:
		return &product{
			Cost:            price,
			MaintenanceCost: price * 0.03,
		}, nil
	case t == LargeProduct:
		return &product{
			Cost:            price,
			MaintenanceCost: price * 0.06,
			ShippingCost:    2500,
		}, nil
	}

	return nil, errInvalidProductType
}
