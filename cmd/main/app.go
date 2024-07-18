package main

import (
	"context"
	"fmt"
	handlermanager "github.com/Hajymuhammet03/internal/handlers/manager"
	"github.com/Hajymuhammet03/pkg/config"
	"github.com/Hajymuhammet03/pkg/logging"
	"github.com/Hajymuhammet03/pkg/postgresql"
	"github.com/jackc/pgx/v4/pgxpool"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/Hajymuhammet03/docs"
	"net"
	"net/http"

	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// @title Swagger API FILM
// @version 1.0
// @description This is upload project
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:2303
// @BasePath /api/v1/dvd
// @schemes http https

// @securityDefinitions.apiKey  ApiKeyAuth
// @security
// @in header
// @name authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
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
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:2303/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	server := &http.Server{
		Handler:        handler,
		WriteTimeout:   10000 * time.Second,
		ReadTimeout:    10000 * time.Second,
		MaxHeaderBytes: 1 << 20 * 10 * 1000 * 5000000,
	}

	logger.Fatal(server.Serve(listener))

}
