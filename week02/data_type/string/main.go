package main

import (
	"fmt"
	"strconv"
	"strings"
)

func fuzhi() {
	s1 := "sdfs sdf sdfsdf 中加 有"           // 字符串里可以包含任意Unicode字条
	s2 := "dsfsdf '\"sdf\tsdfdsf\n \\中加 右" //包含转义字符
	s3 := `sdfdsf '"sdf dsfdsf
\中加 右"'` // 反引号里的转义字符无效，反引号里的原封不动的输出，包括空白符和换行符
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}

func string_function() {

	s1 := "中"
	fmt.Println(len(s1))
	var arr []byte = []byte(s1)
	fmt.Println(arr)
	s := "abcfsd fsd中国人"
	fmt.Println(len(s))                      // 1个英文字母的长度为1,1个汉字占3个长度
	fmt.Println(strings.Split(s, " "))       //以xxx為分隔符，返回值是一个数组！
	fmt.Println(strings.Contains(s, "abc"))  //包含子串，返回值是bool
	fmt.Println(strings.Contains(s, "fs"))   //包含子串，返回值是bool
	fmt.Println(strings.Index(s, "abc"))     //寻找子串第一次出现的位置
	fmt.Println(strings.LastIndex(s, "abc")) //寻找子串最后一次出现的位置
	fmt.Println(strings.HasPrefix(s, "abc")) //以 xxx结尾，返回值是bool
	fmt.Println(strings.HasSuffix(s, "abc")) //以xxx 开头，返回值是bool

}

//字符串拼接
func concat() {
	s1 := "abc"
	s2 := "efg"
	s3 := "中2q"
	space := " "

	s4 := s1 + s2 + s3
	fmt.Println(s4)

	s5 := fmt.Sprintf("%s%s%s", s1, s2, s3)
	fmt.Println(s5)

	arr := []string{s1, s2, s3}
	s6 := strings.Join(arr, "")
	fmt.Println(s6)
	//当有大量的string需要拼接时，用strings.Builder效率最高
	sb := strings.Builder{}
	sb.WriteString(s1)
	sb.WriteString(space)
	sb.WriteString(s2)
	sb.WriteString(space)
	sb.WriteString(s3)
	s7 := sb.String()
	fmt.Println(s7)

}

func string_diceng() {
	s := "My name is 焦明"
	arr := []byte(s) //强制类型转换
	brr := []rune(s)

	fmt.Printf("last byte %d\n", arr[len(arr)-1]) //string可以转换为[]byte或[]rune类型
	fmt.Printf("last byte %c\n", arr[len(arr)-1]) // byte或rune可以转为string
	fmt.Printf("last rune %d\n", arr[len(brr)-1])
	fmt.Printf("last rune %c\n", arr[len(brr)-1])
	L := len(s)
	fmt.Printf("string len %d byte array len %d rune array len %d\n", L, len(arr), len(brr))
	for _, ele := range s {
		fmt.Printf("%c", ele) // string中的每个元素是字符
	}
	fmt.Println()
	for i := 0; i < L; i++ {
		fmt.Printf("%c", s[i]) //[i] 前面应该出现数组或切片，这里自动把string转成了[]byte（而不是[]rune）
	}

}

func string_other_convert() {
	var err error
	var i int = 8
	var i64 int64 = int64(i)
	//	int转string
	var s string = strconv.Itoa(i)
	s = strconv.FormatInt(i64, 10)
	//	string转int
	i, err = strconv.Atoi(s)
	//string转int64
	i64, err = strconv.ParseInt(s, 10, 32)

	//	float转string
	var f float64 = 8.123456789
	s = strconv.FormatFloat(f, 'f', 2, 64) //保留两位小数
	fmt.Println(s)
	//	string转float
	f, err = strconv.ParseFloat(s, 64)

	//string<-->[]byte
	var arr []byte = []byte(s)
	s = string(arr)

	//string<-->[]rune
	var brr []rune = []rune(s)
	s = string(brr)

	fmt.Printf("err %v\n", err)
}

func main() {
	//fuzhi()
	//string_function()
	//concat()
	//string_diceng()
	string_other_convert()
}
