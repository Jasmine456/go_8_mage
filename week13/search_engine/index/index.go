package index

import (
	"fmt"
	"github.com/bytedance/sonic"
	"go_8_mage/week13/search_engine/common"
	"gorm.io/gorm"
	"math"
	"sync"
	"sync/atomic"
)

func SerializeUser(user *common.User) []byte {
	if bs, err := sonic.Marshal(user); err != nil {
		return nil
	} else {
		return bs
	}

}

//遍历mysql里的user表，把每一个user写入正排索引和倒排索引
func BuildIndex(db *gorm.DB, storage *Badger) {
	var maxid = -1
	const BATCH = 500
	for {
		users := make([]common.User, 0, BATCH)
		db.Where("id>?", maxid).Limit(BATCH).Find(&users)
		if len(users) > 0 {
			for _, user := range users {
				if user.Id > maxid {
					maxid = user.Id
				}
				InsertUser2InvertIndex(&user) //写入倒排索引
				//写入正排索引
				k := IntToBytes(user.Uid)
				v := SerializeUser(&user)
				if v != nil {
					storage.Set(k, v)
				}
			}
		} else {
			break
		}
	}
}

const(
	MAX_INVERT_LEN=1000000
)

//设置一个固定长度的内存池，防止重复GC，拉低程序性能
var (
	arrPool=sync.Pool{
		New: func() any{
			arr:=make([]uint32,MAX_INVERT_LEN)
			return arr
		},
	}
)

//课堂写的不完善函数
//func traversInvertList(link []uint64, onFlag uint32) []uint32 {
//	const SEG_LEN = 10000
//	routineCnt := (len(link) + SEG_LEN - 1) / SEG_LEN
//	wg:=sync.WaitGroup{}
//	wg.Add(routineCnt)
//	collection:=make(chan uint32, len(link))
//	for i:=0;i<routineCnt;i++{//启动多个子协程，并行遍历 link
//		go func (i int){
//			defer wg.Done()
//			begin:=i*SEG_LEN
//			end:=begin+SEG_LEN
//			if end>len(link){
//				end=len(link)
//			}
//			for j:=begin;j<end;j++{
//				id:=link[j]
//				docid,bits:=chai64To32(id)
//				if bits &onFlag == onFlag{
//					collection <- docid
//				}
//			}
//		}(i)
//	}
//
//	wg.Wait()
//	close(collection)
//	//rect:=make([]uint32,0,len(collection))
//	rect:=arrPool.Get().([]uint32)
//	for ele:=range collection{
//		rect=append(rect,ele)
//	}
//	return rect
//
//}

//traverseInvertList 遍历某一条倒排链
func traverseInvertList(list []uint64, onFlag uint32, offFlag uint32, orFlags []uint32) ([]uint32, int) {
	const INVERTED_SEG_LEN = 10000   //倒排链很长时切成很多小的片段，并行遍历
	rect := arrPool.Get().([]uint32) //从缓存池里取一个int切片
	// rect := mallocSlice()
	segCount := (len(list) + INVERTED_SEG_LEN - 1) / INVERTED_SEG_LEN
	wg := sync.WaitGroup{}
	wg.Add(segCount)
	var idx int32 = 0
	for segIndex := 0; segIndex < segCount; segIndex++ {
		//分段并行遍历
		go func(segIndex int) {
			defer wg.Done()
			begin := segIndex * INVERTED_SEG_LEN
			end := (segIndex + 1) * INVERTED_SEG_LEN
			for i := begin; i < len(list) && i < end; i++ {
				docId, bits := chai64To32(list[i])
				if docId > 0 && (bits&onFlag == onFlag) && (bits&offFlag == 0) { //确保有效元素都大于0，onFlag和offFlag都满足
					orOk := false
					if len(orFlags) == 0 {
						orOk = true
					} else {
						for _, orFlag := range orFlags {
							if orFlag > 0 && bits&orFlag == orFlag {
								orOk = true
								break
							}
						}
					}
					if orOk { //满足orFlags
						index := atomic.AddInt32(&idx, 1) - 1
						if int(index) >= cap(rect) { //强行截断
							break
						}
						rect[index] = docId //index是多协程共享的，slice支持多协程并发修改，多个协程不会云修改rect里的同一个元素
					}
				}
			}
		}(segIndex)
	}
	wg.Wait()
	if int(idx) < cap(rect) {
		rect[idx] = 0 //最后一个元素置为0，这一个特殊的标记
	}
	return rect, int(idx)
}


