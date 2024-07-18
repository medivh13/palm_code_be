package pages

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type PagesDTOInterface interface {
	Validate() error
}

type PageCreateReqDTO struct {
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	BannerMedia string `json:"banner_media"`
	Content     string `json:"content"`
	CreatedBy   int64  `json:"created_by"`
}

func (dto *PageCreateReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.Title, validation.Required),
		validation.Field(&dto.Slug, validation.Required),
		validation.Field(&dto.BannerMedia, validation.Required),
		validation.Field(&dto.Content, validation.Required),
		validation.Field(&dto.CreatedBy, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type PageUpdateReqDTO struct {
	ID              int64     `json:"id"`
	Title           string    `json:"title"`
	Slug            string    `json:"slug"`
	BannerMedia     string    `json:"banner_media"`
	Content         string    `json:"content"`
	PublicationDate time.Time `json:"publication_date"`
	UpdatedBy       int64     `json:"updated_by"`
}

func (dto *PageUpdateReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.Title, validation.Required),
		validation.Field(&dto.Slug, validation.Required),
		validation.Field(&dto.BannerMedia, validation.Required),
		validation.Field(&dto.Content, validation.Required),
		validation.Field(&dto.PublicationDate, validation.Required),
		validation.Field(&dto.UpdatedBy, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type PagesGetReqDTO struct {
	CreatedByID int64 `json:"id"`
	Page        int64 `json:"page"`
	PerPage     int64 `json:"perPage"`
}

type PagesGetReqByIDDTO struct {
	ID        int64 `json:"id"`
	CreatedBy int64 `json:"created_by"`
}

func (dto *PagesGetReqByIDDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.ID, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type PageDeleteReqDTO struct {
	ID int64 `json:"id"`
}

func (dto *PageDeleteReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.ID, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type PageRespDTO struct {
	ID              int64     `json:"id" db:"id"`
	Title           string    `json:"title" db:"title"`
	Slug            string    `json:"slug" db:"slug"`
	BannerMedia     string    `json:"banner_media" db:"banner_media"`
	Content         string    `json:"content" db:"content"`
	PublicationDate time.Time `json:"publication_date" db:"publication_date"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

type PageRespModel struct {
	ID              int64     `json:"id" db:"id"`
	Title           string    `json:"title" db:"title"`
	Slug            string    `json:"slug" db:"slug"`
	BannerMedia     string    `json:"banner_media" db:"banner_media"`
	Content         string    `json:"content" db:"content"`
	PublicationDate time.Time `json:"publication_date" db:"publication_date"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
	TotalData       int64     `db:"total_data"`
}
