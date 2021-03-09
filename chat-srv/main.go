package main

import (
	"chat-srv/handler"
	pb "chat-srv/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("chat-srv"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterChatSrvHandler(srv.Server(), new(handler.ChatSrv))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
