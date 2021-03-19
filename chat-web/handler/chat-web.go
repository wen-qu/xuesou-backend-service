package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	chatweb "chat-web/proto"
)

type ChatWeb struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *ChatWeb) Call(ctx context.Context, req *chatweb.Request, rsp *chatweb.Response) error {
	log.Info("Received ChatWeb.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *ChatWeb) Stream(ctx context.Context, req *chatweb.StreamingRequest, stream chatweb.ChatWeb_StreamStream) error {
	log.Infof("Received ChatWeb.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&chatweb.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *ChatWeb) PingPong(ctx context.Context, stream chatweb.ChatWeb_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&chatweb.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
