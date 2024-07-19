package media

import (
	dto "palm_code_be/src/app/dto/media"
	"palm_code_be/src/interface/response"

	"log"

	"github.com/jmoiron/sqlx"
)

type MediaRepository interface {
	Create(data *dto.MediaCreateReqDTO) error
	Get(data *dto.MediaGetReqDTO) ([]*dto.MediaRespDTO, *response.Meta, error)
	GetByID(data *dto.MediaGetReqByIDDTO) (*dto.MediaRespDTO, error)
}

const (
	Create = `INSERT INTO media (url, type, created_by)
		values ($1, $2, $3)`

	GetByID = `SELECT id, url, type FROM media WHERE id=$1 and created_by = $2`
	Get     = `SELECT count(id) over() as total_data, id, url, type FROM media 
	WHERE created_by=$1
	order by id asc LIMIT $2 OFFSET $3
	`
)

var statement PreparedStatement

type PreparedStatement struct {
	create  *sqlx.Stmt
	get     *sqlx.Stmt
	getByID *sqlx.Stmt
}

type mediaRepo struct {
	Connection *sqlx.DB
}

func NewMediaRepository(db *sqlx.DB) MediaRepository {
	repo := &mediaRepo{
		Connection: db,
	}
	InitPreparedStatement(repo)
	return repo
}

func (p *mediaRepo) Preparex(query string) *sqlx.Stmt {
	statement, err := p.Connection.Preparex(query)
	if err != nil {
		log.Fatalf("Failed to preparex query: %s. Error: %s", query, err.Error())
	}

	return statement
}

func InitPreparedStatement(m *mediaRepo) {
	statement = PreparedStatement{
		create:  m.Preparex(Create),
		get:     m.Preparex(Get),
		getByID: m.Preparex(GetByID),
	}
}

func (p *mediaRepo) Create(data *dto.MediaCreateReqDTO) error {
	_, err := statement.create.Exec(
		data.URL, data.Type, data.CreatedBy,
	)

	if err != nil {
		log.Println("Failed Query Create Media : ", err.Error())
		return err
	}

	return nil
}

func (p *mediaRepo) Get(data *dto.MediaGetReqDTO) ([]*dto.MediaRespDTO, *response.Meta, error) {
	var resultData []*dto.MediaRespModel
	offset := (data.Page - 1) * data.PerPage
	err := statement.get.Select(&resultData, data.CreatedByID, data.PerPage, offset)

	if err != nil {
		return nil, nil, err
	}

	if len(resultData) == 0 {
		return []*dto.MediaRespDTO{}, &response.Meta{Limit: int(data.PerPage), Skip: int(offset), Total: 0}, nil
	}

	metaData := &response.Meta{}
	metaData.Limit = int(data.PerPage)
	metaData.Skip = int(offset)
	metaData.Total = float64(resultData[0].TotalData)

	return dto.ToMedia(resultData), metaData, nil
}

func (p *mediaRepo) GetByID(data *dto.MediaGetReqByIDDTO) (*dto.MediaRespDTO, error) {
	var resultData []*dto.MediaRespModel

	err := statement.getByID.Select(&resultData, data.ID, data.CreatedBy)

	if err != nil {
		return nil, err
	}

	return dto.ToReturnMedia(resultData[0]), nil
}
