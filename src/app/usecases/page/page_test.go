package page

import (
	"errors"
	mockDTO "palm_code_be/mocks/app/dto/page"
	"testing"

	mockRepo "palm_code_be/mocks/infra/persistence/postgres/page"
	dto "palm_code_be/src/app/dto/pages"
	"palm_code_be/src/interface/response"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockPageUseCase struct {
	mock.Mock
}

type PageUseCaseList struct {
	suite.Suite
	mockDTO    *mockDTO.MockPagesDTO
	mockRepo   *mockRepo.MockPageRepo
	useCase    PageUsecase
	dtoCreate  *dto.PageCreateReqDTO
	dtoUpdate  *dto.PageUpdateReqDTO
	dtoDelete  *dto.PageDeleteReqDTO
	dtoGet     *dto.PagesGetReqDTO
	dtoGetByID *dto.PagesGetReqByIDDTO
	dtoResp    *dto.PageRespDTO
	dtoGetResp []*dto.PageRespDTO
	meta       *response.Meta
}

func (suite *PageUseCaseList) SetupTest() {
	suite.mockDTO = new(mockDTO.MockPagesDTO)
	suite.mockRepo = new(mockRepo.MockPageRepo)
	suite.useCase = NewPageUseCase(suite.mockRepo)

	suite.dtoCreate = &dto.PageCreateReqDTO{
		Title:       "title",
		Slug:        "title",
		BannerMedia: "banner.com",
		Content:     "content",
		CreatedBy:   1,
	}

	suite.dtoUpdate = &dto.PageUpdateReqDTO{
		ID:          1,
		Title:       "title",
		Slug:        "title",
		BannerMedia: "banner.com",
		Content:     "content",
		UpdatedBy:   1,
	}

	suite.dtoDelete = &dto.PageDeleteReqDTO{
		ID: 1,
	}

	suite.dtoGet = &dto.PagesGetReqDTO{
		Page:        1,
		PerPage:     10,
		CreatedByID: 1,
	}

	suite.dtoGetByID = &dto.PagesGetReqByIDDTO{
		ID:        1,
		CreatedBy: 1,
	}

	suite.dtoResp = &dto.PageRespDTO{
		Title:       "title",
		Slug:        "title",
		BannerMedia: "banner.com",
		Content:     "content",
	}

	suite.dtoGetResp = []*dto.PageRespDTO{
		{
			Title:       "title",
			Slug:        "title",
			BannerMedia: "banner.com",
			Content:     "content",
		},
	}

	suite.meta = &response.Meta{
		Skip:  0,
		Limit: 10,
		Total: 1,
	}

}

func (u *PageUseCaseList) TestCreateSuccess() {
	u.mockRepo.Mock.On("Create", u.dtoCreate).Return(nil)
	err := u.useCase.Create(u.dtoCreate)
	u.Equal(nil, err)
}

func (u *PageUseCaseList) TestCreateFailed() {
	u.mockRepo.Mock.On("Create", u.dtoCreate).Return(errors.New(mock.Anything))
	err := u.useCase.Create(u.dtoCreate)
	u.Equal(errors.New(mock.Anything), err)
}

func (u *PageUseCaseList) TestGetSuccess() {
	u.mockRepo.Mock.On("Get", u.dtoGet).Return(u.dtoGetResp, u.meta, nil)
	_, _, err := u.useCase.Get(u.dtoGet)
	u.Equal(nil, err)
}

func (u *PageUseCaseList) TestGetFailed() {
	u.mockRepo.Mock.On("Get", u.dtoGet).Return(nil, nil, errors.New(mock.Anything))
	_, _, err := u.useCase.Get(u.dtoGet)
	u.Equal(errors.New(mock.Anything), err)
}

func (u *PageUseCaseList) TestGetByIDSuccess() {
	u.mockRepo.Mock.On("GetByID", u.dtoGetByID).Return(u.dtoGetResp, nil)
	_, err := u.useCase.GetByID(u.dtoGetByID)
	u.Equal(nil, err)
}

func (u *PageUseCaseList) TestGetByIDFailed() {
	u.mockRepo.Mock.On("GetByID", u.dtoGetByID).Return(nil, errors.New(mock.Anything))
	_, err := u.useCase.GetByID(u.dtoGetByID)
	u.Equal(errors.New(mock.Anything), err)
}

func (u *PageUseCaseList) TestUpdateSuccess() {
	u.mockRepo.Mock.On("Update", u.dtoUpdate).Return(nil)
	err := u.useCase.Update(u.dtoUpdate)
	u.Equal(nil, err)
}

func (u *PageUseCaseList) TestUpdateFailed() {
	u.mockRepo.Mock.On("Update", u.dtoUpdate).Return(errors.New(mock.Anything))
	err := u.useCase.Update(u.dtoUpdate)
	u.Equal(errors.New(mock.Anything), err)
}

func (u *PageUseCaseList) TestDeleteSuccess() {
	u.mockRepo.Mock.On("Delete", u.dtoDelete).Return(nil)
	err := u.useCase.Delete(u.dtoDelete)
	u.Equal(nil, err)
}

func (u *PageUseCaseList) TestDeleteFailed() {
	u.mockRepo.Mock.On("Delete", u.dtoDelete).Return(errors.New(mock.Anything))
	err := u.useCase.Delete(u.dtoDelete)
	u.Equal(errors.New(mock.Anything), err)
}

func TestUsecase(t *testing.T) {
	suite.Run(t, new(PageUseCaseList))
}
