package test

import (
	"fmt"
	"go_8_mage/week13/search_engine/common"
	"go_8_mage/week13/search_engine/index"
	"testing"
)

func TestInvertedIndex(t *testing.T) {
	user := &common.User{
		Uid:      729325696,
		Keywords: "立功|劳动法|开除|员工|规定|请假",
		Gender:   "女",
		City:     "重庆",
		Degree:   "本科",
	}
	index.InsertUser2InvertIndex(user)
	user = &common.User{
		Uid:      729325697,
		Keywords: "立功|劳动法|开除|员工|规定|请假",
		Gender:   "男",
		City:     "重庆",
		Degree:   "本科",
	}
	index.InsertUser2InvertIndex(user)
	user = &common.User{
		Uid:      729325698,
		Keywords: "立功|劳动法|开除|员工|规定|请假",
		Gender:   "男",
		City:     "重庆",
		Degree:   "博士",
	}
	index.InsertUser2InvertIndex(user)

	request := &common.SearchRequest{
		MustKeys:   []string{"立功"},
		ShouldKeys: nil,
		Gender:     "男",
		City:       "重庆",
		Degrees:    []string{"本科", "硕士"},
	}
	mustKeys,shouldKeys,onFlag,offFlag,orFlags:=index.SearchRequest2LookUpRequest(request)
	uids := index.LookUp(mustKeys,shouldKeys,onFlag,offFlag,orFlags)
	fmt.Println(len(uids))

}
