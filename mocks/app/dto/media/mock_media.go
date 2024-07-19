package mock_media

import (
	dto "palm_code_be/src/app/dto/media"

	"github.com/stretchr/testify/mock"
)

type MockMediasDTO struct {
	mock.Mock
}

func NewMockMediasDTO() *MockMediasDTO {
	return &MockMediasDTO{}
}

var _ dto.MediaDTOInterface = &MockMediasDTO{}

func (m *MockMediasDTO) Validate() error {
	args := m.Called()
	var err error
	if n, ok := args.Get(0).(error); ok {
		err = n
		return err
	}

	return nil
}
