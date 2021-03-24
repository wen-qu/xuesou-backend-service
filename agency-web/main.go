package main

import (
	"github.com/wen-qu/xuesou-backend-service/agency-web/handler"
	pb "github.com/wen-qu/xuesou-backend-service/agency-web/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("agency-web"),
		service.Version("latest"),
	)

	// Register handler
	_ = pb.RegisterAgencyWebHandler(srv.Server(), new(handler.AgencyWeb))

	handler.Init()
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
