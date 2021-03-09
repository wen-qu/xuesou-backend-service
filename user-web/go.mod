module user-web

go 1.13

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.8.3
	github.com/micro/micro/v3 v3.0.0
	github.com/wen-qu/xuesou-backend-service/basic v1.2.3 // indirect
	github.com/wen-qu/xuesou-backend-service/user-srv v1.2.3
)

replace github.com/wen-qu/xuesou-backend-service/basic => ../basic

replace github.com/wen-qu/xuesou-backend-service/user-srv => ../user-srv
