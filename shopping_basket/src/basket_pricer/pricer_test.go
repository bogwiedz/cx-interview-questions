package basket_pricer

import (
	"strings"
	"testing"
)

func assert(t *testing.T, expected, actual interface{}, name ...string) {
	if expected != actual {
		t.Errorf("%s: expected '%v' got `%v'", strings.Join(name, " "), expected, actual)
	}
}

var basicCatalogue = Catalogue{
	"Baked Beans":      0.99,
	"Biscuits":         1.20,
	"Sardines":         1.89,
	"Shampoo (Small)":  2.00,
	"Shampoo (Medium)": 2.50,
	"Shampoo (Large)":  3.50,
}

func TestBasicBasket1(t *testing.T) {
	offer := Offer{
		"Baked Beans": OfferItem{
			Buy:  2,
			Free: 1,
		},
		"Sardines": OfferItem{
			Discount: 25,
		},
	}
	pricer, _ := Calculate(offer, basicCatalogue, Basket{
		"Baked Beans": 4,
		"Biscuits":    1,
	})

	assert(t, 516, int(pricer.SubTotal*100+0.5), "Sub-Total")
	assert(t, 99, int(pricer.Discount*100+0.5), "Discount")
	assert(t, 417, int(pricer.Total*100+0.5), "Total")
}

func TestBasicBasket2(t *testing.T) {
	offer := Offer{
		"Baked Beans": OfferItem{
			Buy:  2,
			Free: 1,
		},
		"Sardines": OfferItem{
			Discount: 25,
		},
	}
	pricer, _ := Calculate(offer, basicCatalogue, Basket{
		"Baked Beans": 2,
		"Biscuits":    1,
		"Sardines":    2,
	})

	assert(t, 696, int(pricer.SubTotal*100+0.5), "Sub-Total")
	assert(t, 95, int(pricer.Discount*100+0.5), "Discount")
	assert(t, 601, int(pricer.Total*100+0.5), "Total")
}

func TestBuyPack1(t *testing.T) {
	offer := Offer{
		"Biscuits": OfferItem{
			Buy:  3,
			Free: 1,
		},
		"Sardines": OfferItem{
			Discount: 25,
		},
	}
	pricer, _ := Calculate(offer, basicCatalogue, Basket{
		"Biscuits":        11,
		"Shampoo (Small)": 4,
	})

	assert(t, 2120, int(pricer.SubTotal*100+0.5), "Sub-Total")
	assert(t, 240, int(pricer.Discount*100+0.5), "Discount")
	assert(t, 1880, int(pricer.Total*100+0.5), "Total")
}

func TestBuyPack2(t *testing.T) {
	offer := Offer{
		"Biscuits": OfferItem{
			Buy:  3,
			Free: 2,
		},
		"Sardines": OfferItem{
			Discount: 25,
		},
	}
	pricer, _ := Calculate(offer, basicCatalogue, Basket{
		"Biscuits":        11,
		"Shampoo (Small)": 4,
	})

	assert(t, 2120, int(pricer.SubTotal*100+0.5), "Sub-Total")
	assert(t, 480, int(pricer.Discount*100+0.5), "Discount")
	assert(t, 1640, int(pricer.Total*100+0.5), "Total")
}
