package user

import (
	"errors"
	"testing"
	mockDTO "palm_code_be/mocks/app/dto/user"

	mockRepo "palm_code_be/mocks/infra/persistence/postgres/user"
	dto "palm_code_be/src/app/dto/user"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockUserUseCase struct {
	mock.Mock
}

type UserUseCaseList struct {
	suite.Suite
	mockDTO        *mockDTO.MockUsersDTO
	mockRepo       *mockRepo.MockUsersRepo
	useCase        UserUCInterface
	dtoRegisterReq *dto.RegisterReqDTO
	dtoLoginReq    *dto.LoginReqDTO
	dtoResp        *dto.RegisterRespDTO
}

func (suite *UserUseCaseList) SetupTest() {
	suite.mockDTO = new(mockDTO.MockUsersDTO)
	suite.mockRepo = new(mockRepo.MockUsersRepo)
	suite.useCase = NewUserUseCase(suite.mockRepo)

	suite.dtoRegisterReq = &dto.RegisterReqDTO{
		Email: "jody.almaida@gmail.com",
	}

	suite.dtoLoginReq = &dto.LoginReqDTO{
		Email: "jody.almaida@gmail.com",
	}

	suite.dtoResp = &dto.RegisterRespDTO{
		Token: "asdfghjkl",
	}
}

func (u *UserUseCaseList) TestRegisterSuccess() {
	u.mockRepo.Mock.On("Register", u.dtoRegisterReq).Return(u.dtoResp, nil)
	_, err := u.useCase.Register(u.dtoRegisterReq)
	u.Equal(nil, err)
}

func (u *UserUseCaseList) TestRegisterFailed() {
	u.mockRepo.Mock.On("Register", u.dtoRegisterReq).Return(nil, errors.New(mock.Anything))
	_, err := u.useCase.Register(u.dtoRegisterReq)
	u.Equal(errors.New(mock.Anything), err)
}

func (u *UserUseCaseList) TestLoginSuccess() {
	u.mockRepo.Mock.On("Login", u.dtoLoginReq).Return(u.dtoResp, nil)
	_, err := u.useCase.Login(u.dtoLoginReq)
	u.Equal(nil, err)
}

func (u *UserUseCaseList) TestLoginFailed() {
	u.mockRepo.Mock.On("Login", u.dtoLoginReq).Return(nil, errors.New(mock.Anything))
	_, err := u.useCase.Login(u.dtoLoginReq)
	u.Equal(errors.New(mock.Anything), err)
}

func TestUsecase(t *testing.T) {
	suite.Run(t, new(UserUseCaseList))
}
