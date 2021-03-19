package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	agencyweb "agency-web/proto"
)

type AgencyWeb struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *AgencyWeb) Call(ctx context.Context, req *agencyweb.Request, rsp *agencyweb.Response) error {
	log.Info("Received AgencyWeb.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *AgencyWeb) Stream(ctx context.Context, req *agencyweb.StreamingRequest, stream agencyweb.AgencyWeb_StreamStream) error {
	log.Infof("Received AgencyWeb.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&agencyweb.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *AgencyWeb) PingPong(ctx context.Context, stream agencyweb.AgencyWeb_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&agencyweb.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
