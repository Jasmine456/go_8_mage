package main

/*
2. 对一个任意长度的byte切片，在尾部追加任意元素使切片的长度变成8的整位数。这个过程称为padding。
要求拿到padding后的切片能够还原出原始的切片。
```go
func Padding(originSlice []byte) []byte { }
func UnPadding(paddedSlice []byte) []byte { }
```

參考：
1. https://blog.csdn.net/xz_studying/article/details/94229023
2. 本代碼week08/encryption/common/padding.go
3. 其他同学的实现

func padding(input []byte, x byte) (output []byte) {
	inputLen := len(input)
	num := inputLen % 8
	switch num {
	case 0:
		output = append(input, []byte{x, x, x, x, x, x, x, 8}...)
	case 1:
		output = append(input, []byte{x, x, x, x, x, x, 7}...)
	case 2:
		output = append(input, []byte{x, x, x, x, x, 6}...)
	case 3:
		output = append(input, []byte{x, x, x, x, 5}...)
	case 4:
		output = append(input, []byte{x, x, x, 4}...)
	case 5:
		output = append(input, []byte{x, x, 3}...)
	case 6:
		output = append(input, []byte{x, 2}...)
	case 7:
		output = append(input, []byte{1}...)

	}
	return
}
func unPadding(output []byte) (input []byte) {
	outputLen := len(output)
	num := output[outputLen-1]
	input = output[:outputLen-int(num)]
	return
}
 */

func Padding(originSlice []byte) []byte {

	return []byte
}
func UnPadding(paddedSlice []byte) []byte {

	return []byte
}