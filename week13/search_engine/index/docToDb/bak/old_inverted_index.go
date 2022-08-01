package index

import (
	"go_8_mage/week13/search_engine/common"
	"honnef.co/go/tools/go/ir"
	"strings"
	"sync"
)

var (
	InvertedIndex sync.Map
)

func extractKeyOfInvertLink(user *common.User) []string {
	if user ==nil{
		return nil
	}
	return strings.Split(user.Keywords,"|")
}

//把一个user放入倒排索引
func InsertUser2InvertIndex(user *common.User){
	uid:=user.Uid
	kws:=extractKeyOfInvertLink(user)
	for _,word:=range kws{
		if oldSlice,exists:=InvertedIndex.Load(word);exists{
			newSlice:=append(oldSlice.([]uint32),uid)
			InvertedIndex.Store(word,newSlice)
		} else{
			InvertedIndex.Store(word,[]uint32{uid})
		}
	}

}

//delete 删除
func DeleteUserFromInvertLink(user *common.User) {
	uid:=user.Uid
	kws:=extractKeyOfInvertLink(user)
	for _,word:=range kws{
		if oldSlice,exists:=InvertedIndex.Load(word);exists{
			old:=oldSlice.([]uint32)
			newSlice:=make([]uint32,0,len(old))
			for _,ele:=range old{
				if ele!=uid{
					newSlice=append(newSlice,ele)
				}
			}
			InvertedIndex.Store(word,newSlice)
		}
	}
}