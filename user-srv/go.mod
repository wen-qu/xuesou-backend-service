module github.com/wen-qu/xuesou-backend-service/user-srv

go 1.16

require (
	github.com/golang/protobuf v1.4.3
	github.com/micro/go-micro v1.18.0 // indirect
	github.com/micro/micro/v3 v3.0.0
	github.com/wen-qu/xuesou-backend-service/basic v0.0.0-00010101000000-000000000000
	google.golang.org/protobuf v1.25.0
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/wen-qu/xuesou-backend-service/basic => ../basic
