module github.com/wen-qu/xuesou-backend-service/agency-srv

go 1.15

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/gobwas/httphead v0.0.0-20180130184737-2c6c146eadee // indirect
	github.com/gobwas/pool v0.2.1 // indirect
	github.com/gobwas/ws v1.0.3 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.2
	github.com/jinzhu/copier v0.2.8
	github.com/micro/micro/v3 v3.2.0
	github.com/soheilhy/cmux v0.1.4 // indirect
	github.com/wen-qu/xuesou-backend-service/basic v0.0.0-20210404124158-591e4a74fbb0
	google.golang.org/protobuf v1.25.0
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
