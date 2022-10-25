# grpc hello world

在当前hello目录执行：
# protoc-gen-go
# --go_out=.pb --go_opt=module="go_8_mage/week14_after/micro/rpc/grpc/hello/pb"
# protoc-gen-go-grpc

protoc -I=. --go_out=./pb --go_opt=module="go_8_mage/week14_after/micro/rpc/grpc/hello/pb" \
--go-grpc_out=./pb --go-grpc_opt=module="go_8_mage/week14_after/micro/rpc/grpc/hello/pb" \
pb/hello.proto