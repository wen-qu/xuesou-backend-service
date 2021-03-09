package main

import (
	"agency-srv/handler"
	pb "agency-srv/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("agency-srv"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterAgencySrvHandler(srv.Server(), new(handler.AgencySrv))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
