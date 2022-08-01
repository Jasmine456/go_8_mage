# Golang笔记\_第11周\_Go语言socket和websocket编程





##  socket编程

### 网络通信过程

![image-20220628094208756](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC11%E5%91%A8_Go%E8%AF%AD%E8%A8%80socket%E5%92%8Cwebsocket%E7%BC%96%E7%A8%8B.assets/image-20220628094208756.png)

> - DMA:网卡和磁盘数据拷贝到内存流程比较固定，不涉及到运算操作，且非常耗时。在磁盘嵌入一个DMA芯片，完成上述拷贝工作，把CPU解脱出来，让CPU专注于运算。
> - mmap: 用户空间和内核空间映射同一块内存空间，从而达到省略将数据从内核缓冲区到用户空间的操作，用户空间通过映射直接操作内核缓冲区的数据。



阻塞式网络I/O

![image-20220628094600126](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC11%E5%91%A8_Go%E8%AF%AD%E8%A8%80socket%E5%92%8Cwebsocket%E7%BC%96%E7%A8%8B.assets/image-20220628094600126.png)



非阻塞式网络I/O

![image-20220628094634197](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC11%E5%91%A8_Go%E8%AF%AD%E8%A8%80socket%E5%92%8Cwebsocket%E7%BC%96%E7%A8%8B.assets/image-20220628094634197.png)



多路复用网络I/O

![image-20220628094739174](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC11%E5%91%A8_Go%E8%AF%AD%E8%A8%80socket%E5%92%8Cwebsocket%E7%BC%96%E7%A8%8B.assets/image-20220628094739174.png)

socket把复杂的传输层协议封装成简单的接口，使应用层可以像读写文件一样进行网络数据的传输。

![image-20220628094855302](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC11%E5%91%A8_Go%E8%AF%AD%E8%A8%80socket%E5%92%8Cwebsocket%E7%BC%96%E7%A8%8B.assets/image-20220628094855302.png)

### socket通信过程

###### socket编程接口参照图

![image-20220628095118153](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC11%E5%91%A8_Go%E8%AF%AD%E8%A8%80socket%E5%92%8Cwebsocket%E7%BC%96%E7%A8%8B.assets/image-20220628095118153.png)





### TCP CS 架构

#### 网络通信模型

OSI 7层参考模型

![image-20220628095547591](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC11%E5%91%A8_Go%E8%AF%AD%E8%A8%80socket%E5%92%8Cwebsocket%E7%BC%96%E7%A8%8B.assets/image-20220628095547591.png)

TCP/IP模型

![image-20220628095621013](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC11%E5%91%A8_Go%E8%AF%AD%E8%A8%80socket%E5%92%8Cwebsocket%E7%BC%96%E7%A8%8B.assets/image-20220628095621013.png)



​		传输层数据大小的上限为MSS(Maximum Segment Size 最大分段大小)，网络接口层的数据大小的上限为MTU(Maximum Transmit Unit 最大传输单元)。



#### TCP协议解读

​		MSS=MTU-ip首部-tcp首部，MTU视网络接口层的不同而不同。TCP在建立连接是通常需要协商双方的MSS值。应用层传输的数据大于MSS时需要分段。

![image-20220628100043599](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC11%E5%91%A8_Go%E8%AF%AD%E8%A8%80socket%E5%92%8Cwebsocket%E7%BC%96%E7%A8%8B.assets/image-20220628100043599.png)



TCP首部

![image-20220628100120120](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC11%E5%91%A8_Go%E8%AF%AD%E8%A8%80socket%E5%92%8Cwebsocket%E7%BC%96%E7%A8%8B.assets/image-20220628100120120.png)



> - 前20个字节是固定的，后面还4N个可选字节（TCP选项）
> - 数据偏移：TCP数据部分距TCP开头的偏移量(一个偏移量是4个字节，TCP选项占4N个字节)，亦即TCP首部的长度。所以TCP首部的最大长度是15*4=60个字节，即TCP选项最多有40个字节。
> - 端口在tcp层指定，ip在IP层指定。端口占2个字节，则最大端口号为2^16-1=65535
> - 由于应用层的数据被分段了，为了在接收端对数据按顺序充足，需要为每段数据编个序号。
> - TCP规定在连接建立后所有传送的报文段都必须把ACK设置为1





TCP建立连接

