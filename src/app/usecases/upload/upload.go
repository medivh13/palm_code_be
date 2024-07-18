package upload

import (
	dto "palm_code_be/src/app/dto/upload"
	fStore "palm_code_be/src/infra/integration/firestorage"
	"sync"
)

type UploadUsecase interface {
	Upload(bucketName, objectName, filePath string) (*dto.UploadRespSTO, error)
}

type uploadUseCase struct {
	mutex     sync.Mutex
	FireStore fStore.FireStorage
}

func NewUploadUseCase(fStore fStore.FireStorage) UploadUsecase {
	return &uploadUseCase{
		FireStore: fStore,
	}
}

func (uc *uploadUseCase) Upload(bucketName, objectName, filePath string) (*dto.UploadRespSTO, error) {
	uc.mutex.Lock()         // Acquire the lock of the Mutex
	defer uc.mutex.Unlock() // Ensure the Mutex is unlocked when the function exits
	resp := dto.UploadRespSTO{}
	url, err := uc.FireStore.Upload(bucketName, objectName, filePath)

	if err != nil {
		return nil, err
	}

	resp.URL = url
	return &resp, nil
}
