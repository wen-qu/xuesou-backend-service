package main

import (
	"github.com/wen-qu/xuesou-backend-service/user-srv/handler"
	pb "github.com/wen-qu/xuesou-backend-service/user-srv/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"

	"github.com/wen-qu/xuesou-backend-service/basic"
)

func main() {

	// init database
	basic.Init()

	// Create service
	srv := service.New(
		service.Name("user-srv"),
		service.Version("latest"),
	)

	// Register handler
	_ = pb.RegisterUserSrvHandler(srv.Server(), new(handler.UserSrv))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