//mustFind 关键词必须都命中
func mustFind(words []string, onFlag uint32, offFlag uint32, orFlags []uint32) map[uint32]bool {
	if len(words) == 0 {
		return nil
	}

	results := make([][]uint32, len(words))
	wg := sync.WaitGroup{}
	wg.Add(len(words))
	shortestListIndex := -1
	shortestListLen := math.MaxInt32 //取最短的那个长度
	for i, word := range words {
		go func(key string, i int) { //注意：在子协程中使用for range生成的变量时一定作为参数传给子协程
			defer wg.Done()
			lst, exists := InvertedIndex.Load(key)
			if exists {
				arr := lst.([]uint64)
				if len(arr) > 0 {
					var cl int
					results[i], cl = traverseInvertList(arr, onFlag, offFlag, orFlags)
					if cl < shortestListLen {
						shortestListIndex = i
						shortestListLen = cl //由于并发的缘故，shortestListLen不一定是最短的长度，但这个没关系，shortestListLen只是用来决定hitCountMap的Capacity
					}
				}
			}
		}(word, i) //range产生的参数一定要按值传到routine里面去，因为range每轮迭代使用的是同一个地址
	}
	wg.Wait()

	if shortestListLen == math.MaxInt32 {
		return map[uint32]bool{}
	}
	//先把最短的那条放到hitCountMap里
	hitCountMap := make(map[uint32]int, shortestListLen)
	arr := results[shortestListIndex]
	for _, ele := range arr {
		if ele == 0 { //遍历到arr最后一个元素了
			break
		}
		hitCountMap[ele] = 1
	}
	arrPool.Put(arr) //读完后归还给traverseResultPool
	// freeSlice(arr)

	//再遍历其余的，增加hitCountMap中的计数
	for i, arr := range results {
		if arr == nil || i == shortestListIndex {
			continue
		}
		if len(hitCountMap) > 0 {
			dup := make(map[uint32]bool, len(hitCountMap)) //对单条倒排上的元素进行排重
			for _, ele := range arr {
				if ele == 0 { //遍历到arr最后一个元素了
					break
				}
				if cnt, exists := hitCountMap[ele]; exists && !dup[ele] {
					hitCountMap[ele] = cnt + 1
					dup[ele] = true
				}
			}
		}
		arrPool.Put(arr) //读完后归还给arrPool
		// freeSlice(arr)
	}
	rect := make(map[uint32]bool, len(hitCountMap))
	for docId, cnt := range hitCountMap {
		if cnt >= len(words) { //如果一个doc包含了所有的words
			rect[docId] = true
		}
	}
	return rect
}

//shouldFind 命中一个关键词即可
func shouldFind(words []string, onFlag uint32, offFlag uint32, orFlags []uint32) map[uint32]bool {
	if len(words) == 0 {
		return nil
	}

	results := make([][]uint32, len(words))
	wg := sync.WaitGroup{}
	wg.Add(len(words))
	var total int32 = 0
	for i, word := range words {
		go func(key string, i int) { //注意：在子协程中使用for range生成的变量时一定作为参数传给子协程
			defer wg.Done()
			lst, exists := InvertedIndex.Load(key)
			if exists {
				arr := lst.([]uint64)
				if len(arr) > 0 {
					results[i], _ = traverseInvertList(arr, onFlag, offFlag, orFlags)
					atomic.AddInt32(&total, int32(len(results[i])))
				}
			}
		}(word, i) //range产生的参数一定要按值传到routine里面去，因为range每轮迭代使用的是同一个地址
	}
	wg.Wait()

	rect := make(map[uint32]bool, atomic.LoadInt32(&total))
	for _, arr := range results {
		if arr == nil {
			continue
		}
		for _, ele := range arr {
			if ele == 0 { //遍历到arr最后一个元素了
				break
			}
			rect[ele] = true
		}
		arrPool.Put(arr) //读完后归还给arrPool
		// freeSlice(arr)
	}
	return rect
}

func LookUp(mustKeys, shouldKeys []string, onFlag uint32, offFlag uint32, orFlags []uint32) []uint32 {
	haveMust := false
	haveShould := false
	if len(mustKeys) > 0 {
		haveMust = true
	}
	if len(shouldKeys) > 0 {
		haveShould = true
	}
	if !haveMust && !haveShould {
		return nil
	}

	//must和should并行检索
	var mustResult map[uint32]bool
	var shouldResult map[uint32]bool
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		mustResult = mustFind(mustKeys, onFlag, offFlag, orFlags)
	}()
	go func() {
		defer wg.Done()
		shouldResult = shouldFind(shouldKeys, onFlag, offFlag, orFlags)
	}()
	wg.Wait()

	var finalResult map[uint32]bool = make(map[uint32]bool, 100)
	if haveMust && haveShould {
		//mustResult和shouldResult取交集
		for docid := range mustResult {
			if _, exists := shouldResult[docid]; exists {
				finalResult[docid] = true
			}
		}
	} else if haveMust {
		finalResult = mustResult
	} else if haveShould {
		finalResult = shouldResult
	}

	rect := make([]uint32, 0, len(finalResult))
	fmt.Println(len(finalResult))
	for ele := range finalResult {
		rect = append(rect, ele)
	}

	return rect
}
