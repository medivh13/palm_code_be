package media

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	dto "palm_code_be/src/app/dto/media"
	usecases "palm_code_be/src/app/usecases/media"
	common_error "palm_code_be/src/infra/errors"
	"palm_code_be/src/infra/helper"
	"palm_code_be/src/interface/response"

	"github.com/go-chi/chi/v5"
	_ "github.com/joho/godotenv/autoload"
)

type MediaHandlerInterface interface {
	Get(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
}

type MediaHandler struct {
	response response.IResponseClient
	usecase  usecases.MediaUsecase
}

func NewMediaHandler(r response.IResponseClient, h usecases.MediaUsecase) MediaHandlerInterface {
	return &MediaHandler{
		response: r,
		usecase:  h,
	}
}

func (h *MediaHandler) Get(w http.ResponseWriter, r *http.Request) {

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

	getDTO := dto.MediaGetReqDTO{}
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
		"Successful Get Medias",
		data,
		meta,
	)
}

func (h *MediaHandler) GetByID(w http.ResponseWriter, r *http.Request) {

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

	getDTO := dto.MediaGetReqByIDDTO{}
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
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_RETRIEVE_DATA, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Get Medias",
		data,
		nil,
	)
}
