package user

import (
	"fmt"
	dto "palm_code_be/src/app/dto/user"

	"log"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"

	"errors"
	helper "palm_code_be/src/infra/helper"
)

type UserRepository interface {
	Register(data *dto.RegisterReqDTO) (*dto.RegisterRespDTO, error)
	Login(data *dto.LoginReqDTO) (*dto.RegisterRespDTO, error)
}

const (
	Register = `INSERT INTO public.users (email, password) 
		values ($1, $2)`

	Login = `select u.id, u.email, u.password
	from public.users u
	where u.email = $1`
)

var statement PreparedStatement

type PreparedStatement struct {
	login *sqlx.Stmt
}

type userRepo struct {
	Connection *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	repo := &userRepo{
		Connection: db,
	}
	InitPreparedStatement(repo)
	return repo
}

func (p *userRepo) Preparex(query string) *sqlx.Stmt {
	statement, err := p.Connection.Preparex(query)
	if err != nil {
		log.Fatalf("Failed to preparex query: %s. Error: %s", query, err.Error())
	}

	return statement
}

func InitPreparedStatement(m *userRepo) {
	statement = PreparedStatement{
		login: m.Preparex(Login),
	}
}

func (p *userRepo) Register(data *dto.RegisterReqDTO) (resp *dto.RegisterRespDTO, err error) {
	var resultData dto.RegisterModel

	// Hash the password from the registration data
	pwd, err := hashPassword(data.Password)
	if err != nil {
		return nil, err
	}

	// Begin a new transaction
	tx, err := p.Connection.Beginx()
	if err != nil {
		log.Println("Failed Begin Tx Register  : ", err.Error())
		return nil, err
	}

	// Defer function to handle commit/rollback of the transaction
	defer func() {
		if p := recover(); p != nil {
			// Rollback transaction if a panic occurs
			tx.Rollback()
			log.Println("Recovered in Register:", p)
			err = fmt.Errorf("panic occurred: %v", p)
		} else if err != nil {
			// Rollback transaction if an error occurred
			tx.Rollback()
			log.Println("Rolling back transaction due to:", err)
		} else {
			// Commit the transaction if no error occurred
			err = tx.Commit()
			if err != nil {
				log.Println("Failed to commit transaction:", err.Error())
			}
		}
	}()

	// Execute the registration query and scan the result into resultData

	_, err = tx.Exec(Register, data.Email, pwd)
	if err != nil {
		log.Println("Failed Query Register: ", err.Error())
		return nil, err
	}

	// Initialize response object and generate token
	resp = &dto.RegisterRespDTO{}
	if resp.Token, err = helper.GenerateToken(&resultData); err != nil {
		return nil, err
	}

	// Return the response object if everything is successful
	return resp, nil
}

func (p *userRepo) Login(data *dto.LoginReqDTO) (*dto.RegisterRespDTO, error) {
	var resultData []*dto.RegisterModel
	var resp dto.RegisterRespDTO

	// Execute the login query
	if err := statement.login.Select(&resultData, data.Email); err != nil {
		return nil, err
	}

	// Check if no rows were returned from the query
	if len(resultData) < 1 {
		return nil, errors.New("no rows returned from the query")
	}

	// Verify the password
	if err := verifyPassword(resultData[0].Password, data.Password); err != nil {
		return nil, err
	}

	// Generate token
	token, err := helper.GenerateToken(resultData[0])
	if err != nil {
		return nil, err
	}
	resp.Token = token

	// Return the response object if everything is successful
	return &resp, nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func verifyPassword(hashedPassword, inputPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
}
