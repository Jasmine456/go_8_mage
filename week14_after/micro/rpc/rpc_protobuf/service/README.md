
cd service
protoc -I=. --go_out=./ --go_opt=module="go_8_mage/week14_after/micro/rpc/rpc_protobuf/service" pb/hello.proto
