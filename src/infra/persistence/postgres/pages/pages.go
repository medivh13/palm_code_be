package pages

import (
	dto "palm_code_be/src/app/dto/pages"
	"palm_code_be/src/infra/helper"
	"palm_code_be/src/interface/response"

	"log"

	"github.com/jmoiron/sqlx"
)

type PageRepository interface {
	Create(data *dto.PageCreateReqDTO) error
	Update(data *dto.PageUpdateReqDTO) error
	Get(data *dto.PagesGetReqDTO) ([]*dto.PageRespDTO, *response.Meta, error)
	GetByID(data *dto.PagesGetReqByIDDTO) (*dto.PageRespDTO, error)
	Delete(data *dto.PageDeleteReqDTO) error
}

const (
	Create = `INSERT INTO pages (title, slug, banner_media, content, created_by)
		values ($1, $2, $3, $4, $5)`

	GetByID = `SELECT id, title, slug, banner_media, content, created_at, updated_at FROM pages WHERE id=$1 and created_by = $2`
	Get     = `SELECT count(id) over() as total_data, id, title, slug, banner_media, content, created_at, updated_at FROM pages 
	WHERE created_by=$1
	order by id asc LIMIT $2 OFFSET $3
	`
	QueryLockForUpdate = `
		SELECT 1
		FROM pages	
		WHERE id = $1
		FOR UPDATE
	`

	Update = `
		Update pages SET title = $1, slug = $2,
		banner_media = $3, content = $4,
		updated_at = CURRENT_TIMESTAMP, 
		updated_by = $5 WHERE id = $6
	`

	Delete = `
		Delete from pages where id = $1
	`
)

var statement PreparedStatement

type PreparedStatement struct {
	create  *sqlx.Stmt
	get     *sqlx.Stmt
	getByID *sqlx.Stmt
	delete  *sqlx.Stmt
}

type PageRepo struct {
	Connection *sqlx.DB
}

func NewPageRepository(db *sqlx.DB) PageRepository {
	repo := &PageRepo{
		Connection: db,
	}
	InitPreparedStatement(repo)
	return repo
}

func (p *PageRepo) Preparex(query string) *sqlx.Stmt {
	statement, err := p.Connection.Preparex(query)
	if err != nil {
		log.Fatalf("Failed to preparex query: %s. Error: %s", query, err.Error())
	}

	return statement
}

func InitPreparedStatement(m *PageRepo) {
	statement = PreparedStatement{
		create:  m.Preparex(Create),
		get:     m.Preparex(Get),
		getByID: m.Preparex(GetByID),
		delete:  m.Preparex(Delete),
	}
}

func (p *PageRepo) Create(data *dto.PageCreateReqDTO) error {

	_, err := statement.create.Exec(
		data.Title, data.Slug, data.BannerMedia, data.Content, data.CreatedBy,
	)

	if err != nil {
		log.Println("Failed Query Create Page : ", err.Error())
		return err
	}

	return nil
}

func (p *PageRepo) Update(data *dto.PageUpdateReqDTO) error {

	tx, err := p.Connection.Beginx()
	if err != nil {
		log.Println("Failed to begin transaction:", err.Error())
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
			if err != nil {
				log.Println("Failed to commit transaction:", err.Error())
			}
		}
	}()

	_, err = tx.Exec(QueryLockForUpdate, data.ID)
	if err != nil {
		log.Println("Failed to Lock query:", err.Error())
		return err
	}

	_, err = tx.Exec(Update, data.Title, data.Slug, data.BannerMedia, data.Content, data.UpdatedBy, data.ID)
	if err != nil {
		log.Println("Failed to Update Page query:", err.Error())
		return err
	}

	return nil
}

func (p *PageRepo) Get(data *dto.PagesGetReqDTO) ([]*dto.PageRespDTO, *response.Meta, error) {
	var resultData []*dto.PageRespModel
	offset := (data.Page - 1) * data.PerPage
	err := statement.get.Select(&resultData, data.CreatedByID, data.PerPage, offset)

	if err != nil {
		return nil, nil, err
	}

	if len(resultData) == 0 {
		return []*dto.PageRespDTO{}, &response.Meta{Limit: int(data.PerPage), Skip: int(offset), Total: 0}, nil
	}

	metaData := &response.Meta{}
	metaData.Limit = int(data.PerPage)
	metaData.Skip = int(offset)
	metaData.Total = float64(resultData[0].TotalData)

	return dto.ToPage(resultData), metaData, nil
}

func (p *PageRepo) GetByID(data *dto.PagesGetReqByIDDTO) (*dto.PageRespDTO, error) {
	var resultData []*dto.PageRespModel

	err := statement.getByID.Select(&resultData, data.ID, data.CreatedBy)

	if err != nil {
		return nil, err
	}

	return dto.ToReturnPage(resultData[0]), nil
}

func (p *PageRepo) Delete(data *dto.PageDeleteReqDTO) error {
	result, err := statement.delete.Exec(
		data.ID,
	)

	if err != nil {
		log.Println("Failed Query Delete Page : ", err.Error())
		return err
	}

	row, _ := result.RowsAffected()
	if row < 1 {
		log.Println("Failed Query Delete: ", helper.ErrNotFound)
		err = helper.ErrNotFound
		return err
	}

	return nil
}
