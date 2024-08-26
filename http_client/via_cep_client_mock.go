package http_client

import (
	"github.com/stretchr/testify/mock"
)

type ViaCepClientMock struct {
	mock.Mock
}

func (m *ViaCepClientMock) AddressDetails(zipCode string) (*ViaCepResponse, error) {
	args := m.Called(zipCode)
	return args.Get(0).(*ViaCepResponse), args.Error(1)
}
