package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	classsrv "class-srv/proto"
)

type ClassSrv struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *ClassSrv) Call(ctx context.Context, req *classsrv.Request, rsp *classsrv.Response) error {
	log.Info("Received ClassSrv.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *ClassSrv) Stream(ctx context.Context, req *classsrv.StreamingRequest, stream classsrv.ClassSrv_StreamStream) error {
	log.Infof("Received ClassSrv.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&classsrv.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *ClassSrv) PingPong(ctx context.Context, stream classsrv.ClassSrv_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&classsrv.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
