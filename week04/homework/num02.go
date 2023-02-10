package main

import "fmt"

/*
实现一个函数，接受若干个float64（用不定长参数），返回这些参数乘积的倒数，除数为0时返回error
上题用递归实现

3 1/3
f(a)=1/a
5 1/5
f(b)=1/b
8 1/8
3,5 1/(3*5) 1/3 * 1/5
f(a,b)=f(a)*f(b)

3,5,8 1/(3*5*8)  1/3 * 1/5 * 1/8
f(a,b,c)=f(a)*f(b)*f(c)=f(a,b)*f(c)=f(a)f(b,c)
*/
var count int=0
func printIndent(n int ){
	for i:=0;i<n;i++{
		fmt.Printf("  ")
	}
}
func f(args ...float64)(float64,error){
	count+=count

	if len(args)==1{
		if args[0]==0{
			return 0,fmt.Errorf("divide by 0")
		}
		return 1/args[0],nil
	}
	if len(args)==0{
		return 1,nil
	}
	first:=args[0]
	fmt.Println(first)
	remain:=args[1:]//截取子切片
	printIndent(count)
	fmt.Println(remain)
	part1,err1:=f(first)
	if err1!=nil{
		return 0,err1
	}
	part2,err2:=f(remain...)//把切片转为不定长参O数

	if err2!=nil{
		return 0,err2
	}
	result:=part1*part2
	return result,nil
}


func main(){
	f(1.0,2.0,3.0,4.0)
}