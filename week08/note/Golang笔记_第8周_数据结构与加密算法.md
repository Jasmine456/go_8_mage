[toc]

# Golang笔记\_第8周\_数据结构与加密算法







## 加密算法

### 对称加密

![image-20220606143616865](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC8%E5%91%A8_%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E4%B8%8E%E5%8A%A0%E5%AF%86%E7%AE%97%E6%B3%95.assets/image-20220606143616865.png)



​		加密过程的每一步都是可逆的。 加密和解密用的是同一组密钥。异或是最简单的堆成加密算法。

```GO
//XOR 异或运算，要求plain和key的长度相同
func XOR(plain string, key []byte) string {
	bPlain := []byte(plain)
	bCipher := make([]byte, len(key))
	for i, k := range key {
		bCipher[i] = k ^ bPlain[i]
	}
	cipher := string(bCipher)
	return cipher
}
```





​		DES(Data Encryption Standard)数据加密标准，是目前最流行的加密算法之一。对原始数据（明文）进行分组，每组64位，最后一组不足64位是按一定规则填充。每一组上单独施加DES算法。



DES子秘钥生成

​		初始秘钥64位，实际有效位56位，每隔7位有一个校验位。根据初始秘钥生成16个48位的子秘钥。

![image-20220606143915932](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC8%E5%91%A8_%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E4%B8%8E%E5%8A%A0%E5%AF%86%E7%AE%97%E6%B3%95.assets/image-20220606143915932.png)

N取值从1到16位，N和x有固定的映射表。



DES加密过程

![image-20220606144009444](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC8%E5%91%A8_%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E4%B8%8E%E5%8A%A0%E5%AF%86%E7%AE%97%E6%B3%95.assets/image-20220606144009444.png)

L1, R1 = f(L0, R0, K1)，依此循环，得到L16和R16。
  S盒替换。输入48位，输出32位。各分为8组，输入每组6位，输出每组4位。分别在每组上施加S盒替换，一共8个S盒。

![image-20220606144118220](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC8%E5%91%A8_%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E4%B8%8E%E5%8A%A0%E5%AF%86%E7%AE%97%E6%B3%95.assets/image-20220606144118220.png)



DES加密过程

![image-20220606144141163](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC8%E5%91%A8_%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E4%B8%8E%E5%8A%A0%E5%AF%86%E7%AE%97%E6%B3%95.assets/image-20220606144141163.png)



​		分组模式。CBC（Cipher Block Chaining ）密文分组链接模式，将当前明文分组与前一个密文分组进行异或运算，然后再进行加密。其他分组模式还有ECB, CTR, CFR, OFB。



```GO
func DesEncryptCBC(text string, key []byte) (string, error) {
	src := []byte(text)
	block, err := des.NewCipher(key) //用des创建一个加密器cipher
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()           //分组的大小，blockSize=8
	src = common.ZeroPadding(src, blockSize) //填充

	out := make([]byte, len(src))                   //密文和明文的长度一致
	encrypter := cipher.NewCBCEncrypter(block, key) //CBC分组模式加密
	encrypter.CryptBlocks(out, src)
	return hex.EncodeToString(out), nil
}

func DesDecryptCBC(text string, key []byte) (string, error) {
	src, err := hex.DecodeString(text) //转成[]byte
	if err != nil {
		return "", err
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}

	out := make([]byte, len(src))                   //密文和明文的长度一致
	encrypter := cipher.NewCBCDecrypter(block, key) //CBC分组模式解密
	encrypter.CryptBlocks(out, src)
	out = common.ZeroUnPadding(out) //反填充
	return string(out), nil
}
```



AES（Advanced Encryption Standard）高级加密标准，旨在取代DES。





### 非对称加密

> - 使用公钥加密，使用私钥解密。
> - 公钥和私钥不同
> - 公钥可以公布给所有人。
> - 私钥只有自己保存
> - 相对于对称加密，运算素服非常慢。



