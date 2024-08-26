package http_client

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWeatherDetails(t *testing.T) {
	mockWeatherClient := new(WeatherClientMock)

	expectedResponse := &WeatherAPIResponse{
		Current: Current{
			LastUpdated: "2021-09-06 22:00",
			TempC:       25.0,
			TempF:       77.0,
		},
	}

	mockWeatherClient.On("WeatherDetails", "São Paulo").Return(expectedResponse, nil)

	actualResponse, err := mockWeatherClient.WeatherDetails("São Paulo")

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)

	mockWeatherClient.AssertExpectations(t)
}