![image-20220628100641088](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC11%E5%91%A8_Go%E8%AF%AD%E8%A8%80socket%E5%92%8Cwebsocket%E7%BC%96%E7%A8%8B.assets/image-20220628100641088.png)

> - 第一次握手：TCP首部SYN=1,初始化一个序号=J。SYN报文段不能携带数据。
> - 第二次握手：TCP首部SYN=1,ACK=1,确认号=J+1,初始化一个序号=K.此报文同样不携带数据。
> - 第三次握手：SYN=1,ACK=1,序号=j+1,确认号=k+1。此次一般会携带真正需要传输的数据。
> - 确认号：即希望下次对方发过来的序号值。
> - SYN Flood攻击始终不进行第三次握手，属于DDOS攻击的一种。



TCP释放连接

![image-20220628101048263](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC11%E5%91%A8_Go%E8%AF%AD%E8%A8%80socket%E5%92%8Cwebsocket%E7%BC%96%E7%A8%8B.assets/image-20220628101048263.png)

> - TCP的连接是全双工，（可以同时发送和接收）的连接，因此在关闭连接的时候，必须关闭传送和接收两个方向上的连接。
> - 第一次挥手：FIN=1,序号=M
> - 第二次挥手：ACK=1,序号=M+1.
> - 第三次挥手：FIN=1,序号=N.
> - 第四次挥手：ACK=1,序号=N+1
> - 从TIME_WAIT进入CLOSED需要经过2个MSL(Maxinum Segment Lifetime),RFC793建议MSL=2分钟。



#### Go TCP编程

> - 用三元（IP地址，协议，端口号）给唯一标识网络中的一个进程，如（172.122.121.111，tcp,5656).
> - IPv4的地址位数为32位，分为4段，每段最大取值为255
> - IPv6的地址位数为128位，分为8段，各段用16进制表示，最大取值为ffff。
> - 端口0-1023被熟知的应用程序占用（普通应用程序不可以使用），49152-65535客户端程序运行时动态选择使用。



常用函数讲解



```GO
func ResolveTCPAddr(net, addr string) (*TCPAddr, os.Error)
```

​		net参数是"tcp4","tcp6","tcp"中的任意一个。分别表示TCP(IPv4-only)，TCP(IPv6-only)或者TCP(IPv4，IPv6的任意一个)。addr表示域名或者ip地址，例如“www.qq.com:80"或者”127.0.0.1:22“



```GO
func DialTCP(network string, laddr, raddr *TCPAddr) (*TCPConn, error)
```

network参数是"tcp4","tcp6","tcp"中的任意一个。laddr表示本机地址，一般设置为nil。raddr表示远程的服务地址。



```GO
func net.DialTimeout(network string, address string, timeout time.Duration) (net.Conn, error)
```

创建连接时设置的超时时间



```GO
func (*net.conn) Write(b []byte) (int, error)
```

通过conn发送数据



```GO
func (net.Conn).Read(b []byte) (n int, err error)
```

从conn里读取数据，如果没有数据可读，会阻塞。



```GO
func ioutil.ReadAll(r io.Reader) ([]byte, error)
```

从conn中读取所有内容，直到遇到error（比如连接关闭）或EOF



```GO
func ListenTCP(network string, laddr *TCPAddr) (*TCPListener, error)
```

监听端口



```GO
func (l *TCPListener) Accept() (Conn, error)
```

阻塞，直到有客户端请求建立连接



```GO
func (*net.conn) Close() error
```

关闭连接



```GO
func (c *TCPConn) SetReadDeadline(t time.Time) error 
func (c *TCPConn) SetWriteDeadline(t time.Time) error
```

设置从一个http连接上读取和写入的超时时间。



```GO
func (c *TCPConn) SetKeepAlive(keepalive bool) os.Error
```

当一个tcp连接上没有数据时，操作系统会间隔性的发送心跳包，如果长时间没有收到心跳包会任务连接已经断开。



tcp_server.go

