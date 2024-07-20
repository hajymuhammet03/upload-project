package manager

import (
	"fmt"
	"github.com/Hajymuhammet03/internal/dvd/category"
	categorydb "github.com/Hajymuhammet03/internal/dvd/category/db"
	"github.com/Hajymuhammet03/internal/dvd/film_category"
	filmCategorydb "github.com/Hajymuhammet03/internal/dvd/film_category/db"
	"github.com/Hajymuhammet03/internal/dvd/language"
	languagedb "github.com/Hajymuhammet03/internal/dvd/language/db"
	"github.com/Hajymuhammet03/pkg/logging"
	"github.com/Hajymuhammet03/pkg/postgresql"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	baseURL        = "/api/v1/dvd"
	healthcheckURL = "/api/v1/healthcheck"
)

func Manager(db postgresql.Client, logger *logging.Logger) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc(healthcheckURL, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status": "our news service is running"}`)
	})

	categoryRouterManager := router.PathPrefix(baseURL).Subrouter()
	categoryRouterRepository := categorydb.NewRepository(db, logger)
	categoryRouterHandler := category.NewHandler(categoryRouterRepository, logger)
	categoryRouterHandler.Register(categoryRouterManager)

	filmCategoryRouterManager := router.PathPrefix(baseURL).Subrouter()
	filmCategoryRouterRepository := filmCategorydb.NewRepository(db, logger)
	filmCategoryRouterHandler := film_category.NewHandler(filmCategoryRouterRepository, logger)
	filmCategoryRouterHandler.Register(filmCategoryRouterManager)

	languageRouterManager := router.PathPrefix(baseURL).Subrouter()
	languageRouterRepository := languagedb.NewRepository(db, logger)
	languageRouterHandler := language.NewHandler(languageRouterRepository, logger)
	languageRouterHandler.Register(languageRouterManager)

	return router
}
