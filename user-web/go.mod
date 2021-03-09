module github.com/wen-qu/xuesou-backend-service/user-web

go 1.16

require (
	github.com/coreos/etcd v3.3.18+incompatible // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/micro/cli v0.2.0
	github.com/micro/micro/v3 v3.0.0
	github.com/mwitkow/go-conntrack v0.0.0-20190716064945-2f068394615f // indirect
	github.com/nats-io/nats-server/v2 v2.1.9 // indirect
	github.com/wen-qu/xuesou-backend-service/user-srv v1.2.3
)

replace github.com/wen-qu/xuesou-backend-service/user-srv => ../user-srv

replace github.com/wen-qu/xuesou-backend-service/basic => ../basic

replace github.com/coreos/etcd => github.com/coreos/etcd v3.3.10+incompatible
