package page

import (
	dto "palm_code_be/src/app/dto/pages"
	repo "palm_code_be/src/infra/persistence/postgres/pages"
	"palm_code_be/src/interface/response"
	"sync"
)

type PageUsecase interface {
	Create(data *dto.PageCreateReqDTO) error
	Update(data *dto.PageUpdateReqDTO) error
	Get(data *dto.PagesGetReqDTO) ([]*dto.PageRespDTO, *response.Meta, error)
	GetByID(data *dto.PagesGetReqByIDDTO) (*dto.PageRespDTO, error)
	Delete(data *dto.PageDeleteReqDTO) error
}

type pageUseCase struct {
	Repo  repo.PageRepository
	mutex sync.Mutex
}

func NewPageUseCase(repo repo.PageRepository) PageUsecase {
	return &pageUseCase{
		Repo: repo,
	}
}

func (uc *pageUseCase) Create(data *dto.PageCreateReqDTO) error {
	uc.mutex.Lock()         // Acquire the lock of the Mutex
	defer uc.mutex.Unlock() // Ensure the Mutex is unlocked when the function exits

	err := uc.Repo.Create(data)

	if err != nil {
		return err
	}
	return nil
}

func (uc *pageUseCase) Get(data *dto.PagesGetReqDTO) ([]*dto.PageRespDTO, *response.Meta, error) {
	resp, meta, err := uc.Repo.Get(data)

	if err != nil {
		return nil, nil, err
	}

	return resp, meta, nil
}

func (uc *pageUseCase) GetByID(data *dto.PagesGetReqByIDDTO) (*dto.PageRespDTO, error) {

	resp, err := uc.Repo.GetByID(data)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (uc *pageUseCase) Update(data *dto.PageUpdateReqDTO) error {
	uc.mutex.Lock()         // Acquire the lock of the Mutex
	defer uc.mutex.Unlock() // Ensure the Mutex is unlocked when the function exits

	err := uc.Repo.Update(data)

	if err != nil {
		return nil
	}

	return nil
}

func (uc *pageUseCase) Delete(data *dto.PageDeleteReqDTO) error {
	err := uc.Repo.Delete(data)

	if err != nil {
		return err
	}

	return nil
}
