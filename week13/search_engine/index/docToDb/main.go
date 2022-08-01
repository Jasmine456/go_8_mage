package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"go_8_mage/week13/search_engine/common"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

//把一行内容解析为一个User实体
func parseLine(line string) *common.User {
	arr := strings.Split(line, ",")
	if len(arr) >= 5 { //避免数组越界
		uid, err := strconv.ParseUint(arr[0], 10, 64)
		if err != nil || uid <= 0 {
			glog.Errorf("invalid line: [%s]", line)
			return nil
		}
		return &common.User{
			Uid:      uint32(uid),
			Keywords: arr[1],
			Gender:   arr[2],
			Degree:   arr[3],
			City:     arr[4],
		}
	} else{
		glog.Errorf("line %s 不够5列\n",line)
		return nil
	}
}

//把文件里的内容写入数据库
func MassInsert(db *gorm.DB, corpusFile string) {
	fin, err := os.Open(corpusFile)
	if err != nil {
		glog.Errorf("open corpusFile %s failed: %s\n", corpusFile, err.Error())
		return
	}
	defer fin.Close()

	const Batch_SIZE = 500 //每次往数据库中批量插入BATCH_SIZE条记录
	users := make([]common.User, 0, Batch_SIZE)
	reader := bufio.NewReader(fin) //带 buffer的方式读文本文件
	for {
		line, err := reader.ReadString('\n') //读取一行
		if err != nil {
			if err == io.EOF {
				user := parseLine(line) //已到文件结尾，最后没有换行符
				if user != nil {
					users = append(users, *user)
				}
			} else {
				glog.Errorf("read corpusFile %s failed:%s\n", corpusFile, err.Error())
			}
			break
		} else {
			line = strings.TrimRight(line, "\n") //还没到文件结尾，最后有换行符，需要去掉
			user := parseLine(line)
			if user != nil {
				users = append(users, *user)
			}
		}
		if len(users) >= Batch_SIZE {
			//批量一次性写入mysql
			//在实际工作中，所有error都要捕获，都要打到日志文件里面
			if res := db.Create(users); res.Error != nil {
				glog.Errorf("batch insert to database filed:%s\n", res.Error.Error())
			}
			//清空users
			users = make([]common.User, 0, Batch_SIZE) //批量插入后清空users
		}
	}
	//	users中还残留一些未DB的记录
	if len(users) > 0 {
		if res := db.Create(users); res.Error != nil {
			glog.Errorf("batch insert to database failed: %s\n", res.Error.Error())
		}
	}

}

func main() {
	flag.Parse()
	defer glog.Flush()

	dns := "tester:123456@tcp(10.20.2.117:3306)/test?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		glog.Fatalf("connect to database failed:%s\n", err.Error()) //打完日志后会调用os.Exit(255)
	}
	now:=time.Now()
	MassInsert(db,"../../data/doc.txt")
	glog.Infoln("insert into DB finsh")
	fmt.Println(time.Since(now))

}
