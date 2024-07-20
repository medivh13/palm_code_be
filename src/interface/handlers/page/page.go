package page

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	dto "palm_code_be/src/app/dto/pages"
	usecases "palm_code_be/src/app/usecases/page"
	common_error "palm_code_be/src/infra/errors"
	"palm_code_be/src/infra/helper"
	"palm_code_be/src/interface/response"

	"github.com/go-chi/chi/v5"
	_ "github.com/joho/godotenv/autoload"
)

type PageHandlerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type pageHandler struct {
	response response.IResponseClient
	usecase  usecases.PageUsecase
}

func NewPageHandler(r response.IResponseClient, h usecases.PageUsecase) PageHandlerInterface {
	return &pageHandler{
		response: r,
		usecase:  h,
	}
}

func (h *pageHandler) Create(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		h.response.HttpError(w, common_error.NewError(common_error.UNAUTHORIZED, errors.New("missing authorization token")))
		return
	}

	// Verifikasi token
	dataClaim, err := helper.VerifyToken(tokenString)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.UNAUTHORIZED, err))
		return
	}

	createDTO := dto.PageCreateReqDTO{}
	err = json.NewDecoder(r.Body).Decode(&createDTO)
	if err != nil {
		log.Println(err)
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	createDTO.CreatedBy = dataClaim.UserID

	err = createDTO.Validate()
	if err != nil {
		log.Println(err)
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	err = h.usecase.Create(&createDTO)
	if err != nil {
		log.Println(err)
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_CREATE_DATA, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Create Page",
		nil,
		nil,
	)
}

func (h *pageHandler) Get(w http.ResponseWriter, r *http.Request) {

	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		h.response.HttpError(w, common_error.NewError(common_error.UNAUTHORIZED, errors.New("missing authorization token")))
		return
	}

	// Verifikasi token
	dataClaim, err := helper.VerifyToken(tokenString)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.UNAUTHORIZED, err))
		return
	}

	getDTO := dto.PagesGetReqDTO{}
	getDTO.CreatedByID = dataClaim.UserID

	if r.URL.Query().Get("page") != "" {
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		getDTO.Page = int64(page)
	} else {
		getDTO.Page = helper.Page
	}

	if r.URL.Query().Get("perPage") != "" {
		perPage, _ := strconv.Atoi(r.URL.Query().Get("perPage"))
		getDTO.PerPage = int64(perPage)
	} else {
		getDTO.PerPage = helper.PerPage
	}

	data, meta, err := h.usecase.Get(&getDTO)
	if err != nil {
		log.Println(err)
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_RETRIEVE_DATA, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Get Pages",
		data,
		meta,
	)
}

func (h *pageHandler) GetByID(w http.ResponseWriter, r *http.Request) {

	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		h.response.HttpError(w, common_error.NewError(common_error.UNAUTHORIZED, errors.New("missing authorization token")))
		return
	}

	// Verifikasi token
	dataClaim, err := helper.VerifyToken(tokenString)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.UNAUTHORIZED, err))
		return
	}

	getDTO := dto.PagesGetReqByIDDTO{}
	getDTO.CreatedBy = dataClaim.UserID

	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	getDTO.ID = int64(id)

	err = getDTO.Validate()

	if err != nil {
		log.Println(err)
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	data, err := h.usecase.GetByID(&getDTO)
	if err != nil {
		log.Println(err)
		if err.Error() == "data not found" {
			h.response.HttpError(w, common_error.NewError(common_error.STATUS_PAGE_NOT_FOUND, err))
			return
		}
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_RETRIEVE_DATA, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Get Pages",
		data,
		nil,
	)
}

func (h *pageHandler) Update(w http.ResponseWriter, r *http.Request) {

	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		h.response.HttpError(w, common_error.NewError(common_error.UNAUTHORIZED, errors.New("missing authorization token")))
		return
	}

	// Verifikasi token
	dataClaim, err := helper.VerifyToken(tokenString)
	if err != nil {
		log.Println(err)
		h.response.HttpError(w, common_error.NewError(common_error.UNAUTHORIZED, err))
		return
	}

	updateDTO := dto.PageUpdateReqDTO{}
	err = json.NewDecoder(r.Body).Decode(&updateDTO)
	if err != nil {
		log.Println(err)
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}
	updateDTO.UpdatedBy = dataClaim.UserID
	err = updateDTO.Validate()
	if err != nil {
		log.Println(err)
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	err = h.usecase.Update(&updateDTO)
	if err != nil {
		log.Println(err)
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_CREATE_DATA, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Update Page",
		nil,
		nil,
	)
}

func (h *pageHandler) Delete(w http.ResponseWriter, r *http.Request) {

	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		h.response.HttpError(w, common_error.NewError(common_error.UNAUTHORIZED, errors.New("missing authorization token")))
		return
	}

	// Verifikasi token
	_, err := helper.VerifyToken(tokenString)
	if err != nil {
		log.Println(err)
		h.response.HttpError(w, common_error.NewError(common_error.UNAUTHORIZED, err))
		return
	}

	delDTO := dto.PageDeleteReqDTO{}
	err = json.NewDecoder(r.Body).Decode(&delDTO)
	if err != nil {
		log.Println(err)
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	err = delDTO.Validate()
	if err != nil {
		log.Println(err)
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	err = h.usecase.Delete(&delDTO)
	if err != nil {
		log.Println(err)
		if errors.Is(err, helper.ErrNotFound) {
			h.response.HttpError(w, common_error.NewError(common_error.STATUS_PAGE_NOT_FOUND, err))
			return
		}
		h.response.HttpError(w, common_error.NewError(common_error.UNKNOWN_ERROR, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Delete Page",
		nil,
		nil,
	)
}
