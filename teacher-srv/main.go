package main

import (
	"teacher-srv/handler"
	pb "teacher-srv/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("teacher-srv"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterTeacherSrvHandler(srv.Server(), new(handler.TeacherSrv))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
