package basket_pricer

import "fmt"

type Offer map[string]OfferItem

type OfferItem struct {
	// percent `json:""` of discount [0-100]
	Discount int `json:"discount"`
	// buy pack offer
	// buy - number of items for which customer will pay
	Buy int `json:"buy"`
	// free - numbrer of free items
	Free int `json:"free"`
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

func Calculate(offer Offer, catalogue Catalogue, basket Basket) (*pricer, error) {
	var subTotal float64
	var discount float64

	for product, quantity := range basket {
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

			// Get discount base on percentage discount
			if change.Discount > 0 && change.Discount <= 100 {
				discountPercent = tmpPrice * (float64(change.Discount) / 100.0)
			}

			// Get discount base on buy X get Y free discount
			if change.Buy > 0 && change.Free > 0 {
				discountPack = float64(change.Free) * price * float64(quantity/(change.Buy+change.Free))
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