![image-20220606144413228](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC8%E5%91%A8_%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E4%B8%8E%E5%8A%A0%E5%AF%86%E7%AE%97%E6%B3%95.assets/image-20220606144413228.png)

对称加密和非对称加密结合使用的案例。小明要给小红传输机密文件，他俩先交换各自的公钥，然后：

1. 小明生成一个随机的AES口令，然后用小红的公钥通过RSA加密这个口令，并发给小红。
2. 小红用自己的RSA私钥解密得到AES口令。
3. 双方使用这个共享的AES口令用AES加密通信。



![image-20220606144523086](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC8%E5%91%A8_%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E4%B8%8E%E5%8A%A0%E5%AF%86%E7%AE%97%E6%B3%95.assets/image-20220606144523086.png)

​		RSA是三个发明人名字的缩写：Ron Rivest，Adi Shamir，Leonard Adleman。密钥越长，越难破解。 目前768位的密钥还无法破解（至少没人公开宣布）。因此可以认为1024位的RSA密钥基本安全，2048位的密钥极其安全。RSA的算法原理主要用到了数论。
RSA加密过程



1. 随机选择两个不相等的质数p和q。p=61, q=53
2. 计算p和q的乘积n。n=3233
3. 计算n的欧拉函数φ(n) = (p-1)(q-1)。 φ(n) =3120
4. 随机选择一个整数e，使得1< e < φ(n)，且e与φ(n) 互质。e=17
5. 计算e对于φ(n)的模反元素d，即求解e*d+ φ(n)*y=1。d=2753, y=-15
6. 将n和e封装成公钥，n和d封装成私钥。公钥=(3233，17)，公钥=(3233，2753)



```GO
// RSA加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes) //目前的数字证书一般都是基于ITU（国际电信联盟）制定的X.509标准
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// RSA解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
```



 	ECC（Elliptic Curve Cryptography）椭圆曲线加密算法，相比RSA，ECC可以使用更短的密钥，来实现与RSA相当或更高的安全。定义了椭圆曲线上的加法和二倍运算。椭圆曲线依赖的数学难题是：k为正整数，P是椭圆曲线上的点（称为基点）, k*P=Q , 已知Q和P，很难计算出k。



```GO
func genPrivateKey() (*ecies.PrivateKey, error) {
	pubkeyCurve := elliptic.P256() // 初始化椭圆曲线
	// 随机挑选基点,生成私钥
	p, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader) //用golang标准库生成公私钥对
	if err != nil {
		return nil, err
	} else {
		return ecies.ImportECDSA(p), nil //转换成以太坊的公私钥对
	}
}

//ECCEncrypt 椭圆曲线加密
func ECCEncrypt(plain string, pubKey *ecies.PublicKey) ([]byte, error) {
	src := []byte(plain)
	return ecies.Encrypt(rand.Reader, pubKey, src, nil, nil)
}

//ECCDecrypt 椭圆曲线解密
func ECCDecrypt(cipher []byte, prvKey *ecies.PrivateKey) (string, error) {
	if src, err := prvKey.Decrypt(cipher, nil, nil); err != nil {
		return "", err
	} else {
		return string(src), nil
	}
}
```



### 哈希算法

哈希函数的基本特征



1. 输入可以是任意长度。
2. 输出是固定长度。
3. 根据输入很容易计算出输出。
4. 根据输出很难计算出输入（几乎不可能)。
5. 两个不同的输入几乎不可能得到相同的输出



​		SHA（Secure Hash Algorithm） 安全散列算法，是一系列密码散列函数，有多个不同安全等级的版本：SHA-1,SHA-224,SHA-256,SHA-384,SHA-512。其作用是防伪装，防窜扰，保证信息的合法性和完整性



sha-1算法大致过程

1. 填充。使得数据长度对512求余的结果为448。
2. 在信息摘要后面附加64bit，表示原始信息摘要的长度。
3. 初始化h0到h4，每个h都是32位。
4. h0到h4历经80轮复杂的变换。
5. 把h0到h4拼接起来，构成160位，返回。

