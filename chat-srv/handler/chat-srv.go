package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	chatsrv "chat-srv/proto"
)

type ChatSrv struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *ChatSrv) Call(ctx context.Context, req *chatsrv.Request, rsp *chatsrv.Response) error {
	log.Info("Received ChatSrv.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *ChatSrv) Stream(ctx context.Context, req *chatsrv.StreamingRequest, stream chatsrv.ChatSrv_StreamStream) error {
	log.Infof("Received ChatSrv.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&chatsrv.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *ChatSrv) PingPong(ctx context.Context, stream chatsrv.ChatSrv_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&chatsrv.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
