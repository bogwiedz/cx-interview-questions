package basket_pricer

import "fmt"

type Offer map[string]OfferItem

type OfferItem struct {
	// percent `json:""` of discount [0-100]
	Discount int `json:"discount"`
	Buy      int `json:"buy"`
	Free     int `json:"free"`
}

type Basket map[string]int

type Catalogue map[string]float64

type pricer struct {
	// Base price, without discounts
	SubTotal float64

	Discount float64
	// Final price
	Total float64
}

func Calculate(offer Offer, catalogue Catalogue, basket *Basket) (*pricer, error) {
	var subTotal float64
	var discount float64

	for product, quantity := range *basket {
		var tmpPrice float64

		price, ok := catalogue[product]
		if !ok {
			return nil, fmt.Errorf("Product '%s' not found in catalogue", product)
		}

		// Calculate base price of product
		tmpPrice = float64(quantity) * price

		// update Sub-Total price
		subTotal += tmpPrice

		if change, ok := offer[product]; ok {
			// Get maximum discount

			var discountPercent float64
			var discountPack float64

			if change.Discount > 0 {
				discountPercent = tmpPrice * (float64(change.Discount) / 100.0)
			}

			if change.Buy > 0 && change.Free > 0 {

				discountPack = price * float64(quantity/(change.Buy+change.Free))
			}

			if discountPercent > discountPack {
				discount += discountPercent
			} else {
				discount += discountPack
			}
		}
	}
	p := pricer{
		SubTotal: subTotal,
		Discount: discount,
		Total:    subTotal - discount,
	}

	return &p, nil
}
