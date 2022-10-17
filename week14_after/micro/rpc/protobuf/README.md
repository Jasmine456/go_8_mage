# 编译hello.proto文件

```GO
protoc.exe -I=. --go_out=./pb --go_opt=module="go_8_mage/week14_after/micro/rpc/protobuf/hello/pb" pb/hello.proto
```