```GO
package main

import (
	"encoding/json"
	"fmt"
	"go-course/socket"
	"net"
	"strconv"
	"time"
)

type (
	Request struct {
		A int
		B int
	}
	Response struct {
		Sum int
	}
)

func handleRequest2(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(30 * time.Second)) //30秒后conn.Read会报出i/o timeout
	defer conn.Close()
	for { //长连接，即连接建立后进行多轮的读写交互
		requestBytes := make([]byte, 256) //初始化后byte数组每个元素都是0
		read_len, err := conn.Read(requestBytes)
		if err != nil {
			fmt.Printf("read from socket error: %s\n", err.Error())
			break //到达deadline后，退出for循环，关闭连接。client再用这个连接读写会发生错误
		}
		fmt.Printf("receive request %s\n", string(requestBytes)) //[]byte转string时，0后面的会自动被截掉

		var request socket.Request
		json.Unmarshal(requestBytes[:read_len], &request) //json反序列化时会把0都考虑在内，所以需要指定只读前read_len个字节
		response := socket.Response{Sum: request.A + request.B}

		responseBytes, _ := json.Marshal(response)
		_, err = conn.Write(responseBytes)
		socket.CheckError(err)
		fmt.Printf("write response %s\n", string(responseBytes))
	}
}

//长连接
func main() {
	ip := "127.0.0.1" //ip换成0.0.0.0和空字符串试试
	port := 5656
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ip+":"+strconv.Itoa(port))
	socket.CheckError(err)
	listener, err := net.ListenTCP("tcp4", tcpAddr)
	socket.CheckError(err)
	fmt.Println("waiting for client connection ......")
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		fmt.Printf("establish connection to client %s\n", conn.RemoteAddr().String()) //操作系统会随机给客户端分配一个49152~65535上的端口号
		go handleRequest2(conn)
	}
}
```





### UDP CS 架构

#### UDP协议解读

![image-20220628103204280](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC11%E5%91%A8_Go%E8%AF%AD%E8%A8%80socket%E5%92%8Cwebsocket%E7%BC%96%E7%A8%8B.assets/image-20220628103204280.png)



> - UDP首部占8个字节，所以UDP报文长度最小是8B.
> - 不需要建立连接，直接收发数据，效率很高。
> - 面向报文。
>   - 对应用层交下来的报文，既不合并也不拆分，直接加上边界交给IP层。
>   - TCP是面向字节流
>   - TCP有一个缓冲，当应用程序传送的数据块太长，TCP就可以把它划分短一些再传送；如果应用程序一次只发送一个字节，TCP也可以等待积累有足够多的字节后再构成报文段发送出去。
> - 从机制上不保证顺序（在IP层要对数据分段），可能会丢包（校验和如果出差错就会把这个报文丢弃掉）。在内网环境下分片乱序和数据丢包极少发生。
> - 支持一对一，一对多，多对一和多对多的交互通信。



#### Go UDP编程

```
func net.Dial(network string, address string) (net.Conn, error)
```

network指定为udp，建立udp连接（伪连接）。



```
func net.DialTimeout(network string, address string, timeout time.Duration) (net.Conn, error)
```

network指定为udp，建立连接时指定超时。



```
func net.ResolveUDPAddr(network string, address string) (*net.UDPAddr, error)
```

解析成udp地址



```
func net.ListenUDP(network string, laddr *net.UDPAddr) (*net.UDPConn, error)
```

直接调用Listen就返回一个udp连接



```
func (*net.UDPConn).ReadFromUDP(b []byte) (int, *net.UDPAddr, error)
```

读数据，会返回remote地址



```
func (*net.UDPConn).WriteToUDP(b []byte, addr *net.UDPAddr) (int, error)
```

写数据，需要指定的remote的地址。



udp_server.go

```GO
package main

import (
	"encoding/json"
	"fmt"
	"go-course/socket"
	"net"
	"strconv"
	"time"
)

//长连接
func main() {
	ip := "127.0.0.1" //ip换成0.0.0.0和空字符串试试
	port := 5656
	udpAddr, err := net.ResolveUDPAddr("udp", ip+":"+strconv.Itoa(port))
	socket.CheckError(err)
	conn, err := net.ListenUDP("udp", udpAddr) //UDP不需要创建连接，所以不需要像TCP那样通过Accept()创建连接，这里的conn是个假连接
	socket.CheckError(err)
	conn.SetReadDeadline(time.Now().Add(30 * time.Second))
	defer conn.Close()
	for {
		requestBytes := make([]byte, 256)                           //初始化后byte数组每个元素都是0
		read_len, remoteAddr, err := conn.ReadFromUDP(requestBytes) //一个conn可以对应多个client，ReadFrom可以返回是哪个
		if err != nil {
			fmt.Printf("read from socket error: %s\n", err.Error())
			break //到达deadline后，退出for循环，关闭连接。client再用这个连接读写会发生错误
		}
		fmt.Printf("receive request %s from %s\n", string(requestBytes), remoteAddr.String()) //[]byte转string时，0后面的会自动被截掉

		var request socket.Request
		json.Unmarshal(requestBytes[:read_len], &request) //json反序列化时会把0都考虑在内，所以需要指定只读前read_len个字节
		response := socket.Response{Sum: request.A + request.B}

		responseBytes, _ := json.Marshal(response)
		_, err = conn.WriteToUDP(responseBytes, remoteAddr) //由于UDP conn支持多对多通信，所以通信对方可能有多个EndPoint，通过WriteTo指定要写给哪个EndPoint
		socket.CheckError(err)
		fmt.Printf("write response %s to %s\n", string(responseBytes), remoteAddr.String())
	}
}
```



