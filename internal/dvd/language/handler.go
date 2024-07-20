package language

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
	addLanguage    = "/language"
	getLanguage    = "/get-language"
	getLanguageID  = "/get-language-id"
	deleteLanguage = "/delete-language"
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
	router.HandleFunc(addLanguage, appresult.Middleware(h.AddLanguage)).Methods(http.MethodPost)
}

// Add Language v1
// @Summary v1
// @Description post add language
// @Tags dvd
// @ID add_language_v1
// @Produce json
// @Param LanguageDTO body LanguageDTO true "Language DTO"
// @Success 200 {object} language.UUID
// @Failure 500	{object} string	"some err from db"
// @Router /language [post]
func (h *handler) AddLanguage(w http.ResponseWriter, r *http.Request) error {
	body, errBody := ioutil.ReadAll(r.Body)
	if errBody != nil {
		fmt.Println("error handler in AddLanguage body: ", errBody)
		return appresult.ErrMissingParam
	}

	language := LanguageDTO{}
	errData := json.Unmarshal(body, &language)
	if errData != nil {
		fmt.Println("error handler in AddLanguage data: ", errData)
		return appresult.ErrMissingParam
	}

	data, err := h.repository.AddLanguage(context.TODO(), language)
	if err != nil {
		fmt.Println("error handler in AddLanguage: ", err)
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
