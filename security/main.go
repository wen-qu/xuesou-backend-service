package main

import (
	"github.com/wen-qu/xuesou-backend-service/security/handler"
	pb "github.com/wen-qu/xuesou-backend-service/security/proto"
	"net/http"

	"github.com/micro/micro/v3/plugin"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("security"),
		service.Version("latest"),
	)

	// Register handler
	_ = pb.RegisterSecurityHandler(srv.Server(), new(handler.Security))
	_ = plugin.Register(plugin.NewPlugin(
		plugin.WithName("auth"),
		plugin.WithHandler(checkToken),
		))

	handler.Init()
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}

func checkToken(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		whiteList := []string{"/agency-web/login", "/agency-web/register", "/user-web/login", "/user-web/register"}
		for _, url := range whiteList {
			if r.URL.Path == url {
				h.ServeHTTP(w, r)
				return
			}
		}

		token := r.Header.Get("Authorization")
		if token[:6] == "Bearer" {
			token = token[7:] // delete the "Bearer " prefix
		}

		acc, err := handler.JWTClient.Inspect(token)
		if err != nil {
			_, _ = w.Write([]byte("unauthorized"))
			return
		}

		if acc != nil {
			h.ServeHTTP(w, r)
		}
	})

}