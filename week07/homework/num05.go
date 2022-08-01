package main

import "fmt"

/*

 */

func BikeScream() string {
	return "叮铃铃"
}

func CarScream() string {
	return "嘟嘟嘟"
}

func TrainScream() string {
	return "呜呜呜"
}

//根据传入的参数不同，Run()动态地执行不同的行为。通过接口怎么实现类似的功能？
func Run(foo func() string) {
	fmt.Println("Are you ready?")
	fmt.Println(foo())
}

//接口实现
type Param interface {
	Run()string
}

type Car struct {

}
type Bike struct {

}
type Train struct {

}
func (c *Car)Run()string{
	fmt.Println("Are you ready?\n嘟嘟嘟")
	return "嘟嘟嘟"
}
func (b *Bike)Run()string{
	fmt.Println("Are you ready?\n叮铃铃")
	return "叮铃铃"
}
func(t *Train)Run()string{
	fmt.Println("Are you ready?\n呜呜呜")
	return "呜呜呜"
}


func main(){
	Run(TrainScream)
	//c:=new(Car)
	//b:=new(Bike)
	//t:=new(Train)
	//p:=new(Param)
	//p=c
	//var pa Param
	//pa=c
	//c.Run()
	//b.Run()
	//t.Run()
	//pa.Run()
	c:=&Car{}
	b:=&Bike{}
	t:=&Train{}
	var p Param
	p=c
	p.Run()
	p=b
	p.Run()
	p=t
	p.Run()

}
