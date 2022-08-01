# Golang笔记\_第13周\_项目篇 搜索引擎

# 亲手开发高性能搜索引擎
## 构建倒排索引
&#8195;&#8195;准备数据。data/doc.txt是原始的待索引的文档(document)，一行代表一个user的信息，用逗号分成5列，分别表示uid,keywords,gender,degree,city。形如  
```
729325696,立功|劳动法|开除|员工|规定|请假,女,本科,重庆
```
&#8195;&#8195;一次性把该文件里的内容写入数据库，后续建索引时直接从数据库读数据。操作数据库使用gorm，能批量操作就不要逐条操作。  
正排索引  

![image-20220720142220168](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC13%E5%91%A8_%E9%A1%B9%E7%9B%AE%E7%AF%87%20%E6%90%9C%E7%B4%A2%E5%BC%95%E6%93%8E.assets/image-20220720142220168.png)





正排索引的实现     

|  存储器  |   内部实现    | 语言 |                         适用场景                          |
| :------: | :-----------: | :--: | :-------------------------------------------------------: |
|   Bolt   |     B+树      |  go  |                 读多写少，范围扫描。etcd                  |
| RocksDB  |      LSM      | C++  | 顺序读写时LSM性能优势突显；更改默认配置可实现快速的随机读 |
|  Badger  |      LSM      |  go  |     为随机读做了专门的优化，号称至少比RocksDB快3.5倍      |
| TerarkDB | Succinct Tire | C++  |                超高压缩率，超高随机读性能                 |

RocksDB  
- RocksDB（Facebook）大量复用了LevelDB（Google）的代码，并且还借鉴了许多HBase的设计理念。原始代码从leveldb 1.5 上fork出来
- K,V是任意长度的字节流
- 针对不同的存储硬件（内存,Flash,硬盘,HDFS）做了专门的优化，以使性能达到最优。性能随CPU线性提升
- 支持各种压缩算法，并提供了便捷的生产环境维护和调试工具
- 只能单机部署。支持ReadOnly模式  

Badger  
- 纯go编写，针对SSD（固态硬盘）做了专门优化
- Key存储在.sst文件中，Value存储在.vlog文件中
- WithMaxCacheSize：内存中缓存多少数据
- WithCompression：压缩算法
- WithEncryptionKey：加密磁盘上的数据  

倒排索引  
![image-20220720142240709](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC13%E5%91%A8_%E9%A1%B9%E7%9B%AE%E7%AF%87%20%E6%90%9C%E7%B4%A2%E5%BC%95%E6%93%8E.assets/image-20220720142240709.png)



倒排索引的实现  
- 一个大Map。不支持并发读写
- 一个大sync.Map。一种高效的线程安全的Map
- 多个小Map。写非常少的情况下，自行分段加读写锁，效率比sync.Map更高  

遍历倒排索引   
- 每次遍历倒排链都要创建一个临时容器（显然是切片而非数组），容纳遍历的结果
- 如果要往slice或map里添加大量元素，则创建slice或map时最好指定capacity，避免内存频繁拷贝，产生大量的临时内存，导致多次GC
- 当需要频繁申请大量的临时内存时，可以使用sync.Pool  

并行遍历倒排索引   
- 对于特别长的倒排链，开多个协程，分段并行遍历
- 多协程遍历的结果放到同一个slice里，slice/struct支持多协程并发修改，map不支持。因为slice中的每个元素都有确定的地址（头指针加偏移量），而map中某个key下的value其地址是不固定的（在重哈希阶段）

索引的更新   
- 先拿docid从正排索引上取得doc对应的keys，再从这些keys对应的倒排链上把docid删掉
- 遍历倒排链时加读锁，更新倒排链时加写锁
- 更新的QPS不能太高，否则会影响读的耗时
- 为避免写索引的QPS突增，把所有写请求放到一个缓冲队列（channel）里去，异步执行
## 筛选功能的实现
用key前缀表示Field。KW前缀表示keywords，CT前缀表示城市。   
![image-20220720142304520](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC13%E5%91%A8_%E9%A1%B9%E7%9B%AE%E7%AF%87%20%E6%90%9C%E7%B4%A2%E5%BC%95%E6%93%8E.assets/image-20220720142304520.png)

key中组合多个field。常用的筛选字段放到key里。  
![image-20220720142313325](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC13%E5%91%A8_%E9%A1%B9%E7%9B%AE%E7%AF%87%20%E6%90%9C%E7%B4%A2%E5%BC%95%E6%93%8E.assets/image-20220720142313325.png)

key中组合多个field。只能组合少量特征，防止key爆炸。  
![image-20220720142324260](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC13%E5%91%A8_%E9%A1%B9%E7%9B%AE%E7%AF%87%20%E6%90%9C%E7%B4%A2%E5%BC%95%E6%93%8E.assets/image-20220720142324260.png) 

