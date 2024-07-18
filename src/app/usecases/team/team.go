package team

import (
	dto "palm_code_be/src/app/dto/team"
	repo "palm_code_be/src/infra/persistence/postgres/team"
	"palm_code_be/src/interface/response"
	"sync"
)

type TeamUsecase interface {
	Create(data *dto.TeamCreateReqDTO) error
	Update(data *dto.TeamUpdateReqDTO) error
	Get(data *dto.TeamGetReqDTO) ([]*dto.TeamRespDTO, *response.Meta, error)
	GetByID(data *dto.TeamGetReqByIDDTO) (*dto.TeamRespDTO, error)
	Delete(data *dto.TeamDeleteReqDTO) error
}

type teamUseCase struct {
	Repo  repo.TeamRepository
	mutex sync.Mutex
}

func NewTeamUseCase(repo repo.TeamRepository) TeamUsecase {
	return &teamUseCase{
		Repo: repo,
	}
}

func (uc *teamUseCase) Create(data *dto.TeamCreateReqDTO) error {
	uc.mutex.Lock()         // Acquire the lock of the Mutex
	defer uc.mutex.Unlock() // Ensure the Mutex is unlocked when the function exits

	err := uc.Repo.Create(data)

	if err != nil {
		return err
	}
	return nil
}

func (uc *teamUseCase) Update(data *dto.TeamUpdateReqDTO) error {
	uc.mutex.Lock()         // Acquire the lock of the Mutex
	defer uc.mutex.Unlock() // Ensure the Mutex is unlocked when the function exits

	err := uc.Repo.Update(data)

	if err != nil {
		return err
	}
	return nil
}

func (uc *teamUseCase) Get(data *dto.TeamGetReqDTO) ([]*dto.TeamRespDTO, *response.Meta, error) {
	resp, meta, err := uc.Repo.Get(data)

	if err != nil {
		return nil, nil, err
	}

	return resp, meta, nil
}

func (uc *teamUseCase) GetByID(data *dto.TeamGetReqByIDDTO) (*dto.TeamRespDTO, error) {

	resp, err := uc.Repo.GetByID(data)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (uc *teamUseCase) Delete(data *dto.TeamDeleteReqDTO) error {

	err := uc.Repo.Delete(data)

	if err != nil {
		return err
	}

	return nil
}
