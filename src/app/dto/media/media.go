package media

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type MediaDTOInterface interface {
	Validate() error
}

type MediaCreateReqDTO struct {
	URL       string `json:"url"`
	Type      string `json:"type"`
	CreatedBy int64  `json:"created_by"`
}

func (dto *MediaCreateReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.URL, validation.Required),
		validation.Field(&dto.Type, validation.Required),
		validation.Field(&dto.CreatedBy, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type MediaUpdateReqDTO struct {
	ID        int64  `json:"id"`
	URL       string `json:"url"`
	Type      string `json:"type"`
	UpdatedBy int64  `json:"updated_by"`
}

func (dto *MediaUpdateReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.ID, validation.Required),
		validation.Field(&dto.URL, validation.Required),
		validation.Field(&dto.Type, validation.Required),
		validation.Field(&dto.UpdatedBy, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type MediaDeleteReqDTO struct {
	ID int64 `json:"id"`
}

func (dto *MediaDeleteReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.ID, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type MediaGetReqDTO struct {
	CreatedByID int64 `json:"id"`
	Page        int64 `json:"page"`
	PerPage     int64 `json:"perPage"`
}

type MediaGetReqByIDDTO struct {
	ID        int64 `json:"id"`
	CreatedBy int64 `json:"created_by"`
}

func (dto *MediaGetReqByIDDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.ID, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type MediaRespDTO struct {
	ID        int64     `json:"id" db:"id"`
	URL       string    `json:"url" db:"url"`
	Type      string    `json:"type" db:"type"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type MediaRespModel struct {
	ID        int64     `db:"id"`
	URL       string    `json:"url" db:"url"`
	Type      string    `json:"type" db:"type"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	TotalData int64     `db:"total_data"`
}
