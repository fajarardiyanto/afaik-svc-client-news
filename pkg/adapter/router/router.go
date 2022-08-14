package router

import (
	"github.com/fajarardiyanto/afaik-svc-client-news/app/models"
	"github.com/fajarardiyanto/afaik-svc-client-news/app/service"
	"github.com/fajarardiyanto/afaik-svc-client-news/internal/config"
	"github.com/fajarardiyanto/afaik-svc-client-news/pkg/adapter/middleware"
	"github.com/fajarardiyanto/afaik-svc-client-news/pkg/connection"
	"github.com/fajarardiyanto/flt-go-listener/lib/client"
	"github.com/gorilla/mux"
)

func Router() {
	r := mux.NewRouter()

	r.Use(middleware.Middleware)

	conn := connection.NewConnection(config.GetConfig().Client).InitConn()

	// repo
	newsRepo := models.NewNewsRepository(conn)

	newsSpv := service.NewNewsClient(config.GetLogger(), newsRepo)

	r.HandleFunc("/get-news", newsSpv.Get).Methods("POST")

	client.NewServer(config.GetLogger(), config.GetConfig().Server, r).StartHTTPServer()
}
