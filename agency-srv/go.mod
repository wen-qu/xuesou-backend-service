module github.com/wen-qu/xuesou-backend-service/agency-srv

go 1.15

require (
	github.com/micro/micro/v3 v3.0.0
	github.com/wen-qu/xuesou-backend-service/basic v0.0.0-20210311040543-ea030d0fdf8b
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
