package team

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type TeamDTOInterface interface {
	Validate() error
}

type TeamCreateReqDTO struct {
	Name           string `json:"name"`
	Role           string `json:"role"`
	Bio            string `json:"bio"`
	ProfilePicture string `json:"profile_picture"`
	CreatedBy      int64  `json:"created_by"`
}

func (dto *TeamCreateReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.Name, validation.Required),
		validation.Field(&dto.Role, validation.Required),
		validation.Field(&dto.Bio, validation.Required),
		validation.Field(&dto.ProfilePicture, validation.Required),
		validation.Field(&dto.CreatedBy, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type TeamUpdateReqDTO struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	Role           string `json:"role"`
	Bio            string `json:"bio"`
	ProfilePicture string `json:"profile_picture"`
	UpdatedBy      int64    `json:"updated_by"`
}

func (dto *TeamUpdateReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.ID, validation.Required),
		validation.Field(&dto.Name, validation.Required),
		validation.Field(&dto.Role, validation.Required),
		validation.Field(&dto.Bio, validation.Required),
		validation.Field(&dto.ProfilePicture, validation.Required),
		validation.Field(&dto.UpdatedBy, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type TeamDeleteReqDTO struct {
	ID int64 `json:"id"`
}

func (dto *TeamDeleteReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.ID, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type TeamGetReqDTO struct {
	CreatedByID int64 `json:"id"`
	Page        int64 `json:"page"`
	PerPage     int64 `json:"perPage"`
}

type TeamGetReqByIDDTO struct {
	ID        int64 `json:"id"`
	CreatedBy int64 `json:"created_by"`
}

func (dto *TeamGetReqByIDDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.ID, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type TeamRespDTO struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name"`
	Role           string    `json:"role"`
	Bio            string    `json:"bio"`
	ProfilePicture string    `json:"profile_picture"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type TeamRespModel struct {
	ID             int64     `db:"id"`
	Name           string    `db:"name"`
	Role           string    `db:"role"`
	Bio            string    `db:"bio"`
	ProfilePicture string    `db:"profile_picture"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
	TotalData      int64     `db:"total_data"`
}
