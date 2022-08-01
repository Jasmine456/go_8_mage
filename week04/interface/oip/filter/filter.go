package filter

import "go_8_mage/week04/interface/oip/common"

type Filter interface{
	Filter([]*common.Product) []*common.Product
	Name() string
}