构造Must筛选条件  
- query：关键词=算法|工程师，城市=北京
- 构造key：
    1. 关键词和城市分别构造key
{KW算法，KW工程师},{CT北京}
    2. 求上述两个集合的笛卡尔积
{CT北京_KW算法，CT北京_KW工程师}
    3. 取上述2条倒排链的交集  

构造Should筛选条件  
- query：关键词=“算法|工程师”或“算法|经理”，城市=北京
- 构造key：
    1. must，所有key必须同时命中。 {CT北京_KW算法}
    2. should，命中任意一个key即可。{CT北京_KW工程师,CT北京_KW经理}
    3. must和should结果求交集

bits筛选  
- 倒排key只能完成一两个field的条件筛选
- 当field是离散属性且取值较少时，可以用bit标识
- 连续属性也可以进行离散化  

用bit表示离散取值：   
![image-20220720142342078](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC13%E5%91%A8_%E9%A1%B9%E7%9B%AE%E7%AF%87%20%E6%90%9C%E7%B4%A2%E5%BC%95%E6%93%8E.assets/image-20220720142342078.png)



离散属性有多个取值：   
![image-20220720142414860](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC13%E5%91%A8_%E9%A1%B9%E7%9B%AE%E7%AF%87%20%E6%90%9C%E7%B4%A2%E5%BC%95%E6%93%8E.assets/image-20220720142414860.png)



连续属性离散化：  
![image-20220720142436595](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC13%E5%91%A8_%E9%A1%B9%E7%9B%AE%E7%AF%87%20%E6%90%9C%E7%B4%A2%E5%BC%95%E6%93%8E.assets/image-20220720142436595.png)



通过bits只是完成了初步筛选，最后还需要拿着doc详情进行精确筛选。  
bits附在docid后面：  
![image-20220720142446873](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC13%E5%91%A8_%E9%A1%B9%E7%9B%AE%E7%AF%87%20%E6%90%9C%E7%B4%A2%E5%BC%95%E6%93%8E.assets/image-20220720142446873.png)



## 分布式索引
水平切分索引  
- 优势：遍历单条倒排链很快
- 劣势：倒排链跨机器求交集，逻辑复杂；正排索引冗余地存储在多台机器上
![image-20220720142504983](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC13%E5%91%A8_%E9%A1%B9%E7%9B%AE%E7%AF%87%20%E6%90%9C%E7%B4%A2%E5%BC%95%E6%93%8E.assets/image-20220720142504983.png)



垂直切分索引  
- 优势：某一台宕机后对搜索结果影响不大；每一台的搜索行为与单机检索时完全相同
- 劣势：倒排链长度分配不均，链查询速度由最慢的那台机器决定
![image-20220720142519513](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC13%E5%91%A8_%E9%A1%B9%E7%9B%AE%E7%AF%87%20%E6%90%9C%E7%B4%A2%E5%BC%95%E6%93%8E.assets/image-20220720142519513.png)



分布式索引  
![image-20220720142532781](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC13%E5%91%A8_%E9%A1%B9%E7%9B%AE%E7%AF%87%20%E6%90%9C%E7%B4%A2%E5%BC%95%E6%93%8E.assets/image-20220720142532781.png)

 

- 由多个Group垂直切分整个倒排索引，每个Group内有多台Server做冗余备份
- Server之间使用socket通信
- Group内部使用最小并发度算法做负载均衡

## 提供搜索服务
RPC远程过程调用
![image-20220720142550471](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC13%E5%91%A8_%E9%A1%B9%E7%9B%AE%E7%AF%87%20%E6%90%9C%E7%B4%A2%E5%BC%95%E6%93%8E.assets/image-20220720142550471.png)



- 序列化协议：json(HTTP), protobuf(GRPC), thrift
- 服务注册：zookeeper, Consul, etcd
- 负载均衡：加权轮询法，最小并发度法
- 通信：TCP, epoll

## 性能监控
在线profile代码：
```Go
import (
	"net/http"
	_ "net/http/pprof"
)
func main() {
        // 在8080端口上接收profile请求，这样可在程序运行期间对其进行性能监控
        go http.ListenAndServe("localhost:8080", nil) 
}
```

终端交互  
- go tool pprof `http://127.0.0.1:8080/debug/pprof/profile`  //采集CPU使用信息
- go tool pprof `http://localhost:8080/debug/pprof/heap` //采集内存信息
- 常用命令。top -n列出最耗CPU或内存的前n个函数。list <func>列出函数中每一行消耗CPU或内存的情况  

web界面  
- 直接在游览器中访问`http://127.0.0.1:8080/debug/pprof/`
- 运行go tool pprof -http=:8081 `http://localhost:8080/debug/pprof/profile`，然后在游览器中访问`http://localhost:8081/ui/`