package film_category

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
	addFilmCategory    = "/film-category"
	getFilmCategory    = "/get-film-category"
	getFilmCategoryID  = "/get-film-category-id"
	deleteFilmCategory = "/delete-film-category"
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
	router.HandleFunc(addFilmCategory, appresult.Middleware(h.AddFilmCategory)).Methods(http.MethodPost)
	router.HandleFunc(getFilmCategory, appresult.Middleware(h.GetFilmCategory)).Methods(http.MethodPost)
	router.HandleFunc(getFilmCategoryID, appresult.Middleware(h.GetFilmCategoryID)).Methods(http.MethodPost)
	router.HandleFunc(deleteFilmCategory, appresult.Middleware(h.DeleteFilmCategory)).Methods(http.MethodPost)
}

// Add Film Category v1
// @Summary v1
// @Description post add film category
// @Tags dvd
// @ID add_film_category_v1
// @Produce json
// @Param FilmCategoryReq body FilmCategoryReq true "Film Category Request"
// @Success 200 {object} category.UUID
// @Failure 500	{object} string	"some err from db"
// @Router /film-category [post]
func (h *handler) AddFilmCategory(w http.ResponseWriter, r *http.Request) error {
	body, errBody := ioutil.ReadAll(r.Body)
	if errBody != nil {
		fmt.Println("Error body in handler Add film Category:", errBody)
		return appresult.ErrMissingParam
	}

	film := FilmCategoryReq{}
	errData := json.Unmarshal(body, &film)
	if errData != nil {
		fmt.Println("Error Data in handler Add film Category:", errData)
		return appresult.ErrMissingParam
	}

	data, err := h.repository.AddFilmCategory(context.TODO(), film)
	if err != nil {
		fmt.Println("Error in handler Add film Category:", err)
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

// Get Film Category v1
// @Summary v1
// @Description post get film category
// @Tags dvd
// @ID get_film_category_v1
// @Produce json
// @Param PaginationDTO body PaginationDTO true "Pagination DTO"
// @Success 200 {object} GetFilmCategoryResult
// @Failure 500	{object} string	"some err from db"
// @Router /get-film-category [post]
func (h *handler) GetFilmCategory(w http.ResponseWriter, r *http.Request) error {
	body, errBody := ioutil.ReadAll(r.Body)
	if errBody != nil {
		fmt.Println("Error body in handler GetFilmCategory:", errBody)
		return appresult.ErrMissingParam
	}

	dto := PaginationDTO{}
	errData := json.Unmarshal(body, &dto)
	if errData != nil {
		fmt.Println("Error Data in handler GetFilmCategory:", errData)
		return appresult.ErrMissingParam
	}

	data, count, err := h.repository.GetFilmCategory(context.TODO(), dto)
	if err != nil {
		fmt.Println("Error in handler GetFilmCategory:", err)
		return appresult.ErrInternalServer
	}

	successResult := appresult.Success
	successResult.Data = GetFilmCategoryResult{
		FilmCategory: data,
		Count:        count,
	}
	w.Header().Add(appresult.HeaderContentTypeJson())
	err = json.NewEncoder(w).Encode(successResult)
	if err != nil {
		return err
	}
	return nil
}

// Get Film Category ID v1
// @Summary v1
// @Description post get film category ID
// @Tags dvd
// @ID get_film_category_id_v1
// @Produce json
// @Param UUID body UUID true "UUID"
// @Success 200 {object} GetFilmCategory
// @Failure 500	{object} string	"some err from db"
// @Router /get-film-category-id [post]
func (h *handler) GetFilmCategoryID(w http.ResponseWriter, r *http.Request) error {
	body, errBody := ioutil.ReadAll(r.Body)
	if errBody != nil {
		fmt.Println("Error body in handler GetFilmCategoryID:", errBody)
		return appresult.ErrMissingParam
	}

	id := UUID{}
	errData := json.Unmarshal(body, &id)
	if errData != nil {
		fmt.Println("Error Data in handler GetFilmCategoryID:", errData)
		return appresult.ErrMissingParam
	}

	uuid, err := h.repository.GetFilmCategoryID(context.TODO(), id)
	if err != nil {
		fmt.Println("Error in handler GetFilmCategoryID:", err)
		return appresult.ErrInternalServer
	}
	successResult := appresult.Success
	successResult.Data = uuid
	w.Header().Add(appresult.HeaderContentTypeJson())
	err = json.NewEncoder(w).Encode(successResult)
	if err != nil {
		return err
	}
	return nil
}

// Delete Film Category v1
// @Summary v1
// @Description post delete film category
// @Tags dvd
// @ID delete_film_category_v1
// @Produce json
// @Param UUID body UUID true "UUID"
// @Success 200
// @Failure 500	{object} string	"some err from db"
// @Router /delete-film-category [post]
func (h *handler) DeleteFilmCategory(w http.ResponseWriter, r *http.Request) error {
	body, errBody := ioutil.ReadAll(r.Body)
	if errBody != nil {
		fmt.Println("Error body in handler DeleteFilmCategory:", errBody)
		return appresult.ErrMissingParam
	}
	id := UUID{}
	errData := json.Unmarshal(body, &id)
	if errData != nil {
		fmt.Println("Error Data in handler DeleteFilmCategory:", errData)
		return appresult.ErrMissingParam
	}
	err := h.repository.DeleteFilmCategory(context.TODO(), id)
	if err != nil {
		fmt.Println("Error in handler DeleteFilmCategory:", err)
		return appresult.ErrInternalServer
	}
	w.Header().Add(appresult.HeaderContentTypeJson())
	err = json.NewEncoder(w).Encode(appresult.Success)
	if err != nil {
		return err
	}
	return nil
}
