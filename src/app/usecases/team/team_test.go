package team

import (
	"errors"
	mockDTO "palm_code_be/mocks/app/dto/team"
	"testing"

	mockRepo "palm_code_be/mocks/infra/persistence/postgres/team"
	dto "palm_code_be/src/app/dto/team"
	"palm_code_be/src/interface/response"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockTeamUseCase struct {
	mock.Mock
}

type TeamUseCaseList struct {
	suite.Suite
	mockDTO    *mockDTO.MockTeamDTO
	mockRepo   *mockRepo.MockTeamRepo
	useCase    TeamUsecase
	dtoCreate  *dto.TeamCreateReqDTO
	dtoUpdate  *dto.TeamUpdateReqDTO
	dtoDelete  *dto.TeamDeleteReqDTO
	dtoGet     *dto.TeamGetReqDTO
	dtoGetByID *dto.TeamGetReqByIDDTO
	dtoResp    *dto.TeamRespDTO
	dtoGetResp []*dto.TeamRespDTO
	meta       *response.Meta
}

func (suite *TeamUseCaseList) SetupTest() {
	suite.mockDTO = new(mockDTO.MockTeamDTO)
	suite.mockRepo = new(mockRepo.MockTeamRepo)
	suite.useCase = NewTeamUseCase(suite.mockRepo)

	suite.dtoCreate = &dto.TeamCreateReqDTO{
		Name:           "name",
		Role:           "role",
		Bio:            "bio",
		ProfilePicture: "banner.com",
		CreatedBy:      1,
	}

	suite.dtoUpdate = &dto.TeamUpdateReqDTO{
		ID:             1,
		Name:           "name",
		Role:           "role",
		Bio:            "bio",
		ProfilePicture: "banner.com",
		UpdatedBy:      1,
	}

	suite.dtoDelete = &dto.TeamDeleteReqDTO{
		ID: 1,
	}

	suite.dtoGet = &dto.TeamGetReqDTO{
		Page:        1,
		PerPage:     10,
		CreatedByID: 1,
	}

	suite.dtoGetByID = &dto.TeamGetReqByIDDTO{
		ID:        1,
		CreatedBy: 1,
	}

	suite.dtoResp = &dto.TeamRespDTO{
		ID:             1,
		Name:           "name",
		Role:           "role",
		Bio:            "bio",
		ProfilePicture: "banner.com",
	}

	suite.dtoGetResp = []*dto.TeamRespDTO{
		{
			ID:             1,
			Name:           "name",
			Role:           "role",
			Bio:            "bio",
			ProfilePicture: "banner.com",
		},
	}

	suite.meta = &response.Meta{
		Skip:  0,
		Limit: 10,
		Total: 1,
	}

}

func (u *TeamUseCaseList) TestCreateSuccess() {
	u.mockRepo.Mock.On("Create", u.dtoCreate).Return(nil)
	err := u.useCase.Create(u.dtoCreate)
	u.Equal(nil, err)
}

func (u *TeamUseCaseList) TestCreateFailed() {
	u.mockRepo.Mock.On("Create", u.dtoCreate).Return(errors.New(mock.Anything))
	err := u.useCase.Create(u.dtoCreate)
	u.Equal(errors.New(mock.Anything), err)
}

func (u *TeamUseCaseList) TestGetSuccess() {
	u.mockRepo.Mock.On("Get", u.dtoGet).Return(u.dtoGetResp, u.meta, nil)
	_, _, err := u.useCase.Get(u.dtoGet)
	u.Equal(nil, err)
}

func (u *TeamUseCaseList) TestGetFailed() {
	u.mockRepo.Mock.On("Get", u.dtoGet).Return(nil, nil, errors.New(mock.Anything))
	_, _, err := u.useCase.Get(u.dtoGet)
	u.Equal(errors.New(mock.Anything), err)
}

func (u *TeamUseCaseList) TestGetByIDSuccess() {
	u.mockRepo.Mock.On("GetByID", u.dtoGetByID).Return(u.dtoGetResp, nil)
	_, err := u.useCase.GetByID(u.dtoGetByID)
	u.Equal(nil, err)
}

func (u *TeamUseCaseList) TestGetByIDFailed() {
	u.mockRepo.Mock.On("GetByID", u.dtoGetByID).Return(nil, errors.New(mock.Anything))
	_, err := u.useCase.GetByID(u.dtoGetByID)
	u.Equal(errors.New(mock.Anything), err)
}

func (u *TeamUseCaseList) TestUpdateSuccess() {
	u.mockRepo.Mock.On("Update", u.dtoUpdate).Return(nil)
	err := u.useCase.Update(u.dtoUpdate)
	u.Equal(nil, err)
}

func (u *TeamUseCaseList) TestUpdateFailed() {
	u.mockRepo.Mock.On("Update", u.dtoUpdate).Return(errors.New(mock.Anything))
	err := u.useCase.Update(u.dtoUpdate)
	u.Equal(errors.New(mock.Anything), err)
}

func (u *TeamUseCaseList) TestDeleteSuccess() {
	u.mockRepo.Mock.On("Delete", u.dtoDelete).Return(nil)
	err := u.useCase.Delete(u.dtoDelete)
	u.Equal(nil, err)
}

func (u *TeamUseCaseList) TestDeleteFailed() {
	u.mockRepo.Mock.On("Delete", u.dtoDelete).Return(errors.New(mock.Anything))
	err := u.useCase.Delete(u.dtoDelete)
	u.Equal(errors.New(mock.Anything), err)
}

func TestUsecase(t *testing.T) {
	suite.Run(t, new(TeamUseCaseList))
}
