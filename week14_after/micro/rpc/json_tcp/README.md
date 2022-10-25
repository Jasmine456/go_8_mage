# json on tcp

沿用之前rpc_interface里的interface包，接下来采用json编码

#测试
由于没有合适的tcp工具, 比如nc, 同学可以下来自己验证

$ echo -e '{"method":"HelloService.Hello","params":["hello"],"id":1}' | nc localhost 1234
{"id":1,"result":"hello:hello","error":null}