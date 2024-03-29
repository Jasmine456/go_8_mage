package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go_8_mage/week11/socket/common"
	"strconv"
	"sync"
	"time"
)

const TIME_LAYOUT = "2006-01-02"

var (
	loc *time.Location
)

func init() {
	loc, _ = time.LoadLocation("Asia/Shanghai")
}

//insert 通过stmt插入数据
func insert(db *sql.DB) {
	//	一条sql，插入2行记录
	stmt, err := db.Prepare("insert into student(name,province,city,enrollment)value(?,?,?,?),(?,?,?,?)")
	common.CheckError(err)
	//	字符串解析为时间，注意要使用time.ParseInLocatino() 函数指定时区，time.Parse()函数使用默认的UTC时区
	date1, err := time.ParseInLocation(TIME_LAYOUT, "2021-03-18", loc)
	common.CheckError(err)
	date2, err := time.ParseInLocation(TIME_LAYOUT, "2021-03-26", loc)
	common.CheckError(err)

	//	执行修改操作通过stmt.Exec，执行查询操作通过stmt.Query
	res, err := stmt.Exec("小明", "深圳", "深圳", date1, "小红", "上海", "上海", date2)
	common.CheckError(err)
	lastId, err := res.LastInsertId() //ID自增，用过的id（即使对应的行已delete）不会重复使用
	common.CheckError(err)
	fmt.Printf("after insert last id %d\n", lastId)
	rows, err := res.RowsAffected() // 插入两行，所以影响了2行
	common.CheckError(err)
	fmt.Printf("insert affect %d row\n", rows)
}

// replace 通过stmt插入（覆盖）数据
func replace(db *sql.DB) {
	//	由于name字段上有唯一索引，insert重复的name会报错。而使用replace会先删除，再插入
	stmt, err := db.Prepare("replace into student (name,province,city,enrollment) values (?,?,?,?), (?,?,?,?)")
	common.CheckError(err)
	//	字符串解析为时间。注意要使用time.ParseInLocation()函数执行时区，time.Parse()函数使用默认的UTC时区
	date1, err := time.ParseInLocation(TIME_LAYOUT, "2021-04-18", loc)
	common.CheckError(err)
	date2, err := time.ParseInLocation(TIME_LAYOUT, "2021-04-26", loc)
	common.CheckError(err)
	//	执行修改操作通过stmt.Exec，执行查询操作通过stmt.Query
	res, err := stmt.Exec("小明", "深圳", "深圳", date1, "小红", "上海", "上海", date2)
	common.CheckError(err)
	lastId, err := res.LastInsertId() //ID自增，用过的id（即使对应的行已delete）不会重复使用
	common.CheckError(err)
	fmt.Printf("after insert last id %d\n", lastId)
	rows, err := res.RowsAffected() //先删除，后插入，影响了4行
	common.CheckError(err)
	fmt.Printf("insert affect %d row\n", rows)

}

//update 通过stmt修改数据
func update(db *sql.DB) {
	//不同的city加不同的分数
	stmt, err := db.Prepare("update student set score=score+? where city=?")
	common.CheckError(err)
	//	执行修改操作通过stmt.Exec，执行查询操作通过stmt.Query
	res, err := stmt.Exec(10, "上海") //上海加10分
	common.CheckError(err)
	res, err = stmt.Exec(9, "深圳") //深圳加9分
	common.CheckError(err)
	lastId, err := res.LastInsertId()
	common.CheckError(err)
	fmt.Printf("after update last id %d\n", lastId)
	rows, err := res.RowsAffected() //where city=?名字的几行，就会影响几行
	common.CheckError(err)
	fmt.Printf("update affect %d row\n", rows)

}

//query 通过stmt查询数据
func query(db *sql.DB) {
	stmt, err := db.Prepare("select id,name,city,score from student where id>?")
	common.CheckError(err)
	//	执行修改操作通过stmt.Exec，执行查询操作通过stmt.Query
	rows, err := stmt.Query(2)
	common.CheckError(err)
	for rows.Next() {
		var id int
		var score float32
		var name, city string
		err = rows.Scan(&id, &name, &city, &score) //通过scan把db里的数据赋给go变量
		common.CheckError(err)
		fmt.Printf("id=%d,score=%.2f,name=%s,city=%s\n", id, score, name, city)
	}
}

