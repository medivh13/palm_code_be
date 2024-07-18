package user

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type UserDTOInterface interface {
	Validate() error
}

type RegisterReqDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (dto *RegisterReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.Email, validation.Required, is.Email),
		validation.Field(&dto.Password, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type RegisterRespDTO struct {
	Token string `json:"token"`
}

type LoginReqDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (dto *LoginReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.Email, validation.Required, is.Email),
		validation.Field(&dto.Password, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type RegisterModel struct {
	ID       int64  `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
}
