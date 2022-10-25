# 编译helo.proto文件

protoc.exe -I=. --go_out=./pb --go_opt=module="go_8_mage/week14_after/micro/rpc/protobuf/hello/pb" pb/hello.proto
