package team

import (
	dto "palm_code_be/src/app/dto/team"
	"palm_code_be/src/infra/helper"
	"palm_code_be/src/interface/response"

	"log"

	"github.com/jmoiron/sqlx"
)

type TeamRepository interface {
	Create(data *dto.TeamCreateReqDTO) error
	Update(data *dto.TeamUpdateReqDTO) error
	Get(data *dto.TeamGetReqDTO) ([]*dto.TeamRespDTO, *response.Meta, error)
	GetByID(data *dto.TeamGetReqByIDDTO) (*dto.TeamRespDTO, error)
	Delete(data *dto.TeamDeleteReqDTO) error
}

const (
	Create = `INSERT INTO team (name, role, bio, profile_picture, created_by)
		values ($1, $2, $3, $4, $5)`

	GetByID = `SELECT id, name, role, bio, profile_picture FROM team WHERE id=$1 and created_by = $2`
	Get     = `SELECT count(id) over() as total_data, id, name, role, bio, profile_picture FROM team 
	WHERE created_by=$1
	order by id asc LIMIT $2 OFFSET $3
	`
	QueryLockForUpdate = `
		SELECT 1
		FROM team	
		WHERE id = $1
		FOR UPDATE
	`

	Update = `
		Update team SET name = $1, role = $2, bio = $3, 
		profile_picture = $4, updated_at = CURRENT_TIMESTAMP, 
		updated_by = $5 WHERE id = $6
	`

	Delete = `
		Delete from team where id = $1
	`
)

var statement PreparedStatement

type PreparedStatement struct {
	create  *sqlx.Stmt
	get     *sqlx.Stmt
	getByID *sqlx.Stmt
	delete  *sqlx.Stmt
}

type teamRepo struct {
	Connection *sqlx.DB
}

func NewTeamRepository(db *sqlx.DB) TeamRepository {
	repo := &teamRepo{
		Connection: db,
	}
	InitPreparedStatement(repo)
	return repo
}

func (p *teamRepo) Preparex(query string) *sqlx.Stmt {
	statement, err := p.Connection.Preparex(query)
	if err != nil {
		log.Fatalf("Failed to preparex query: %s. Error: %s", query, err.Error())
	}

	return statement
}

func InitPreparedStatement(m *teamRepo) {
	statement = PreparedStatement{
		create:  m.Preparex(Create),
		get:     m.Preparex(Get),
		getByID: m.Preparex(GetByID),
		delete:  m.Preparex(Delete),
	}
}

func (p *teamRepo) Create(data *dto.TeamCreateReqDTO) error {
	_, err := statement.create.Exec(
		data.Name, data.Role, data.Bio, data.ProfilePicture, data.CreatedBy,
	)

	if err != nil {
		log.Println("Failed Query Create Team : ", err.Error())
		return err
	}

	return nil
}

func (p *teamRepo) Update(data *dto.TeamUpdateReqDTO) error {

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

	_, err = tx.Exec(Update, data.Name, data.Role, data.Bio, data.ProfilePicture, data.UpdatedBy, data.ID)
	if err != nil {
		log.Println("Failed to Update Team query:", err.Error())
		return err
	}

	return nil
}

func (p *teamRepo) Get(data *dto.TeamGetReqDTO) ([]*dto.TeamRespDTO, *response.Meta, error) {
	var resultData []*dto.TeamRespModel
	offset := (data.Page - 1) * data.PerPage
	err := statement.get.Select(&resultData, data.CreatedByID, data.PerPage, offset)

	if err != nil {
		return nil, nil, err
	}

	if len(resultData) == 0 {
		return []*dto.TeamRespDTO{}, &response.Meta{Limit: int(data.PerPage), Skip: int(offset), Total: 0}, nil
	}

	metaData := &response.Meta{}
	metaData.Limit = int(data.PerPage)
	metaData.Skip = int(offset)
	metaData.Total = float64(resultData[0].TotalData)

	return dto.ToTeam(resultData), metaData, nil
}

func (p *teamRepo) GetByID(data *dto.TeamGetReqByIDDTO) (*dto.TeamRespDTO, error) {
	var resultData []*dto.TeamRespModel

	err := statement.getByID.Select(&resultData, data.ID, data.CreatedBy)

	if err != nil {
		return nil, err
	}

	return dto.ToReturnTeam(resultData[0]), nil
}

func (p *teamRepo) Delete(data *dto.TeamDeleteReqDTO) error {
	result, err := statement.delete.Exec(
		data.ID,
	)

	if err != nil {
		log.Println("Failed Query Delete Team : ", err.Error())
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
