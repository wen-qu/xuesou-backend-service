module github.com/wen-qu/xuesou-backend-service/agency-web

go 1.15

require (
	github.com/wen-qu/xuesou-backend-service/agency-srv v0.0.0-20210331132935-3de36a15f6ef
	github.com/wen-qu/xuesou-backend-service/class-srv v0.0.0-20210325152105-ea9f2abd9285

)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