```GO
func Sha1(data string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(data))
	return hex.EncodeToString(sha1.Sum(nil))
}
```



​		MD5(Message-Digest Algorithm 5)信息-摘要算法 5，算法流程跟SHA-1大体相似。MD5的输出是128位，比SHA-1断了32位。

MD5相对易受密码分析的攻击，运算速度比SHA-1快。



```GO
func Md5(data string) string {
	md5 := md5.New()
	md5.Write([]byte(data))
	return hex.EncodeToString(md5.Sum(nil))
}
```



哈希函数的应用场景

> - 用户密码的存储
> - 文件上传/下载完整性校验
> - mysql大字段的快速对比。数字签名



![image-20220606150301626](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC8%E5%91%A8_%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E4%B8%8E%E5%8A%A0%E5%AF%86%E7%AE%97%E6%B3%95.assets/image-20220606150301626.png)

​		比特币中验证交易记录的真实性用的就是数字签名。先hash 再用私钥加密的原因是： 非对称加密计算量比较大，先hash可以把原始数据转一条很短的信息，提高计算效率。







## 数据结构与算法

### 链表

![image-20220606150737574](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC8%E5%91%A8_%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E4%B8%8E%E5%8A%A0%E5%AF%86%E7%AE%97%E6%B3%95.assets/image-20220606150737574.png)



​		链表的一个应用案例。LRU（Least Recently Used, 最近最少使用)缓存淘汰的总体思路：缓存的key放到链表中，头部的元素表示最近刚使用。



> - 如果命中缓存，从链表中找到对应的key，移到链表头部。
> - 如果没命中缓存：
>   - 如果缓存容量没超，放入缓存，并把key放到链表头部。
>   - 如果超出缓存容量，删除链表尾部元素，再把key放到链表头部。





![image-20220606151647735](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC8%E5%91%A8_%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E4%B8%8E%E5%8A%A0%E5%AF%86%E7%AE%97%E6%B3%95.assets/image-20220606151647735.png)

​		ring的应用：基于滑动窗口的统计。比如最近100次接口调用的平均耗时、最近10笔订单的平均值、最近30个交易日股票的最高点。ring的容量即为滑动窗口的大小，把待观察变量按时间顺序不停的写入ring即可。





```GO
package main

import (
	"container/ring"
	"fmt"
)

func TraverseRing(ring *ring.Ring) {
	ring.Do(func(i interface{}) { //通过Do()来遍历ring，内部实际上调用了Next()而非Prev()
		fmt.Printf("%v ", i)
	})
	fmt.Println()
}

func main() {
	ring := ring.New(5) //必须指定长度，各元素被初始化为nil
	ring2 := ring.Prev()
	for i := 0; i < 3; i++ {
		ring.Value = i
		ring = ring.Next()
	}
	for i := 0; i < 3; i++ {
		ring2.Value = i
		ring2 = ring2.Prev()
	}
	TraverseRing(ring)
	TraverseRing(ring2) //ring和ring2当前所在的指针位置不同，所以遍历出来的顺序也不同
}
```





### 栈

​		栈是一种先进后出的数据结构，push把元素压入栈底，pop弹出栈顶的元素。编程语言的编译系统也用到了栈的思想。



![image-20220606152101464](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC8%E5%91%A8_%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E4%B8%8E%E5%8A%A0%E5%AF%86%E7%AE%97%E6%B3%95.assets/image-20220606152101464.png)



​	go 自带的List已经包含了栈的功能，这里实现了一个线程安全的栈。



