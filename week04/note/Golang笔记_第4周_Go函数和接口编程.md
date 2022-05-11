# Golang笔记\_第4周\_Go函数编程和接口编程





[toc]

## 函数

### 函数的基本形式

```go
// 函数定义 a b是形参
func argf(a int,b int){
    a = a+b
}
var x ,y int =3,6
argf(x,y)// 函数调用，x，y是实参
```



> - 形参是函数内部的局部变量，实参的值会拷贝给形参
> - 参数定义时的第一个大括号不能另起一行
> - 形参可以有0个或多个
> - 参数类型相同时可以只写一次，比如argf(a,b int)
> - 在函数内部修改形参的值，实参的值不受影响
> - 如果想通过函数修改实参，就需要指针类型



```GO
func argf(a,b *int){
    *a=*a+*b
    *b = 888
}
var x,y int = 3,6
argf(&x,&y)
```



slice 、map、channel都是引用类型，它们作为函数参数是其实跟普通struct没什么区别，都是对struct内部的各个字段做一次拷贝传到函数内部。

```GO
package main

import "fmt"

func slice_arg_1(arr []int){
    
}
```









