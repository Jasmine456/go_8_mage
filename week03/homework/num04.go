package main

import (
	"fmt"
	"math/rand"
)

/*
4.随机初始化两个8*5的矩阵，求两个矩阵的和（逐元素相加）

??? 怎么解
 */

func main(){
	const SIZE1 =8
	const SIZE2 =5
	A := [SIZE1][SIZE2]int{}
	B := [SIZE1][SIZE2]int{}
	for i:=0;i<SIZE1;i++{
		for j:=0;j<SIZE2;j++{
			A[i][j]=rand.Intn(10)
			B[i][j]=rand.Intn(10)
		}
	}

	C:=[SIZE1][SIZE2]int{}
	for i:=0;i<SIZE1;i++{
		for j:=0;j<SIZE2;j++{
			C[i][j]=A[i][j]+B[i][j]
		}
	}
	fmt.Println(C)




}
