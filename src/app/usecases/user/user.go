package user

import (
	"log"

	dto "palm_code_be/src/app/dto/user"
	"sync"
	repo "palm_code_be/src/infra/persistence/postgres/user"
)

type UserUCInterface interface {
	Register(data *dto.RegisterReqDTO) (*dto.RegisterRespDTO, error)
	Login(data *dto.LoginReqDTO) (*dto.RegisterRespDTO, error)
}

type userUseCase struct {
	Repo repo.UserRepository
	mutex sync.Mutex
}

func NewUserUseCase(repo repo.UserRepository) UserUCInterface {
	return &userUseCase{
		Repo: repo,
	}
}

func (uc *userUseCase) Register(data *dto.RegisterReqDTO) (*dto.RegisterRespDTO, error) {
	uc.mutex.Lock()
	defer uc.mutex.Unlock()
	
	result, err := uc.Repo.Register(data)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (uc *userUseCase) Login(data *dto.LoginReqDTO) (*dto.RegisterRespDTO, error) {

	result, err := uc.Repo.Login(data)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}
