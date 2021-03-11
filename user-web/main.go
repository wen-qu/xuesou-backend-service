package main

import (
	"github.com/wen-qu/xuesou-backend-service/user-web/handler"
	pb "github.com/wen-qu/xuesou-backend-service/user-web/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("user-web"),
		service.Version("latest"),
	)

	// Register handler
	_ = pb.RegisterUserWebHandler(srv.Server(), new(handler.UserWeb))

	handler.Init()
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
