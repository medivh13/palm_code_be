package team

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	dto "palm_code_be/src/app/dto/team"
	usecases "palm_code_be/src/app/usecases/team"
	common_error "palm_code_be/src/infra/errors"
	"palm_code_be/src/infra/helper"
	"palm_code_be/src/interface/response"

	"github.com/go-chi/chi/v5"
	_ "github.com/joho/godotenv/autoload"
)

type TeamHandlerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type teamHandler struct {
	response response.IResponseClient
	usecase  usecases.TeamUsecase
}

func NewTeamHandler(r response.IResponseClient, h usecases.TeamUsecase) TeamHandlerInterface {
	return &teamHandler{
		response: r,
		usecase:  h,
	}
}

func (h *teamHandler) Create(w http.ResponseWriter, r *http.Request) {
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

	createDTO := dto.TeamCreateReqDTO{}
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
		"Successful Create Team",
		nil,
		nil,
	)
}

func (h *teamHandler) Get(w http.ResponseWriter, r *http.Request) {

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

	getDTO := dto.TeamGetReqDTO{}
	getDTO.CreatedByID = dataClaim.UserID

	if r.URL.Query().Get("Team") != "" {
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		getDTO.Page = int64(page)
	} else {
		getDTO.Page = helper.Page
	}

	if r.URL.Query().Get("perTeam") != "" {
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
		"Successful Get Teams",
		data,
		meta,
	)
}

func (h *teamHandler) GetByID(w http.ResponseWriter, r *http.Request) {

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

	getDTO := dto.TeamGetReqByIDDTO{}
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
		"Successful Get Teams",
		data,
		nil,
	)
}

func (h *teamHandler) Update(w http.ResponseWriter, r *http.Request) {

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

	updateDTO := dto.TeamUpdateReqDTO{}
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
		"Successful Update Team",
		nil,
		nil,
	)
}

func (h *teamHandler) Delete(w http.ResponseWriter, r *http.Request) {

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

	delDTO := dto.TeamDeleteReqDTO{}
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
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_CREATE_DATA, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Delete Team",
		nil,
		nil,
	)
}
