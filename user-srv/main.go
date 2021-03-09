package main

import (
	"user-srv/handler"
	pb "user-srv/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"

	"github.com/wen-qu/xuesou-backend-service/basic"
)

func main() {

	// init database
	basic.Init()

	// Create service
	srv := service.New(
		service.Name("go.micro.user.srv"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterUserSrvHandler(srv.Server(), new(handler.UserSrv))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
