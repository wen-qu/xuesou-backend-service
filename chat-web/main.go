package main

import (
	"chat-web/handler"
	pb "chat-web/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("chat-web"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterChatWebHandler(srv.Server(), new(handler.ChatWeb))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
