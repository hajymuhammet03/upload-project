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
	addCategory    = "/add-category"
	getCategory    = "/get-category"
	getCategoryID  = "/get-category-id"
	deleteCategory = "/delete-category"
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
	router.HandleFunc(getCategory, appresult.Middleware(h.GetCategory)).Methods(http.MethodPost)
	router.HandleFunc(getCategoryID, appresult.Middleware(h.GetCategoryID)).Methods(http.MethodPost)
	router.HandleFunc(deleteCategory, appresult.Middleware(h.DeleteCategory)).Methods(http.MethodPost)
}

// Add Category v1
// @Summary v1
// @Description post add category
// @Tags dvd
// @ID add_category_v1
// @Produce json
// @Param AddCategory body AddCategory true "Add Category"
// @Success 200 {object} category.UUID
// @Failure 500	{object} string	"some err from db"
// @Router /add-category [post]
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

// Get Category v1
// @Summary v1
// @Description post get category
// @Tags dvd
// @ID get_category_v1
// @Produce json
// @Param PaginationDTO body PaginationDTO true "Pagination DTO"
// @Success 200 {object} category.GetCategoryResult
// @Failure 500	{object} string	"some err from db"
// @Router /get-category [post]
func (h *handler) GetCategory(w http.ResponseWriter, r *http.Request) error {
	body, errBody := ioutil.ReadAll(r.Body)
	if errBody != nil {
		fmt.Println("Get Category in handler err body:", errBody)
		return appresult.ErrMissingParam
	}

	category := PaginationDTO{}
	errData := json.Unmarshal(body, &category)
	if errData != nil {
		fmt.Println("Get Category in handler err data: ", errData)
		return appresult.ErrMissingParam
	}

	data, count, err := h.repository.GetCategory(context.TODO(), category)
	if err != nil {
		fmt.Println("Get Category in handler err: ", err)
		return appresult.ErrInternalServer
	}

	successResult := appresult.Success
	successResult.Data = GetCategoryResult{
		Category: data,
		Count:    count,
	}

	w.Header().Add(appresult.HeaderContentTypeJson())
	err = json.NewEncoder(w).Encode(successResult)
	if err != nil {
		return err
	}
	return nil
}

// Get Category ID v1
// @Summary v1
// @Description post get category id
// @Tags dvd
// @ID get_category_id_v1
// @Produce json
// @Param category.UUID body category.UUID true "UUID"
// @Success 200 {object} category.GetCategory
// @Failure 500	{object} string	"some err from db"
// @Router /get-category-id [post]
func (h *handler) GetCategoryID(w http.ResponseWriter, r *http.Request) error {
	body, errBody := ioutil.ReadAll(r.Body)
	if errBody != nil {
		fmt.Println("Get Category ID in handler err body:", errBody)
		return appresult.ErrMissingParam
	}

	id := UUID{}
	errData := json.Unmarshal(body, &id)
	if errData != nil {
		fmt.Println("Get Category ID in handler err data: ", errData)
		return appresult.ErrMissingParam
	}

	data, err := h.repository.GetCategoryID(context.TODO(), id)
	if err != nil {
		fmt.Println("Get Category ID in handler err: ", err)
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

// DELETE Category v1
// @Summary v1
// @Description post delete category
// @Tags dvd
// @ID delete_category_v1
// @Produce json
// @Param category.UUID body category.UUID true "UUID"
// @Success 200
// @Failure 500	{object} string	"some err from db"
// @Router /delete-category [post]
func (h *handler) DeleteCategory(w http.ResponseWriter, r *http.Request) error {
	body, errBody := ioutil.ReadAll(r.Body)
	if errBody != nil {
		fmt.Println("Delete Category in handler err body:", errBody)
		return appresult.ErrMissingParam
	}
	id := UUID{}
	errData := json.Unmarshal(body, &id)
	if errData != nil {
		fmt.Println("Delete Category in handler err data: ", errData)
		return appresult.ErrMissingParam
	}
	err := h.repository.DeleteCategory(context.TODO(), id)
	if err != nil {
		fmt.Println("Delete Category in handler err: ", err)
		return appresult.ErrInternalServer
	}
	w.Header().Add(appresult.HeaderContentTypeJson())
	successResult := appresult.Success
	successResult.Data = ""
	err = json.NewEncoder(w).Encode(successResult)
	if err != nil {
		return err
	}
	return nil
}