```GO
type (
	node struct {
		value interface{}
		prev  *node
	}
	MyStack struct {
		top    *node
		length int
		lock   *sync.RWMutex
	}
)

func NewMyStack() *MyStack {
	return &MyStack{nil, 0, &sync.RWMutex{}}
}

func (stack *MyStack) Push(value interface{}) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	n := &node{value, stack.top}
	stack.top = n
	stack.length++
}

func (stack *MyStack) Pop() interface{} {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	if stack.length == 0 {
		return nil
	}
	n := stack.top
	stack.top = n.prev
	stack.length--
	return n.value
}

func (stack *MyStack) Peak() interface{} {
	stack.lock.RLock()
	defer stack.lock.RUnlock()
	if stack.length == 0 {
		return nil
	}
	return stack.top.value
}

func (stack *MyStack) Len() int {
	return stack.length
}

func (stack *MyStack) Empty() bool {
	return stack.Len() == 0
}
```





### 堆

​		堆是一颗二叉树。大根堆即仁义街店的值都大于等于其子节点。反之为小根堆。

用数组来表示堆，下标为i的结点的父节点下标为(i-1)/2 ,其左右子节点分别为（2i+1）、（2i+2）。

![image-20220606152517736](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC8%E5%91%A8_%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E4%B8%8E%E5%8A%A0%E5%AF%86%E7%AE%97%E6%B3%95.assets/image-20220606152517736.png)



构建堆

![image-20220606152536146](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC8%E5%91%A8_%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E4%B8%8E%E5%8A%A0%E5%AF%86%E7%AE%97%E6%B3%95.assets/image-20220606152536146.png)



```GO
package main

import "fmt"

//AdjustTraingle 如果只是修改slice里的元素，不需要传slice的指针；如果要往slice里append或让slice指向新的子切片，则需要传slice指针
func AdjustTraingle(arr []int, parent int) {
	left := 2*parent + 1
	if left >= len(arr) {
		return
	}

	right := 2*parent + 2
	minIndex := parent
	minValue := arr[minIndex]
	if arr[left] < minValue {
		minValue = arr[left]
		minIndex = left
	}
	if right < len(arr) {
		if arr[right] < minValue {
			minValue = arr[right]
			minIndex = right
		}
	}
	if minIndex != parent {
		arr[minIndex], arr[parent] = arr[parent], arr[minIndex]
		AdjustTraingle(arr, minIndex) //递归。每当有元素调整下来时，要对以它为父节点的三角形区域进行调整
	}
}

func ReverseAdjust(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}
	lastIndex := n / 2 * 2
	fmt.Println(lastIndex)
	for i := lastIndex; i > 0; i -= 2 { //逆序检查每一个三角形区域
		right := i
		parent := (right - 1) / 2
		fmt.Println(parent)
		AdjustTraingle(arr, parent)
	}
}

func buildHeap() {
	arr := []int{62, 40, 20, 30, 15, 10, 49}
	ReverseAdjust(arr)
	fmt.Println(arr)
}

```



​		每当有元素调整下来时，要对以她为父节点的三角形区域进行调整。



插入元素

![image-20220606152657085](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC8%E5%91%A8_%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E4%B8%8E%E5%8A%A0%E5%AF%86%E7%AE%97%E6%B3%95.assets/image-20220606152657085.png)



删除堆顶

![image-20220606152725211](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC8%E5%91%A8_%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E4%B8%8E%E5%8A%A0%E5%AF%86%E7%AE%97%E6%B3%95.assets/image-20220606152725211.png)



下面讲几个堆的应用。

堆排序

1. 构建堆 O(N)
2. 不断地删除堆顶O(NlogN)。



求集合中最大的k个元素

	1. 用集合的前K个元素构建小根堆。
 	2. 注意遍历集合的其他元素，如果比堆顶小直接丢弃；否则替换掉堆顶，然后向下调整堆。



把超时的元素从缓存成中删除

1. 按key的到期时间把key插入小根堆中。
2. 周期扫描堆顶元素，如果它的到期时间早于当前时刻，则从堆和缓存中删除，然后向下调整堆。

golang中的container/heap实现了小根堆，但需要自己定义一个类，实现以下接口：

- Len() int
- Less(i, j int) bool
- Swap(i, j int)
- Push(x interface{})
- Pop() x interface{}



