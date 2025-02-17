package user

import (
	"encoding/json"
	"log"
	"net/http"

	dto "palm_code_be/src/app/dto/user"
	usecases "palm_code_be/src/app/usecases/user"
	common_error "palm_code_be/src/infra/errors"
	"palm_code_be/src/interface/response"

	"github.com/lib/pq"
)

type UserHandlerInterface interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	response response.IResponseClient
	usecase  usecases.UserUCInterface
}

func NewUserandler(r response.IResponseClient, h usecases.UserUCInterface) UserHandlerInterface {
	return &userHandler{
		response: r,
		usecase:  h,
	}
}

func (h *userHandler) Register(w http.ResponseWriter, r *http.Request) {

	postDTO := dto.RegisterReqDTO{}
	err := json.NewDecoder(r.Body).Decode(&postDTO)
	if err != nil {
		log.Println(err)

		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	err = postDTO.Validate()
	if err != nil {
		log.Println(err)
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	data, err := h.usecase.Register(&postDTO)
	if err != nil {
		log.Println(err)
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			h.response.HttpError(w, common_error.NewError(common_error.USER_ALREADY_EXIST, err))
			return
		}
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_CREATE_DATA, err))
		return
	}
	w.WriteHeader(http.StatusCreated)
	h.response.JSON(
		w,
		"Successful Register New User",
		data,
		nil,
	)
}

func (h *userHandler) Login(w http.ResponseWriter, r *http.Request) {

	postDTO := dto.LoginReqDTO{}
	err := json.NewDecoder(r.Body).Decode(&postDTO)
	if err != nil {
		log.Println(err)
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}
	err = postDTO.Validate()
	if err != nil {
		log.Println(err)
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	data, err := h.usecase.Login(&postDTO)
	if err != nil {
		log.Println(err)

		h.response.HttpError(w, common_error.NewError(common_error.UNAUTHORIZED, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Login",
		data,
		nil,
	)
}
