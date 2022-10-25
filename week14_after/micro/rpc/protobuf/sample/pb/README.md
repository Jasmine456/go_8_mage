# 编译sample.proto文件

必须在sample目录下执行
protoc.exe -I=. --go_out=./pb --go_opt=module="go_8_mage/week14_after/micro/rpc/protobuf/sample/pb" pb/enum.proto


protoc.exe -I=. --go_out=./pb --go_opt=module="go_8_mage/week14_after/micro/rpc/protobuf/sample/pb" pb/array.proto


protoc.exe -I=. --go_out=./pb --go_opt=module="go_8_mage/week14_after/micro/rpc/protobuf/sample/pb" pb/oneof.proto

protoc.exe -I=. -I=/usr/local/include --go_out=./pb --go_opt=module="go_8_mage/week14_after/micro/rpc/protobuf/sample/pb" pb/any.proto

protoc.exe -I=. -I=/usr/local/include --go_out=./other_pb --go_opt=module="go_8_mage/week14_after/micro/rpc/protobuf/sample/other_pb" other_pb/other.proto
