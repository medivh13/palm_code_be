package mock_team

import (
	dto "palm_code_be/src/app/dto/team"
	repo "palm_code_be/src/infra/persistence/postgres/team"
	"palm_code_be/src/interface/response"

	"github.com/stretchr/testify/mock"
)

type MockTeamRepo struct {
	mock.Mock
}

func NewMockTeamRepo() *MockTeamRepo {
	return &MockTeamRepo{}
}

var _ repo.TeamRepository = &MockTeamRepo{}

func (m *MockTeamRepo) Create(data *dto.TeamCreateReqDTO) error {
	args := m.Called(data)
	var (
		err error
	)

	if n, ok := args.Get(0).(error); ok {

		err = n
	}

	return err
}

func (m *MockTeamRepo) Update(data *dto.TeamUpdateReqDTO) error {
	args := m.Called(data)
	var (
		err error
	)

	if n, ok := args.Get(0).(error); ok {

		err = n
	}

	return err
}

func (m *MockTeamRepo) Get(data *dto.TeamGetReqDTO) ([]*dto.TeamRespDTO, *response.Meta, error) {
	args := m.Called(data)
	var (
		err  error
		res  []*dto.TeamRespDTO
		meta *response.Meta
	)

	if n, ok := args.Get(0).([]*dto.TeamRespDTO); ok {

		res = n
	}

	if n, ok := args.Get(1).(*response.Meta); ok {

		meta = n
	}

	if n, ok := args.Get(2).(error); ok {

		err = n
	}

	return res, meta, err
}

func (m *MockTeamRepo) GetByID(data *dto.TeamGetReqByIDDTO) (*dto.TeamRespDTO, error) {
	args := m.Called(data)
	var (
		err error
		res *dto.TeamRespDTO
	)

	if n, ok := args.Get(0).(*dto.TeamRespDTO); ok {

		res = n
	}

	if n, ok := args.Get(1).(error); ok {

		err = n
	}

	return res, err
}

func (m *MockTeamRepo) Delete(data *dto.TeamDeleteReqDTO) error {
	args := m.Called(data)
	var (
		err error
	)

	if n, ok := args.Get(0).(error); ok {

		err = n
	}

	return err
}
