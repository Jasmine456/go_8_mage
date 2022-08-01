package recall

import "go_8_mage/week04/interface/oip/common"

type Recaller interface {
	Recall(n int) []*common.Product //生成一批推荐候选集
	Name() string
}