udp_client.go

```GO
package main

import (
	"encoding/json"
	"fmt"
	"go-course/socket"
	"net"
	"strconv"
	"sync"
	"time"
)

//长连接
func main() {
	ip := "127.0.0.1" //ip换成0.0.0.0和空字符串试试
	port := 5656
	//跟tcp_client的唯一区别就是这行代码
	conn, err := net.DialTimeout("udp", ip+":"+strconv.Itoa(port), 30*time.Minute) //一个conn绑定一个本地端口
	socket.CheckError(err)
	defer conn.Close()
	const P = 10
	wg := sync.WaitGroup{}
	wg.Add(P)
	for i := 0; i < P; i++ {
		request := socket.Request{A: 7, B: 4}
		requestBytes, _ := json.Marshal(request)
		go func() { //多协程，共用一个conn
			defer wg.Done()
			for { //长连接，即连接建立后进行多轮的读写交互
				_, err = conn.Write(requestBytes)
				socket.CheckError(err)
				fmt.Printf("write request %s\n", string(requestBytes))
				responseBytes := make([]byte, 256) //初始化后byte数组每个元素都是0
				read_len, err := conn.Read(responseBytes)
				socket.CheckError(err)
				var response socket.Response
				json.Unmarshal(responseBytes[:read_len], &response) //json反序列化时会把0都考虑在内，所以需要指定只读前read_len个字节
				fmt.Printf("receive response: %d\n", response.Sum)
				time.Sleep(1 * time.Second)
			}
		}()
	}
	wg.Wait()
}
```



​		由于UDP不需要建立连接，所以通过Dial()创建的是一个虚拟连接，Dial() 总是会返回成功，即使对方还没有准备好。所以UDP可以开启client，再启server。由于是虚拟连接所以多个client可以共用一个conn，所以Server端往conn里写数据时需要指定写给那个client，同理从conn里读数据会返回client的Address，即WriteToUDP(b []byte,addr *net.UDPAddr)和ReadFromUDP(b []byte) (int, *net.UDPAddr,error).由于UDP是无连接，和对方关闭连接后，本方再在conn上调用Write和Read不会报错。



​		应用层的一条完整数据称为报文。TCP是面向字节流的，一次Read到的数据可能包含了多个报文，也可能只包含了半个报文，一条报文在什么地方结束需要通信双方事先约定好。 UDP是面向报文的，一次Read只读一个保温，如果没有把一个报文读完，后面的内容会被丢弃掉，下次就读不到了。







## WebSocket编程

#### WebSocket协议解读

![image-20220628114253065](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC11%E5%91%A8_Go%E8%AF%AD%E8%A8%80socket%E5%92%8Cwebsocket%E7%BC%96%E7%A8%8B.assets/image-20220628114253065.png)



websocket和http协议的关联：

> - 都是应用层协议，都基于tcp传输协议
> - 跟http有良好的的兼容性，ws和http的默认端口都是80，wss和https的默认端口都是443.
> - websocket在握手阶段采用http发送数据。



webdocket和http协议的差异：

> - http是半双工的，而websocket通过多路复用实现了全双工
> - http只能有client竹筒发起数据请求，而websocket还可以有server主动向client推送数据。在需要及时刷新的场景中，http只能靠client高频的轮询，浪费严重。
> - http是短连接（也可以是实现长连接，HTTP1.1的链接默认使用长连接），每次数据请求都得经过三次握手重新建立连接，而websocket是长连接。
> - http长连接中每次请求都要带上header，而websocket在传输数据阶段不需要带header。