```GO
type Item struct {
	Value    string
	priority int //优先级，数字越大，优先级越高
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority //golang默认提供的是小根堆，而优先队列是大根堆，所以这里要反着定义Less
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

//往slice里append,需要传slice指针
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

//让slice指向新的子切片，需要传slice指针
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]   //数组最后一个元素
	*pq = old[0 : n-1] //去掉最一个元素
	return item
}
```





### Trie树

trie树又叫字典权。
  现有term集合：{分散，分散精力，分散投资，分布式，工程，工程师}，把它们放到Trie树里如下图:

![image-20220606154630165](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC8%E5%91%A8_%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E4%B8%8E%E5%8A%A0%E5%AF%86%E7%AE%97%E6%B3%95.assets/image-20220606154630165.png)





​		Trie树的根节点是总入口，不存储字符。对于英文，第个节点有26个子节点，子节点可以存到数组里；中文由于汉字很多，用数组存子节点太浪费内存，可以用map存子节点。从根节点到叶节点的完整路径是一个term。从根节点到某个中间节点也可能是一个term，即一个term可能是另一个term的前缀。上图中红圈表示从根节点到本节点是一个完整的term。



```GO
package main

import "fmt"

type TrieNode struct {
	Word     rune               //当前节点存储的字符。byte只能表示英文字符，rune可以表示任意字符
	Children map[rune]*TrieNode //孩子节点，用一个map存储
	Term     string
}

type TrieTree struct {
	root *TrieNode
}

//add 把words[beginIndex:]插入到Trie树中
func (node *TrieNode) add(words []rune, term string, beginIndex int) {
	if beginIndex >= len(words) { //words已经遍历完了
		node.Term = term
		return
	}

	if node.Children == nil {
		node.Children = make(map[rune]*TrieNode)
	}

	word := words[beginIndex] //把这个word放到node的子节点中
	if child, exists := node.Children[word]; !exists {
		newNode := &TrieNode{Word: word}
		node.Children[word] = newNode
		newNode.add(words, term, beginIndex+1) //递归
	} else {
		child.add(words, term, beginIndex+1) //递归
	}
}

//walk words[0]就是当前节点上存储的字符，按照words的指引顺着树往下走，最终返回words最后一个字符对应的节点
func (node *TrieNode) walk(words []rune, beginIndex int) *TrieNode {
	if beginIndex == len(words)-1 {
		return node
	}
	beginIndex += 1
	word := words[beginIndex]
	if child, exists := node.Children[word]; exists {
		return child.walk(words, beginIndex)
	} else {
		return nil
	}
}

//traverseTerms 遍历一个node下面所有的term。注意要传数组的指针，才能真正修改这个数组
func (node *TrieNode) traverseTerms(terms *[]string) {
	if len(node.Term) > 0 {
		*terms = append(*terms, node.Term)
	}
	for _, child := range node.Children {
		child.traverseTerms(terms)
	}
}

func (tree *TrieTree) AddTerm(term string) {
	if len(term) <= 1 {
		return
	}
	words := []rune(term)

	if tree.root == nil {
		tree.root = new(TrieNode)
	}

	tree.root.add(words, term, 0)
}

func (tree *TrieTree) Retrieve(prefix string) []string {
	if tree.root == nil || len(tree.root.Children) == 0 {
		return nil
	}
	words := []rune(prefix)
	firstWord := words[0]
	if child, exists := tree.root.Children[firstWord]; exists {
		end := child.walk(words, 0)
		if end == nil {
			return nil
		} else {
			terms := make([]string, 0, 100)
			end.traverseTerms(&terms)
			return terms
		}
	} else {
		return nil
	}
}

func main() {
	tree := new(TrieTree)
	tree.AddTerm("分散")
	tree.AddTerm("分散精力")
	tree.AddTerm("分散投资")
	tree.AddTerm("分布式")
	tree.AddTerm("工程")
	tree.AddTerm("工程师")

	terms := tree.Retrieve("分散")
	fmt.Println(terms)
	terms = tree.Retrieve("人工")
	fmt.Println(terms)
}
```



