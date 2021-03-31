package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	teachersrv "teacher-srv/proto"
)

type TeacherSrv struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *TeacherSrv) Call(ctx context.Context, req *teachersrv.Request, rsp *teachersrv.Response) error {
	log.Info("Received TeacherSrv.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *TeacherSrv) Stream(ctx context.Context, req *teachersrv.StreamingRequest, stream teachersrv.TeacherSrv_StreamStream) error {
	log.Infof("Received TeacherSrv.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&teachersrv.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *TeacherSrv) PingPong(ctx context.Context, stream teachersrv.TeacherSrv_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&teachersrv.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}