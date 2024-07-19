package mock_firestore

import (
	fSore "palm_code_be/src/infra/integration/firestorage"

	"github.com/stretchr/testify/mock"
)

type MockFstoreRepo struct {
	mock.Mock
}

func NewMockFstore() *MockFstoreRepo {
	return &MockFstoreRepo{}
}

var _ fSore.FireStorage = &MockFstoreRepo{}

func (m *MockFstoreRepo) Upload(bucketName, objectName, filePath string) (string, error) {
	args := m.Called(bucketName, objectName, filePath)
	var (
		err error
		res string
	)

	if n, ok := args.Get(0).(string); ok {

		res = n
	}

	if n, ok := args.Get(1).(error); ok {

		err = n
	}

	return res, err
}
