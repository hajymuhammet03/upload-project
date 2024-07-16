package category

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Hajymuhammet03/internal/appresult"
	"github.com/Hajymuhammet03/internal/handlers"
	"github.com/Hajymuhammet03/pkg/logging"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

const (
	addCategory = "/category"
)

type handler struct {
	logger     *logging.Logger
	repository Repository
}

func NewHandler(repository Repository, logger *logging.Logger) handlers.Handler {
	return &handler{
		repository: repository,
		logger:     logger,
	}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc(addCategory, appresult.Middleware(h.AddCategory)).Methods(http.MethodPost)
}

func (h *handler) AddCategory(w http.ResponseWriter, r *http.Request) error {
	body, errBody := ioutil.ReadAll(r.Body)
	if errBody != nil {
		fmt.Println("Category in handler ErrBody:", errBody)
		return appresult.ErrMissingParam
	}

	category := AddCategory{}
	errData := json.Unmarshal(body, &category)
	if errData != nil {
		fmt.Println("Category in handler ErrData:", errData)
		return appresult.ErrMissingParam
	}

	data, err := h.repository.AddCategory(context.TODO(), category)
	if err != nil {
		return appresult.ErrInternalServer
	}

	successResult := appresult.Success
	successResult.Data = data

	w.Header().Add(appresult.HeaderContentTypeJson())
	err = json.NewEncoder(w).Encode(successResult)
	if err != nil {
		return err
	}
	return nil
}
