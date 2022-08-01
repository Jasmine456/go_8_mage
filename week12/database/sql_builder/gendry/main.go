package main

import (
	"database/sql"
	"fmt"
	"github.com/didi/gendry/builder"
	manager "github.com/didi/gendry/manager"
	_ "github.com/go-sql-driver/mysql"
	"go_8_mage/week11/socket/common"
)

func query() {
	where := map[string]interface{}{
		"city":     []string{"北京", "上海", "杭州"},
		"score<":   30,
		"addr":     builder.IsNotNull,
		"_orderby": "score desc",
		"_groupby": "province",
	}
	table := "student"
	fields := []string{"id", "name", "city", "score"}
	template, values, err := builder.BuildSelect(table, where, fields)
	common.CheckError(err)
	fmt.Println(template)
	fmt.Println(values...)
}

func insert() {
	data := []map[string]interface{}{
		{"name": "王五", "province": "河南", "city": "郑州", "enrollment": "2021-05-01"},
		{"name": "大王", "province": "浙江", "city": "杭州", "enrollment": "2021-04-01"},
	}
	table := "student"
	template, values, err := builder.BuildInsert(table, data)
	common.CheckError(err)
	fmt.Println(template)
	fmt.Println(values...)
}

func update() {
	where := map[string]interface{}{
		"city": []string{"北京", "上海", "杭州"},
	}
	data := map[string]interface{}{
		"score": 25,
	}
	table := "student"
	template, values, err := builder.BuildUpdate(table, where, data)
	common.CheckError(err)
	fmt.Println(template)
	fmt.Println(values...)

}
func delete() {
	where := map[string]interface{}{
		"city": "杭州",
	}
	table := "student"
	template, values, err := builder.BuildDelete(table, where)
	common.CheckError(err)
	fmt.Println(template)  //包含占位符的sql模板
	fmt.Println(values...) //占位符的具体值
}

func query2(db *sql.DB) {
	where := map[string]interface{}{
		"city":     []string{"北京", "上海", "杭州"},
		"score<":   30,
		"addr":     builder.IsNotNull,
		"_orderby": "score desc",
	}
	table := "student"
	fields := []string{"id", "name", "city", "score"}
	//	准备stmt模板
	template, values, err := builder.BuildSelect(table, where, fields)
	common.CheckError(err)
	//	执行stmt模板
	rows,err:=db.Query(template,values...)
	common.CheckError(err)
	for rows.Next(){
		var id int
		var name,city string
		var score float32
		err:=rows.Scan(&id,&name,&city,&score)
		common.CheckError(err)
		fmt.Printf("%d %s %s %.2f\n",id,name,city,score)
	}
}

func insert2(db *sql.DB){
	data := []map[string]interface{}{
		{"name": "王五", "province": "河南", "city": "郑州", "enrollment": "2021-05-01"},
		{"name": "大王", "province": "浙江", "city": "杭州", "enrollment": "2021-04-01"},
	}
	table := "student"
	template, values, err :=builder.BuildReplaceInsert(table,data)//使用replace
	common.CheckError(err)
	res,err:=db.Exec(template,values...)
	common.CheckError(err)
	rows,err:=res.RowsAffected()
	common.CheckError(err)
	fmt.Printf("insert affect %d rows\n",rows)
}

func update2(db *sql.DB){
	where:=map[string]interface{}{
		"city": []string{"北京", "上海", "杭州"},
	}
	data:=map[string]interface{}{
		"score":25,
	}
	table:="student"
	template,values,err:=builder.BuildUpdate(table,where,data)
	common.CheckError(err)
	res,err:=db.Exec(template,values...)
	common.CheckError(err)
	rows,err:=res.RowsAffected()
	common.CheckError(err)
	fmt.Printf("update affect %d rows\n",rows)
}

func delete2(db *sql.DB){
	where := map[string]interface{}{
		"city": "杭州",
	}
	table := "student"
	template, values, err:=builder.BuildDelete(table,where)
	common.CheckError(err)
	res,err:=db.Exec(template,values...)
	common.CheckError(err)
	rows,err:=res.RowsAffected()
	common.CheckError(err)
	fmt.Printf("delete affect %d rows]n",rows)
}


func main() {
	/*
		生成sql语句

	*/
	dbName := "test"
	user := "tester"
	password := "123456"
	host := "10.20.2.117"
	db, err := manager.New(dbName, user, password, host).Set(
		manager.SetCharset("utf8"),
	).Port(3306).Open(true)
	common.CheckError(err)
	//fmt.Println(db)
	//query()
	//insert()
	//update()
	//delete()
	//query2(db)
	//insert2(db)
	//update2(db)
	delete2(db)
}
