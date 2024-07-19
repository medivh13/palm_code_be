package mock_media

import (
	dto "palm_code_be/src/app/dto/media"
	repo "palm_code_be/src/infra/persistence/postgres/media"
	"palm_code_be/src/interface/response"

	"github.com/stretchr/testify/mock"
)

type MockMediaRepo struct {
	mock.Mock
}

func NewMockMediaRepo() *MockMediaRepo {
	return &MockMediaRepo{}
}

var _ repo.MediaRepository = &MockMediaRepo{}

func (m *MockMediaRepo) Create(data *dto.MediaCreateReqDTO) error {
	args := m.Called(data)
	var (
		err error
	)

	if n, ok := args.Get(0).(error); ok {

		err = n
	}

	return err
}

func (m *MockMediaRepo) Get(data *dto.MediaGetReqDTO) ([]*dto.MediaRespDTO, *response.Meta, error) {
	args := m.Called(data)
	var (
		err  error
		res  []*dto.MediaRespDTO
		meta *response.Meta
	)

	if n, ok := args.Get(0).([]*dto.MediaRespDTO); ok {

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

func (m *MockMediaRepo) GetByID(data *dto.MediaGetReqByIDDTO) (*dto.MediaRespDTO, error) {
	args := m.Called(data)
	var (
		err error
		res *dto.MediaRespDTO
	)

	if n, ok := args.Get(0).(*dto.MediaRespDTO); ok {

		res = n
	}

	if n, ok := args.Get(1).(error); ok {

		err = n
	}

	return res, err
}
