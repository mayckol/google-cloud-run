package http_client

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddressDetails(t *testing.T) {
	mockViaCepClient := new(ViaCepClientMock)

	expectedResponse := &ViaCepResponse{
		Cep:        "01001-000",
		Logradouro: "Praça da Sé",
		Localidade: "São Paulo",
		Uf:         "SP",
	}

	mockViaCepClient.On("AddressDetails", "01001-000").Return(expectedResponse, nil)

	actualResponse, err := mockViaCepClient.AddressDetails("01001-000")

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)

	mockViaCepClient.AssertExpectations(t)
}
