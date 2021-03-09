package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	usersrv "user-srv/proto"
)

type UserSrv struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *UserSrv) Call(ctx context.Context, req *usersrv.Request, rsp *usersrv.Response) error {
	log.Info("Received UserSrv.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *UserSrv) Stream(ctx context.Context, req *usersrv.StreamingRequest, stream usersrv.UserSrv_StreamStream) error {
	log.Infof("Received UserSrv.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&usersrv.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *UserSrv) PingPong(ctx context.Context, stream usersrv.UserSrv_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&usersrv.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
