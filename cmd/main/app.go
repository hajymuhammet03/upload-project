package main

import (
	"context"
	"fmt"
	handlermanager "github.com/Hajymuhammet03/internal/handlers/manager"
	"github.com/Hajymuhammet03/pkg/config"
	"github.com/Hajymuhammet03/pkg/logging"
	"github.com/Hajymuhammet03/pkg/postgresql"
	"github.com/jackc/pgx/v4/pgxpool"

	"net"
	"net/http"

	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	cfg := config.GetConfig()
	logger := logging.GetLogger()

	postgresSQLClient, err := postgresql.NewClient(context.TODO(), 3, cfg.Storage)
	if err != nil {
		logger.Fatalf("%v", err)
	}
	start(handlermanager.Manager(postgresSQLClient, logger), cfg, postgresSQLClient)
}

func start(router *mux.Router, cfg *config.Config, pGPool *pgxpool.Pool) {
	logger := logging.GetLogger()
	logger.Info("start application")

	//go autorun.AutoListen(pGPool)
	var listener net.Listener
	var listenErr error

	logger.Info("listen tcp")
	listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
	logger.Infof("server is listening port %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)

	if listenErr != nil {
		logger.Fatal(listenErr)
	}

	fileServer := http.FileServer(http.Dir("/home/user/shared/public/"))
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fileServer))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders: []string{
			"*",
		},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})
	handler := c.Handler(router)

	server := &http.Server{
		Handler:        handler,
		WriteTimeout:   10000 * time.Second,
		ReadTimeout:    10000 * time.Second,
		MaxHeaderBytes: 1 << 20 * 10 * 1000 * 5000000,
	}

	logger.Fatal(server.Serve(listener))

}
