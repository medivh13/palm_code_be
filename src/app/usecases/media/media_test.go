package media

import (
	"errors"
	mockDTO "palm_code_be/mocks/app/dto/media"
	"testing"

	mockRepo "palm_code_be/mocks/infra/persistence/postgres/media"
	dto "palm_code_be/src/app/dto/media"
	"palm_code_be/src/interface/response"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockMediaUseCase struct {
	mock.Mock
}

type MediaUseCaseList struct {
	suite.Suite
	mockDTO    *mockDTO.MockMediasDTO
	mockRepo   *mockRepo.MockMediaRepo
	useCase    MediaUsecase
	dtoCreate  *dto.MediaCreateReqDTO
	dtoGet     *dto.MediaGetReqDTO
	dtoGetByID *dto.MediaGetReqByIDDTO
	dtoResp    *dto.MediaRespDTO
	dtoGetResp []*dto.MediaRespDTO
	meta       *response.Meta
}

func (suite *MediaUseCaseList) SetupTest() {
	suite.mockDTO = new(mockDTO.MockMediasDTO)
	suite.mockRepo = new(mockRepo.MockMediaRepo)
	suite.useCase = NewMediaUseCase(suite.mockRepo)

	suite.dtoCreate = &dto.MediaCreateReqDTO{
		URL:       "image.com",
		Type:      "image",
		CreatedBy: 1,
	}

	suite.dtoGet = &dto.MediaGetReqDTO{
		Page:        1,
		PerPage:     10,
		CreatedByID: 1,
	}

	suite.dtoGetByID = &dto.MediaGetReqByIDDTO{
		ID:        1,
		CreatedBy: 1,
	}

	suite.dtoResp = &dto.MediaRespDTO{
		ID:   1,
		URL:  "image.com",
		Type: "image",
	}

	suite.dtoGetResp = []*dto.MediaRespDTO{
		{
			ID:   1,
			URL:  "image.com",
			Type: "image",
		},
	}

	suite.meta = &response.Meta{
		Skip:  0,
		Limit: 10,
		Total: 1,
	}

}

func (u *MediaUseCaseList) TestCreateSuccess() {
	u.mockRepo.Mock.On("Create", u.dtoCreate).Return(nil)
	err := u.useCase.Create(u.dtoCreate)
	u.Equal(nil, err)
}

func (u *MediaUseCaseList) TestCreateFailed() {
	u.mockRepo.Mock.On("Create", u.dtoCreate).Return(errors.New(mock.Anything))
	err := u.useCase.Create(u.dtoCreate)
	u.Equal(errors.New(mock.Anything), err)
}

func (u *MediaUseCaseList) TestGetSuccess() {
	u.mockRepo.Mock.On("Get", u.dtoGet).Return(u.dtoGetResp, u.meta, nil)
	_, _, err := u.useCase.Get(u.dtoGet)
	u.Equal(nil, err)
}

func (u *MediaUseCaseList) TestGetFailed() {
	u.mockRepo.Mock.On("Get", u.dtoGet).Return(nil, nil, errors.New(mock.Anything))
	_, _, err := u.useCase.Get(u.dtoGet)
	u.Equal(errors.New(mock.Anything), err)
}

func (u *MediaUseCaseList) TestGetByIDSuccess() {
	u.mockRepo.Mock.On("GetByID", u.dtoGetByID).Return(u.dtoGetResp, nil)
	_, err := u.useCase.GetByID(u.dtoGetByID)
	u.Equal(nil, err)
}

func (u *MediaUseCaseList) TestGetByIDFailed() {
	u.mockRepo.Mock.On("GetByID", u.dtoGetByID).Return(nil, errors.New(mock.Anything))
	_, err := u.useCase.GetByID(u.dtoGetByID)
	u.Equal(errors.New(mock.Anything), err)
}

func TestUsecase(t *testing.T) {
	suite.Run(t, new(MediaUseCaseList))
}
