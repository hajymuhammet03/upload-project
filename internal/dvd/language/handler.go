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
	getLanguageID  = "/get-language/{id}"
	deleteLanguage = "/delete-language/{id}"
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
	router.HandleFunc(getLanguage, appresult.Middleware(h.GetLanguage)).Methods(http.MethodGet)
	router.HandleFunc(getLanguageID, appresult.Middleware(h.GetLanguageID)).Methods(http.MethodGet)
	router.HandleFunc(deleteLanguage, appresult.Middleware(h.DeleteLanguage)).Methods(http.MethodDelete)

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
		fmt.Println("error handler in Add Language body: ", errBody)
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

// Get Language v1
// @Summary Get language
// @Description Get a language based on a search query parameter
// @Tags dvd
// @ID get_language_v1
// @Accept json
// @Produce json
// @Param search query string true "Search query parameter"
// @Success 200 {object} []language.Language "Successful operation"
// @Failure 500 {object} string	"some err from db"
// @Router /get-language [get]
func (h *handler) GetLanguage(w http.ResponseWriter, r *http.Request) error {
	search := r.URL.Query().Get("search")

	data, err := h.repository.GetLanguage(context.TODO(), search)
	if err != nil {
		fmt.Println("error handler in GetLanguage: ", err)
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

// GetLanguageID v1
// @Summary Get language by ID
// @Description Get details of a specific language by its ID
// @Tags dvd
// @Param id path string true "Language ID"
// @Produce json
// @Success 200 {object} language.UUID "Successful operation"
// @Failure 500 {object} string	"some err from db"
// @Router /get-language/{id} [get]
func (h *handler) GetLanguageID(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	data, err := h.repository.GetLanguageID(context.TODO(), id)
	if err != nil {
		fmt.Println("error handler in GetLanguageID: ", err)
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

// DeleteLanguage v1
// @Summary Delete language by ID
// @Description Delete a specific language by its ID
// @Tags dvd
// @Param id path string true "Language ID"
// @Produce json
// @Success 200
// @Failure 500 {object} string	"some err from db"
// @Router /delete-language/{id} [delete]
func (h *handler) DeleteLanguage(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	err := h.repository.DeleteLanguage(context.TODO(), id)
	if err != nil {
		fmt.Println("error handler in DeleteLanguage: ", err)
		return appresult.ErrInternalServer
	}
	successResult := appresult.Success
	successResult.Data = ""
	w.Header().Add(appresult.HeaderContentTypeJson())
	err = json.NewEncoder(w).Encode(successResult)
	if err != nil {
		return err
	}
	return nil
}
