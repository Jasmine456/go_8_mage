package main

import (
	"bufio"
	"fmt"
	exec "golang.org/x/sys/execabs"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func format() {
	var i int = 1234
	var f float32 = 3.1415
	var stu struct {
		Name string
		Age  int
	}
	stu.Name = "张三"
	stu.Age = 18
	fmt.Printf("%b\n", i) //二进制表示
	fmt.Printf("%d\n", i)
	fmt.Printf("%16d\n", i) //左边补空格，不够8位
	fmt.Printf("%08d\n", i) //左边补0，补够8位
	fmt.Printf("%f\n", f)   //默认6位小数
	fmt.Printf("%.8f\n", f)
	fmt.Printf("%e\n", f) //科学记数法，默认6位小数
	fmt.Printf("%.8e\n", f)
	fmt.Printf("%g\n", f)     //根据实际情况采用%e 或%f格式(获得更简洁、更准确的输出)
	fmt.Printf("%t\n", 3 > 9) //true或false
	fmt.Printf("%s\n", stu.Name)

	fmt.Printf("%T\n", stu)
	fmt.Printf("%v\n", stu)
	fmt.Printf("%+v\n", stu)
	fmt.Printf("%#v\n", stu)

}

//从标准输入读入数据
func scan() {
	fmt.Println("Please input a word")
	var word string
	fmt.Scan(&word) //读入第1个空格前的单词
	fmt.Println(word)

	fmt.Println("Please input two word")
	var word1 string
	var word2 string
	fmt.Scan(&word1, &word2) //读入多个单词，空格分隔。如果输入了更多单词会被缓存起来，丢给下一次scan
	fmt.Println(word1, word2)

	fmt.Println("Please input an int")
	var i int
	fmt.Scanf("%d", &i) //类似于scan，转为特定格式的数据
	fmt.Println(i)
}

func read_file() {
	file_path := "./io/data/digit.txt"
	if fin, err := os.Open(file_path); err != nil {
		fmt.Printf("read file failed: %v\n", err) //比如文件不存在
	} else {
		defer fin.Close() //别忘了关闭文件句柄

		//	读二进制文件
		cont := make([]byte, 10)
		//if n,err := fin.Read(cont);err != nil{ //读出len(cont)个字节，返回成功读取的字节数
		//	fmt.Printf("read file failed: %v\n",err)
		//}else{
		//	fmt.Println(string(cont[:n]))
		//	if m,err := fin.ReadAt(cont,int64(n));err !=nil{
		//		fmt.Printf("read file failed :%v\n",err)
		//	}else{
		//		fmt.Println(string(cont[:m]))
		//	}
		//	fin.Seek(int64(n),0) //whence：0从文件开头计算偏移量，1从当前位置计算偏移量，2到文件末尾的偏移量
		//	if n,err = fin.Read(cont);err !=nil{
		//		fmt.Printf("read file failed:%v\n",err)
		//	}else{
		//		fmt.Println(string(cont[:n]))
		//	}
		//}
		for {
			n, err := fin.Read(cont)
			if err != nil {
				if err == io.EOF { //已经读到文件末尾
					break
				} else {
					fmt.Println(err)
				}
			} else {
				fmt.Print(string(cont[0:n]))
			}
		}
	}

}

func read_file2() {
	file_path := "io/data/digit.txt"
	if fin, err := os.Open(file_path); err != nil {
		fmt.Printf("open file failed:%v\n", err)
	} else {
		defer fin.Close()
		//	读文本文件建议用bufio.Reader
		fin.Seek(0, 0) //定位到文件开头
		reader := bufio.NewReader(fin)
		for {
			if line, err := reader.ReadString('\n'); err != nil { //指定分隔符
				if err == io.EOF {
					if len(line) > 0 { //如果最后一行没有换行符，则此时最后一行就存在line里
						fmt.Println(line)
					}
					break //已读到文件末尾
				} else {
					fmt.Printf("read file failed %v\n", err)
				}
			} else {
				line = strings.TrimRight(line, "\n")
				fmt.Println(line)
			}
		}
	}
}

func write_file() {
	//	OpenFile()比Open()有更多的参数选项，os.O_WRONLY以止泻的方式打开文件，os.O_TRUNC把文件之前的内容先清空掉，
	//	os.O_CREATE如果文件不存在先创建，0666新建文件的权限设置
	file_name := "io/data/write01.txt"
	if fout, err := os.OpenFile(file_name, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666); err != nil {
		fmt.Printf("open file failed: %v\n", err)
	} else {
		defer fout.Close() //别忘了关闭文件句柄

		//	写文本文件建议使用 bufio
		writer := bufio.NewWriter(fout)
		writer.WriteString("明月多情应笑我")
		writer.WriteString("\n")
		writer.WriteString("笑我如今")
		writer.Flush() //buffer中的数据量累积到一定程度后才会真正写入磁盘。调用Flush强行把缓冲中的所有内容写入磁盘
	}
}

func create_file() {
	file_name := "./io/data/write01.txt" //先删除，不去理会Remove可能返回的error
	os.Remove(file_name)
	if file, err := os.Create(file_name); err != nil {
		fmt.Printf("open file failed: %v\n", err)
	} else {
		file.Chmod(0666)                 //设置文件权限
		fmt.Printf("fd=%d\n", file.Fd()) //获取文件描述符file descriptor，这是一个整数
		info, _ := file.Stat()
		fmt.Printf("is dir %t\n", info.IsDir())
		fmt.Printf("modify time %s\n", info.ModTime())
		fmt.Printf("mode %v\n", info.Mode())
		fmt.Printf("file name %s\n", info.Name())
		fmt.Printf("size %d\n", info.Size())
		fmt.Printf("sys %v\n", info.Sys())
	}

	os.Mkdir("data/sys", os.ModePerm)          //创建目录并设置权限
	os.MkdirAll("data/sys/a/b/c", os.ModePerm) //增强版Mkdir，沿途的目录不存在时会一并创建

	os.Rename("data/sys/a", "data/sys/p")       // 给文件或目录重命名
	os.Rename("data/sys/p/b/c", "data/sys/p/c") //Rename 还可以实现move的功能

	err := os.Remove("data/sys") //删除文件或目录，目录不为空时才能删除成功
	fmt.Println(err)
	os.RemoveAll("data/sys") //增强版Remove，所有子目录会递归删除
}

//遍历一个目录
func walk(path string) error {
	if fileInfos, err := ioutil.ReadDir(path); err != nil {
		return err
	} else {
		for _, fileInfo := range fileInfos {
			fmt.Println(fileInfo.Name())
			if fileInfo.IsDir() {
				if err := walk(filepath.Join(path, fileInfo.Name())); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func logger(){
	log.Printf("%d+%d=%d\n",3,4,3+4)
	log.Println("Hello Golang")
	//log.Fatalln("Bye,the world") //日志输出后会执行os.Exit(1)

//	以append方式打开日志文件
	fout,err := os.OpenFile("io/data/test.log",os.O_CREATE|os.O_APPEND|os.O_WRONLY,0666)
	if err != nil{
		fmt.Printf("open log file failed: %v\n",err)
	}
	defer fout.Close()
	logWriter := log.New(fout,"[MY_BIZ]",log.Ldate|log.Lmicroseconds)
	logWriter.Printf("%d+%d=%d\n",3,4,3+4)
	logWriter.Println("hello Golang")
}

//执行系统命令
func sys_call(){
//	查看系统命令所在的目录，确保命令已安装
	cmd_path,err := exec.LookPath("dir")
	if err !=nil{
		fmt.Println("could not found command echo")
	}
	fmt.Printf("command echo in path %s\n",cmd_path) // /bin/df
	cmd := exec.Command("df","-h")// 相当于命令df -h,注意Command的每一个参数都不能包含空格

	//cmd.Output() 运行命令并获得其输出结果
	if output,err:= cmd.Output();err !=nil{
		fmt.Println("got output failed",err)
	} else{
		fmt.Println(string(output))
	}

	cmd = exec.Command("rm","./data/test.log")
	// 如果不需要获得命令的输出，直接调用cmd.Run()即可
	if err !=nil{
		fmt.Println("run failed",err)
	}
}



func main() {
	//format()
	//scan()
	//read_file()
	//read_file2()
	//write_file()
	//create_file()
	//walk("data")
	//logger()
	sys_call()

}
