package main

import (
	"github.com/wen-qu/xuesou-backend-service/security/handler"
	pb "github.com/wen-qu/xuesou-backend-service/security/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("security"),
		service.Version("latest"),
	)

	// Register handler
	_ = pb.RegisterSecurityHandler(srv.Server(), new(handler.Security))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
