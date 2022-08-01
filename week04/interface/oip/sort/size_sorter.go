package sort

import (
	"go_8_mage/week04/interface/oip/common"
	"sort"
)

type SizeSorter struct {
	Tag string
}

func (s SizeSorter) Name() string{
	return s.Tag
}

func (s SizeSorter) Sort(products []*common.Product) []*common.Product{
	sort.Slice(products,func(i,j int) bool{
	//	按尺寸升序盘列
		return products[i].Size > products[j].Size
	})
	return products
}

/*
sort.slice 排序实例
func demoSortSlice(){
    a := []int{6,3,9,8,1,2,5,7}
    sort.Slice(a, func(i, j int) bool {
        return a[i]>a[j]
    })
    fmt.Println(a)
    //[9 8 7 6 5 3 2 1]
}
 */