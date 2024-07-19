package mock_team

import (
	dto "palm_code_be/src/app/dto/team"

	"github.com/stretchr/testify/mock"
)

type MockTeamDTO struct {
	mock.Mock
}

func NewMockTeamsDTO() *MockTeamDTO {
	return &MockTeamDTO{}
}

var _ dto.TeamDTOInterface = &MockTeamDTO{}

func (m *MockTeamDTO) Validate() error {
	args := m.Called()
	var err error
	if n, ok := args.Get(0).(error); ok {
		err = n
		return err
	}

	return nil
}
