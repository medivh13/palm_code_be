package mock_page

import (
	dto "palm_code_be/src/app/dto/pages"

	"github.com/stretchr/testify/mock"
)

type MockPagesDTO struct {
	mock.Mock
}

func NewMockPagesDTO() *MockPagesDTO {
	return &MockPagesDTO{}
}

var _ dto.PagesDTOInterface = &MockPagesDTO{}

func (m *MockPagesDTO) Validate() error {
	args := m.Called()
	var err error
	if n, ok := args.Get(0).(error); ok {
		err = n
		return err
	}

	return nil
}