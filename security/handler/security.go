package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	security "security/proto"
)

type Security struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Security) Call(ctx context.Context, req *security.Request, rsp *security.Response) error {
	log.Info("Received Security.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Security) Stream(ctx context.Context, req *security.StreamingRequest, stream security.Security_StreamStream) error {
	log.Infof("Received Security.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&security.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Security) PingPong(ctx context.Context, stream security.Security_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&security.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
