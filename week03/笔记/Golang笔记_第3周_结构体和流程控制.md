# Golang笔记\_第三周\_结构体和流程控制





## 目录

[TOC]

> - if语句
> - switch语句
> - for循环
> - break与 continue
> - goto语句与Label





### if语句



#### 常用示例

```
if 5> 9 {
	fmt.Println("5>9")
}
```



> - 如果逻辑表达式成立，就会执行{}里的内容
> - 逻辑表达式不需要加（）
> - “{”必须跟在逻辑表达式后面，不能另起一行



#### 复杂的逻辑表达式 if

```GO
if c,d,e := 5,9,2;c<d && (c>e||c>3){//初始化多个局部变量，复杂的逻辑表达式
    fmt.Println("fit") 
}
```



> - 逻辑表达式中可以含有变量或常量
> - if句子中允许包含1个（仅1个）分号，在分号前初始化一些局部变量(即只在if块内可见)



#### if规范

if-else代码演示

![image-20220422143407754](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC3%E5%91%A8_%E7%BB%93%E6%9E%84%E4%BD%93%E5%92%8C%E6%B5%81%E7%A8%8B%E6%8E%A7%E5%88%B6.assets/image-20220422143407754.png)





if嵌套最好不要超过 3层，规避多层嵌套



#### 多层if更换方式代码示例

![image-20220422143717052](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC3%E5%91%A8_%E7%BB%93%E6%9E%84%E4%BD%93%E5%92%8C%E6%B5%81%E7%A8%8B%E6%8E%A7%E5%88%B6.assets/image-20220422143717052.png)



#### if map的值不存在代码示例

![image-20220422150136474](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC3%E5%91%A8_%E7%BB%93%E6%9E%84%E4%BD%93%E5%92%8C%E6%B5%81%E7%A8%8B%E6%8E%A7%E5%88%B6.assets/image-20220422150136474.png)



#### if代码实操---公交车道

![image-20220422150044379](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC3%E5%91%A8_%E7%BB%93%E6%9E%84%E4%BD%93%E5%92%8C%E6%B5%81%E7%A8%8B%E6%8E%A7%E5%88%B6.assets/image-20220422150044379.png)









### switch语句

#### 语法示例

![image-20220422150232452](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC3%E5%91%A8_%E7%BB%93%E6%9E%84%E4%BD%93%E5%92%8C%E6%B5%81%E7%A8%8B%E6%8E%A7%E5%88%B6.assets/image-20220422150232452.png)





#### 规范

> - switch-case-default可能模拟 if-else if-else,但只能实现相等判断
> - switch和case后面可以跟常量、变量或函数表达式，只要它们表示的数据类型相同就行
> - case后面可以跟多个值，只要有一个值满足就行







#### 空的switch

> - switch后带表达式时，switch-case只能模拟相等的情况；如果switch后不带表达式，case后就可以跟任意的条件表达式。



```GO
switch{
    case add(5)>10;
    	fmt.Println("right")
    default:
    	fmt.Println("wrong")
}
```







#### switch Type

![image-20220422150706624](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC3%E5%91%A8_%E7%BB%93%E6%9E%84%E4%BD%93%E5%92%8C%E6%B5%81%E7%A8%8B%E6%8E%A7%E5%88%B6.assets/image-20220422150706624.png)