WebSocket是HTML5的产物，能更好的节省服务器资源和带宽，websocket应用场景举例：

> - html5多人游戏
> - 聊天室
> - 协同编辑
> - 基于实时位置的应用
> - 股票实时报价
> - 弹幕
> - 视频会议



webdocket握手协议：

Request Header

```GO
Sec-Websocket-Version:13
Upgrade:websocket
Connection:Upgrade
Sec-Websocket-Key:duR0pUQxNgBJsRQKj2Jxsw==
```



Response Header

```GO
Upgrade:websocket
Connection:Upgrade
Sec-Websocket-Accept:a1y2oy1zvgHsVyHMx+hZ1AYrEHI=
```



> - Upgrade:websocket 和Connection：Upgrade知名使用WebSocket协议
> - Sec-WebSocket-Version指定Websocket协议版本
> - sec-WebSocket-Key是一个Base64 encode的值，是浏览器随机生成的。
> - 服务端收到Sec-WebSocket-Key后拼接上一个固定的GUID，进行一次SHA-1摘要，再转成Base64编码，得到Sec-WebSocket-Accept 返回给客户端。客户端对本地的Sec-WebSocket-Key执行同样的操作跟服务端返回的结果进行对比，如果不一直会返回错误关闭连接。如此操作是为了把websocket header 跟 http header区分开。



#### WebSocket CS架构实现

​		首先需要安装gorilla的webdocket包。

```GO
go get github.com/gorilla/websocket
```



1.将http升级到WebSocket协议



```
func (u *Upgrader) Upgrade(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*websocket.Conn, error)
```



2.客户端发起握手，请求建立连接

```
func (*websocket.Dialer) Dial(urlStr string, requestHeader http.Header) (*websocket.Conn, *http.Response, error)
```



3.基于connection进行read和write

ws_server.go

```GO
package main

import (
	"fmt"
	"go-course/socket"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

type WsServer struct {
	listener net.Listener
	addr     string
	upgrade  *websocket.Upgrader
}

func NewWsServer(port int) *WsServer {
	ws := new(WsServer)
	ws.addr = "0.0.0.0:" + strconv.Itoa(port)
	ws.upgrade = &websocket.Upgrader{
		HandshakeTimeout: 5 * time.Second, //握手超时时间
		ReadBufferSize:   2048,            //读缓冲大小
		WriteBufferSize:  1024,            //写缓冲大小
		//请求检查函数，用于统一的链接检查，以防止跨站点请求伪造。如果Origin请求头存在且原始主机不等于请求主机头，则返回false
		CheckOrigin: func(r *http.Request) bool {
			fmt.Printf("request url %s\n", r.URL)
			fmt.Println("handshake request header")
			for key, values := range r.Header {
				fmt.Printf("%s:%s\n", key, values[0])
			}
			return true
		},
		//http错误响应函数
		Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {},
	}
	return ws
}

//httpHandler必须实现ServeHTTP接口
func (ws *WsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/add" {
		fmt.Println("path error")
		http.Error(w, "请求的路径不存在", 222) //把出错的话术写到ResponseWriter里
		return
	}
	conn, err := ws.upgrade.Upgrade(w, r, nil) //将http协议升级到websocket协议
	if err != nil {
		fmt.Printf("upgrade http to websocket error: %v\n", err)
		return
	}
	fmt.Printf("establish conection to client %s\n", conn.RemoteAddr().String())
	go ws.handleConnection(conn)
}

//处理连接里发来的请求数据
func (ws *WsServer) handleConnection(conn *websocket.Conn) {
	defer func() {
		conn.Close()
	}()
	for { //长连接
		conn.SetReadDeadline(time.Now().Add(20 * time.Second))
		var request socket.Request
		if err := conn.ReadJSON(&request); err != nil {
			//判断是不是超时
			if netError, ok := err.(net.Error); ok { //如果ok==true，说明类型断言成功
				if netError.Timeout() {
					fmt.Printf("read message timeout, remote %s\n", conn.RemoteAddr().String())
					return
				}
			}
			//忽略websocket.CloseGoingAway/websocket.CloseNormalClosure这2种closeErr，如果是其他closeErr就打一条错误日志
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
				fmt.Printf("read message from %s error %v\n", conn.RemoteAddr().String(), err)
			}
			return //只要ReadMessage发生错误，就关闭这条连接
		} else {
			response := socket.Response{Sum: request.A + request.B}
			if err = conn.WriteJSON(&response); err != nil {
				fmt.Printf("write response failed: %v", err)
			} else {
				fmt.Printf("write response %d\n", response.Sum)
			}
		}
	}
}

func (ws *WsServer) Start() (err error) {
	ws.listener, err = net.Listen("tcp", ws.addr) //http和websocket都是建立在tcp之上的
	if err != nil {
		fmt.Printf("listen error:%s\n", err)
		return
	}
	err = http.Serve(ws.listener, ws) //开始对外提供http服务。可以接收很多连接请求，其他一个连接处理出错了，也不会影响其他连接
	if err != nil {
		fmt.Printf("http server error: %v\n", err)
		return
	}

	// if err:=http.ListenAndServe(ws.addr, ws);err!=nil{	//Listen和Serve两步合成一步
	// 	fmt.Printf("http server error: %v\n", err)
	// 	return
	// }
	return nil
}

func main() {
	ws := NewWsServer(5657)
	ws.Start()
}
```



