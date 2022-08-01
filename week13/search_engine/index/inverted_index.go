package index

import (
	"go_8_mage/week13/search_engine/common"
	"strings"
	"sync"
)

var (
	InvertedIndex sync.Map
)

func extractKeyOfInvertLink(user *common.User) []string {
	if user == nil {
		return nil
	}
	city := user.City
	kws := strings.Split(user.Keywords, "|")
	rect := make([]string, 0, len(kws))
	for _, ele := range kws {
		rect = append(rect, city+"_"+ele)
	}
	return rect
}

const (
	//最后一位为1：男
	BIT_MAN      uint32 = 1
	//倒数第二位为1：女
	BIT_WOMAN    uint32 = 1 << 1
	//倒数第三、四、五位表示学历，000无学历要求，001大专，010本科，011硕士，100博士
	BIT_JUNIOR   uint32 = 1 << 2
	BIT_BACHELOR uint32 = 1 << 3
	BIT_MASTER   uint32 = 1 << 4
	BIT_DOCTOR   uint32 = 1 << 5
)

//将两个32位的uint32合并为一个uint64
func Combine32To64(a,b uint32) uint64{
	return uint64(a) <<32|uint64(b)
}

//将一个uint64 拆为两个uint32
func chai64To32(c uint64)(uint32,uint32){
	docId:=uint32(c>>32)
	bits:=uint32(c<<32>>32)
	return docId,bits
}



//属性离散化(抽象) 为 位(bit)表示
// 000000000000000000000001 11 1
func BuildBitsByUserInfo(user *common.User) uint32 {
	bits := uint32(0)
	if user.Gender == "男" {
		bits |= BIT_MAN
	} else if user.Gender == "女" {
		bits |= BIT_WOMAN
	}
	switch  user.Degree {
	case "大专":
		bits|=BIT_JUNIOR
	case "本科":
		bits|=BIT_BACHELOR
	case "硕士":
		bits|=BIT_MASTER
	case "博士":
		bits|=BIT_DOCTOR
	}
	return bits
}




//把一个user放入倒排索引
func InsertUser2InvertIndex(user *common.User) {
	uid := user.Uid
	kws := extractKeyOfInvertLink(user)
	for _, word := range kws {
		if oldSlice, exists := InvertedIndex.Load(word); exists {
			newSlice := append(oldSlice.([]uint32), uid)
			InvertedIndex.Store(word, newSlice)
		} else {
			InvertedIndex.Store(word, []uint32{uid})
		}
	}

}

//delete 删除
func DeleteUserFromInvertLink(user *common.User) {
	uid := user.Uid
	kws := extractKeyOfInvertLink(user)
	for _, word := range kws {
		if oldSlice, exists := InvertedIndex.Load(word); exists {
			old := oldSlice.([]uint32)
			newSlice := make([]uint32, 0, len(old))
			for _, ele := range old {
				if ele != uid {
					newSlice = append(newSlice, ele)
				}
			}
			InvertedIndex.Store(word, newSlice)
		}
	}
}

const (
	KEYWORDS_PREFIX = "KW"
	CITY_PREFIX     = "CT"
	KEY_CONNECTOR   = "_"
)

func SearchRequest2LookUpRequest(request *common.SearchRequest) ([]string, []string, uint32, uint32, []uint32) {
	mustKeys := make([]string, 0, len(request.MustKeys))
	shouldKeys := make([]string, 0, len(request.ShouldKeys))
	var onFlag, offFlag uint32
	var orFlags []uint32 = make([]uint32, 0, 4)

	for _, ele := range request.MustKeys {
		// mustKeys = append(mustKeys, CITY_PREFIX+request.City+KEY_CONNECTOR+KEYWORDS_PREFIX+ele)
		sb := strings.Builder{}
		sb.WriteString(CITY_PREFIX)
		sb.WriteString(request.City)
		sb.WriteString(KEY_CONNECTOR)
		sb.WriteString(KEYWORDS_PREFIX)
		sb.WriteString(ele)
		mustKeys = append(mustKeys, sb.String())
	}
	for _, ele := range request.ShouldKeys {
		// shouldKeys = append(shouldKeys, CITY_PREFIX+request.City+KEY_CONNECTOR+KEYWORDS_PREFIX+ele)
		sb := strings.Builder{}
		sb.WriteString(CITY_PREFIX)
		sb.WriteString(request.City)
		sb.WriteString(KEY_CONNECTOR)
		sb.WriteString(KEYWORDS_PREFIX)
		sb.WriteString(ele)
		shouldKeys = append(shouldKeys, sb.String())
	}
	if request.Gender == "男" {
		onFlag = BIT_MAN
	} else if request.Gender == "女" {
		onFlag = BIT_WOMAN
	}
	for _, degree := range request.Degrees {
		switch degree {
		case "大专":
			orFlags = append(orFlags, BIT_JUNIOR)
		case "本科":
			orFlags = append(orFlags, BIT_BACHELOR)
		case "硕士":
			orFlags = append(orFlags, BIT_MASTER)
		case "博士":
			orFlags = append(orFlags, BIT_DOCTOR)
		}
	}
	return mustKeys, shouldKeys, onFlag, offFlag, orFlags
}
