package handler

import (
	"encoding/json"
	"github.com/mayckol/stress-test/http_client"
	"github.com/mayckol/stress-test/utils"
	"net/http"
	"strconv"
	"strings"
)

type WeatherHandler struct {
	ViaCepClient  http_client.ViaCepClientInterface
	WeatherClient http_client.WeatherClientInterface
}

func NewWeatherHandler(viaCepClient http_client.ViaCepClientInterface, weatherClient http_client.WeatherClientInterface) *WeatherHandler {
	return &WeatherHandler{
		ViaCepClient:  viaCepClient,
		WeatherClient: weatherClient,
	}
}

func (h *WeatherHandler) Weather(w http.ResponseWriter, r *http.Request) {
	pathParam := r.PathValue("zipCode")
	// The current net/http library does not support path parameters, when executing the test, the path parameter is not being passed correctly.
	if pathParam == "" {
		pathParts := strings.Split(r.URL.Path, "/")
		if len(pathParts) < 3 {
			w.WriteHeader(http.StatusUnprocessableEntity)
			if _, err := w.Write([]byte("invalid zipcode")); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}
		pathParam = pathParts[2]
	}
	serializedZipCode := utils.ZipCode(pathParam)
	if !serializedZipCode.IsValid() {
		w.WriteHeader(http.StatusUnprocessableEntity)
		if _, err := w.Write([]byte("invalid zipcode")); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}
	addressDetails, err := h.ViaCepClient.AddressDetails(serializedZipCode.Raw())
	if err != nil {
		slices := strings.Split(err.Error(), "@")
		if len(slices) != 2 {
			if _, err := w.Write([]byte("error making request")); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
		message, statusCode := slices[0], slices[1]
		code, _ := strconv.Atoi(statusCode)
		w.WriteHeader(code)
		w.Write([]byte(message))
		return
	}

	if addressDetails == nil {
		w.WriteHeader(http.StatusNotFound)
		if _, err := w.Write([]byte("can not find zipcode")); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	weatherAPIResponse, err := h.WeatherClient.WeatherDetails(addressDetails.Localidade)
	if err != nil {
		slices := strings.Split(err.Error(), "@")
		if len(slices) != 2 {
			if _, err := w.Write([]byte("error making request")); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
		message, statusCode := slices[0], slices[1]
		code, _ := strconv.Atoi(statusCode)
		w.WriteHeader(code)
		w.Write([]byte(message))
		return
	}

	if weatherAPIResponse == nil {
		w.WriteHeader(http.StatusNotFound)
		if _, err := w.Write([]byte("can not find zipcode")); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	tempK := utils.RoundToDecimal(weatherAPIResponse.Current.TempC+273.15, 1)
	weatherResponse := http_client.WeatherResponse{
		TempC: utils.RoundToDecimal(weatherAPIResponse.Current.TempC, 1),
		TempF: utils.RoundToDecimal(weatherAPIResponse.Current.TempF, 1),
		TempK: tempK,
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(weatherResponse); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
