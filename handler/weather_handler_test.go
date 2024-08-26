package handler_test

import (
	"errors"
	"github.com/mayckol/stress-test/handler"
	"github.com/mayckol/stress-test/http_client"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestWeatherHandler_Weather(t *testing.T) {
	mockViaCepClient := new(http_client.ViaCepClientMock)
	mockWeatherClient := new(http_client.WeatherClientMock)

	weatherHandler := handler.NewWeatherHandler(mockViaCepClient, mockWeatherClient)
	validZipCode := "01001000"
	invalidZipCode := "1234"
	validURLPath := "/weather/" + validZipCode
	invalidURLPath := "/weather/" + invalidZipCode
	tests := []struct {
		name            string
		urlPath         string
		zipCode         string
		viaCepResponse  *http_client.ViaCepResponse
		weatherResponse *http_client.WeatherAPIResponse
		viaCepError     error
		weatherError    error
		expectedStatus  int
		expectedBody    string
	}{
		{
			name:           "Valid request",
			urlPath:        validURLPath,
			zipCode:        validZipCode,
			viaCepResponse: &http_client.ViaCepResponse{Localidade: "São Paulo"},
			weatherResponse: &http_client.WeatherAPIResponse{
				Current: http_client.Current{
					LastUpdated: "2021-09-07 22:00",
					TempC:       25.0,
					TempF:       77.0,
				},
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"temp_c":25.0,"temp_f":77.0,"temp_k":298.2}`,
		},
		{
			name:           "Invalid zip code",
			urlPath:        invalidURLPath,
			zipCode:        invalidZipCode,
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody:   "invalid zipcode",
		},
		{
			name:           "Weather API error",
			urlPath:        validURLPath,
			zipCode:        validZipCode,
			viaCepResponse: &http_client.ViaCepResponse{Localidade: "São Paulo"},
			weatherError:   errors.New("error getting weather@500"),
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "error getting weather",
		},
		{
			name:           "Address not found",
			urlPath:        validURLPath,
			zipCode:        validZipCode,
			viaCepResponse: nil,
			viaCepError:    mockViaCepError("can not find zipcode", http.StatusNotFound),
			expectedStatus: http.StatusNotFound,
			expectedBody:   "can not find zipcode",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.viaCepResponse != nil || tt.viaCepError != nil {
				mockViaCepClient.On("AddressDetails", tt.zipCode).
					Return(tt.viaCepResponse, tt.viaCepError).
					Once()
			}
			if tt.weatherResponse != nil || tt.weatherError != nil {
				mockWeatherClient.On("WeatherDetails", "São Paulo").
					Return(tt.weatherResponse, tt.weatherError).
					Once()
			}

			req, err := http.NewRequest(http.MethodGet, tt.urlPath, nil)
			assert.NoError(t, err)

			rr := httptest.NewRecorder()

			weatherHandler.Weather(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)

			if tt.expectedStatus == http.StatusOK {
				assert.JSONEq(t, tt.expectedBody, rr.Body.String())
			} else {
				assert.Equal(t, tt.expectedBody, rr.Body.String())
			}

			mockViaCepClient.AssertExpectations(t)
			mockWeatherClient.AssertExpectations(t)
		})
	}
}

func mockViaCepError(message string, statusCode int) error {
	return errors.New(message + "@" + strconv.Itoa(statusCode))
}

func mockWeatherError(message string, statusCode int) error {
	return errors.New(message + "@" + strconv.Itoa(statusCode))
}