ws_client.go

```GO
package main

import (
	"encoding/json"
	"fmt"
	"go-course/socket"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	dialer := &websocket.Dialer{}
	header := http.Header{
		"Cookie": []string{"name=zcy"},
	}
	conn, resp, err := dialer.Dial("ws://localhost:5657/add", header) //Dial:握手阶段，会发送一条http请求。请求一个不存在的路径试试看
	defer resp.Body.Close()
	if err != nil {
		fmt.Printf("dial server error:%v\n", err)
		fmt.Println(resp.StatusCode)
		msg, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(msg))
		return
	}
	fmt.Println("handshake response header")
	for key, values := range resp.Header {
		fmt.Printf("%s:%s\n", key, values[0])
	}
	// time.Sleep(5 * time.Second)
	defer conn.Close()
	for i := 0; i < 10; i++ {
		request := socket.Request{A: 7, B: 4}
		requestBytes, _ := json.Marshal(request)
		err = conn.WriteJSON(request) //websocket.Conn直接提供发json序列化和反序列化方法
		socket.CheckError(err)
		fmt.Printf("write request %s\n", string(requestBytes))
		var response socket.Response
		err = conn.ReadJSON(&response)
		socket.CheckError(err)
		fmt.Printf("receive response: %d\n", response.Sum)
		time.Sleep(1 * time.Second)
	}
	time.Sleep(30 * time.Second)
}
```





​		websocket发送的消息类型有5中：

- TextMessag

- BinaryMessage

- CloseMessage

- PingMessage

- PongMessage

  ​		TextMessag和BinaryMessage费别表示发送文本消息和二进制信息。 CloseMessage关闭帧，接收方收到这个消息就关闭链接 PingMessage和PongMessage时保持心跳的帧，发送方接收方是PingMessage，接收方发送方是PongMessage，目前浏览器没有相关api发送ping给服务器，只能有服务器发ping给浏览器，浏览器返回pong信息。



#### 聊天室实现

​		gorilla的websocket项目中有一个聊天室的demo，此处讲一下它的设计思路。

官方代码：



总体架构如下图所示

![image-20220628140358225](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC11%E5%91%A8_Go%E8%AF%AD%E8%A8%80socket%E5%92%8Cwebsocket%E7%BC%96%E7%A8%8B.assets/image-20220628140358225.png)



Hub

> - Hub持有每一个Client的指针，broadcast管道里有数据是把它写入每一个client的send管道中
> - 注销Client是关闭Client的send管道



Client

> - 前端（browser）请求建立websocket连接时，为这条websocket连接专门开启一个协程，创建一个client。
> - client把前端发过来的数据写入hub的broadcast管道。
> - client把自身send管道里的数据写给前端。
> - client跟前端的连接断开时请求从hub那儿注销自己。





Front

> - 当打开浏览器页面时，前端会请求建立websocket连接。
> - 关闭浏览器页面时会主动关闭websocket连接。





存活监测

> - 当hub发现client的send管道写不进数据时，把client注销掉。
> - client给websocket连接设置一个读超时，并周期性地给前端发ping消息，如果没有收到pong消息则下一次的conn.read()会报出超时错误，此时client关闭websocket连接。





