package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	agencysrv "agency-srv/proto"
)

type AgencySrv struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *AgencySrv) Call(ctx context.Context, req *agencysrv.Request, rsp *agencysrv.Response) error {
	log.Info("Received AgencySrv.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *AgencySrv) Stream(ctx context.Context, req *agencysrv.StreamingRequest, stream agencysrv.AgencySrv_StreamStream) error {
	log.Infof("Received AgencySrv.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&agencysrv.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *AgencySrv) PingPong(ctx context.Context, stream agencysrv.AgencySrv_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&agencysrv.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