//delete 通过stmt删除数据
func delete(db *sql.DB) {
	stmt, err := db.Prepare("delete from student where id>?")
	common.CheckError(err)
	//执行修改操作通过stmt.Exec，执行查询操作通过stmt.Query
	res, err := stmt.Exec(13) //删除id大于13的记录
	common.CheckError(err)
	rows, err := res.RowsAffected() //where id>?命中了几行，就会影响几行
	common.CheckError(err)
	fmt.Printf("delete affect %d row\n", rows)

}

//大量的数据行插入操作
func hugeInsert(db *sql.DB) {
	begin := time.Now()
	stmt, err := db.Prepare("insert into student (name,province,city,enrollment) values (?,?,?,?)")
	common.CheckError(err)
	date, err := time.ParseInLocation("20060102", "20211204", loc)

	for i := 0; i < 10000; i++ {
		stmt.Exec("宋江"+strconv.Itoa(i), "山西", "大同", date)
	}
	fmt.Printf("huge insert use %d ms\n", time.Since(begin).Milliseconds())

}
func goHugeInsert(db *sql.DB) {
	begin := time.Now()
	stmt, err := db.Prepare("insert into student (name,province,city,enrollment) values (?,?,?,?),(?,?,?,?),(?,?,?,?),(?,?,?,?)")
	common.CheckError(err)
	date, err := time.ParseInLocation("20060102", "20211204", loc)

	const P = 10
	task := make(chan int, 10000)
	for i := 0; i < 10000; i += 4 {
		task <- i
	}
	close(task)
	wg := sync.WaitGroup{}
	wg.Add(P)
	for i := 0; i < P; i++ {

		go func() {
			for {
				i, ok := <-task
				if !ok {
					wg.Done()
					return
				} else {
					stmt.Exec("宋江"+strconv.Itoa(i), "山西", "大同", date,
						"宋江"+strconv.Itoa(i+1), "山西", "大同", date,
						"宋江"+strconv.Itoa(i+2), "山西", "大同", date,
						"宋江"+strconv.Itoa(i+3), "山西", "大同", date)
				}
			}
		}()
	}
	wg.Wait()
	fmt.Printf("huge insert use %d ms\n", time.Since(begin).Milliseconds())

}

func traverse1(db *sql.DB) {
	var offset int
	begin := time.Now()
	stmt, _ := db.Prepare("select id,name,province from student limit ?,100")
	for i := 0; i < 1000; i++ {
		//t0 := time.Now()
		rows, _ := stmt.Query(offset)
		offset += 100
		//fmt.Println(i,time.Since(t0))

		for rows.Next() {
			var id int
			var name string
			var province string
			rows.Scan(&id, &name, &province)
		}
	}
	fmt.Println("total", time.Since(begin))
}

//traverse2 借助于主键自增ID，通过 where id>maxid遍历表
func traverse2(db *sql.DB) {
	var maxid int
	begin := time.Now()
	stmt, _ := db.Prepare("select id,name,province from student where id>? limit 100")
	for i := 0; i < 1000; i++ {
		//t0:=time.Now()
		rows, _ := stmt.Query(maxid)
		//fmt.Printf("%d %dus\n",i,time.Since(t0).Microseconds())

		for rows.Next() {
			var id int
			var name string
			var province string
			rows.Scan(&id, &name, &province)
			//fmt.Println(id,name,province)
			if id > maxid {
				maxid = id
			}
		}
	}
	fmt.Println("total", time.Since(begin))
}

func main() {
	db, err := sql.Open("mysql", "tester:123456@tcp(10.20.2.117:3306)/test?charset=utf8")
	common.CheckError(err)
	//insert(db)
	//replace(db)
	//update(db)
	//query(db)
	//delete(db)
	//hugeInsert(db)
	//goHugeInsert(db)
	//traverse1(db)
	traverse2(db)
}
