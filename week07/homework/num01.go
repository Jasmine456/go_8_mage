package main

/*
1. 比较两个byte切片是否相同
```go
func bsEqual(arr, brr []byte) bool {}
```
 */

func BsEqual(arr,brr []byte)bool{
	if &arr == &brr{
		return true
	}else {
		return false
	}
}

func BsEqual2(arr,brr []byte)bool{
	arr_len:=len(arr)
	brr_len:=len(brr)

	if arr_len != brr_len{
		return false
	}else {
		for i:=0;i<arr_len;i++{
			if arr[i] !=brr[i]{
				return false
			}
		}
	}
	return true
}

