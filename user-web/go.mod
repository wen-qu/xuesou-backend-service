module github.com/wen-qu/xuesou-backend-service/user-web

go 1.13

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.8.3
	github.com/micro/micro/v3 v3.0.0
	github.com/nats-io/nats-server/v2 v2.1.9 // indirect
	github.com/wen-qu/xuesou-backend-service/user-srv v0.0.0-00010101000000-000000000000 // indirect
)

replace github.com/wen-qu/xuesou-backend-service/user-srv => ../user-srv
