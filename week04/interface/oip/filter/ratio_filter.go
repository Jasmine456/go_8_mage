package filter

import "go_8_mage/week04/interface/oip/common"

type RatioFilter struct {
	Tag string
}

func (a RatioFilter) Name() string {
	return a.Tag
}

func (self RatioFilter) Filter(products []*common.Product) []*common.Product {
	rect := make([]*common.Product,0,len(products))
	for _,product := range products{
		if product.RatioCount > 10 && product.PositiveRatio > 0.8 {
			rect = append(rect,product)
		}
	}
	return rect
}
