package http_client

import (
	"github.com/stretchr/testify/mock"
)

type WeatherClientMock struct {
	mock.Mock
}

func (m *WeatherClientMock) AddressDetails(zipCode string) (*ViaCepResponse, error) {
	args := m.Called(zipCode)
	return args.Get(0).(*ViaCepResponse), args.Error(1)
}

func (m *WeatherClientMock) WeatherDetails(locale string) (*WeatherAPIResponse, error) {
	args := m.Called(locale)
	return args.Get(0).(*WeatherAPIResponse), args.Error(1)
}
