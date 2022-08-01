package main

import "fmt"

/*
3. 改变切片中的元素（元素是结构体）。
```go
type Student struct {
	weight float32 //体重，kg
	height float32 //身高，m
	bmi    float32 //weight/height^2
}

//计算每个学生的BMI。函数返回后BMI并没有改变，如何把程序改对？
func CalBMI1(students []Student) {
	for _, student := range students {
		student.bmi = student.weight / (student.height * student.height)
	}
}
 */
type Student struct {
	weight float32 //体重，kg
	height float32 //身高，m
	bmi    float32 //weight/height^2
}

//计算每个学生的BMI。函数返回后BMI并没有改变，如何把程序改对？
func CalBMI1(students []*Student) {
	for _, student := range students {
		student.bmi = student.weight / (student.height * student.height)
	}
}

func main(){
	stu:=[]*Student{{80.0,178.0,0},{75.0,180.0,0}}
	CalBMI1(stu)
	fmt.Println(stu[0].bmi,stu[1].bmi)
}