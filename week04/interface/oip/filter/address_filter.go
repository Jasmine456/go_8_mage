package filter

import "go_8_mage/week04/interface/oip/common"

type AddressFilter struct {
	Tag string
	City string
}

func (a AddressFilter) Name()string{
	return a.Tag
}

func (a AddressFilter) Filter(products []*common.Product) []*common.Product{
	rect := make([]*common.Product,0,len(products))
	for _,product:= range products{
		if product.ShipAddress == a.City{
			rect = append(rect,product)
		}
	}
	return rect
}
