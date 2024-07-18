package media

import (
	dto "palm_code_be/src/app/dto/media"
	repo "palm_code_be/src/infra/persistence/postgres/media"
	"palm_code_be/src/interface/response"
	"sync"
)

type MediaUsecase interface {
	Create(data *dto.MediaCreateReqDTO) error
	Get(data *dto.MediaGetReqDTO) ([]*dto.MediaRespDTO, *response.Meta, error)
	GetByID(data *dto.MediaGetReqByIDDTO) (*dto.MediaRespDTO, error)
}

type mediaUseCase struct {
	Repo  repo.MediaRepository
	mutex sync.Mutex
}

func NewMediaUseCase(repo repo.MediaRepository) MediaUsecase {
	return &mediaUseCase{
		Repo: repo,
	}
}

func (uc *mediaUseCase) Create(data *dto.MediaCreateReqDTO) error {
	uc.mutex.Lock()         // Acquire the lock of the Mutex
	defer uc.mutex.Unlock() // Ensure the Mutex is unlocked when the function exits

	err := uc.Repo.Create(data)

	if err != nil {
		return err
	}
	return nil
}

func (uc *mediaUseCase) Get(data *dto.MediaGetReqDTO) ([]*dto.MediaRespDTO, *response.Meta, error) {
	resp, meta, err := uc.Repo.Get(data)

	if err != nil {
		return nil, nil, err
	}

	return resp, meta, nil
}

func (uc *mediaUseCase) GetByID(data *dto.MediaGetReqByIDDTO) (*dto.MediaRespDTO, error) {

	resp, err := uc.Repo.GetByID(data)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
