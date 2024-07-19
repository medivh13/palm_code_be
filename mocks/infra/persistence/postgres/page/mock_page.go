package mock_page

import (
	dto "palm_code_be/src/app/dto/pages"
	repo "palm_code_be/src/infra/persistence/postgres/pages"
	"palm_code_be/src/interface/response"

	"github.com/stretchr/testify/mock"
)

type MockPageRepo struct {
	mock.Mock
}

func NewMockPageRepo() *MockPageRepo {
	return &MockPageRepo{}
}

var _ repo.PageRepository = &MockPageRepo{}

func (m *MockPageRepo) Create(data *dto.PageCreateReqDTO) error {
	args := m.Called(data)
	var (
		err error
	)

	if n, ok := args.Get(0).(error); ok {

		err = n
	}

	return err
}

func (m *MockPageRepo) Update(data *dto.PageUpdateReqDTO) error {
	args := m.Called(data)
	var (
		err error
	)

	if n, ok := args.Get(0).(error); ok {

		err = n
	}

	return err
}

func (m *MockPageRepo) Get(data *dto.PagesGetReqDTO) ([]*dto.PageRespDTO, *response.Meta, error) {
	args := m.Called(data)
	var (
		err  error
		res  []*dto.PageRespDTO
		meta *response.Meta
	)

	if n, ok := args.Get(0).([]*dto.PageRespDTO); ok {

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

func (m *MockPageRepo) GetByID(data *dto.PagesGetReqByIDDTO) (*dto.PageRespDTO, error) {
	args := m.Called(data)
	var (
		err error
		res *dto.PageRespDTO
	)

	if n, ok := args.Get(0).(*dto.PageRespDTO); ok {

		res = n
	}

	if n, ok := args.Get(1).(error); ok {

		err = n
	}

	return res, err
}

func (m *MockPageRepo) Delete(data *dto.PageDeleteReqDTO) error {
	args := m.Called(data)
	var (
		err error
	)

	if n, ok := args.Get(0).(error); ok {

		err = n
	}

	return err
}
