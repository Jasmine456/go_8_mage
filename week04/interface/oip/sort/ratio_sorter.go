package sort

import (
	"go_8_mage/week04/interface/oip/common"
	"sort"
)

type RatioSorter struct {
	Tag string
}

func (r RatioSorter) Name() string{
	return r.Tag
}

func(r RatioSorter) Sort(products []*common.Product) []*common.Product {
	sort.Slice(products,func(i,j int) bool{
		// 按好评率降序排列
		return products[i].PositiveRatio > products[j].PositiveRatio
	})
	return products
}
