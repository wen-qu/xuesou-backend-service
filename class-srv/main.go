package main

import (
	"class-srv/handler"
	pb "class-srv/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("class-srv"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterClassSrvHandler(srv.Server(), new(handler.ClassSrv))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
