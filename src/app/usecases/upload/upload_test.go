package upload

import (
	"errors"
	mockInteg "palm_code_be/mocks/infra/integration/firestore"
	dto "palm_code_be/src/app/dto/upload"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockUploadUseCase struct {
	mock.Mock
}

type UploadUseCaseList struct {
	suite.Suite
	useCase   UploadUsecase
	dtoResp   *dto.UploadRespSTO
	mockInteg *mockInteg.MockFstoreRepo
}

func (suite *UploadUseCaseList) SetupTest() {

	suite.mockInteg = new(mockInteg.MockFstoreRepo)
	suite.useCase = NewUploadUseCase(suite.mockInteg)

	suite.dtoResp = &dto.UploadRespSTO{
		URL: "asdfghjkl.com",
	}
}

func (u *UploadUseCaseList) TestUploadSuccess() {
	u.mockInteg.Mock.On("Upload", "test", "test", "test").Return(u.dtoResp, nil)
	_, err := u.useCase.Upload("test", "test", "test")
	u.Equal(nil, err)
}

func (u *UploadUseCaseList) TestUploadFail() {
	u.mockInteg.Mock.On("Upload", "test", "test", "test").Return(nil, errors.New(mock.Anything))
	_, err := u.useCase.Upload("test", "test", "test")
	u.Equal(errors.New(mock.Anything), err)
}

func TestUsecase(t *testing.T) {
	suite.Run(t, new(UploadUseCaseList))
}
