package application_test

import (
	"github.com/filipedtristao/hexagonal-architecture/application"
	"github.com/stretchr/testify/require"
	"github.com/satori/go.uuid"
	"testing"
)

func TestProductEnable(t *testing.T) {
	product := application.Product{
		Id:    "1",
		Name:  "Product 1",
		Status: application.DISABLED,
		Price: 10.0,
	}

	err := product.Enable()
	require.Nil(t, err)
	require.Equal(t, application.ENABLED, product.Status)

	product.Price = 0.0

	err = product.Enable()

	require.Equal(t, "the product price must be greater than zero", err.Error())
}

func TestProductDisable(t *testing.T) {
	product := application.Product{
		Id:    "1",
		Name:  "Product 1",
		Status: application.ENABLED,
		Price: 0.0,
	}

	err := product.Disable()
	require.Nil(t, err)
	require.Equal(t, application.DISABLED, product.Status)

	product.Price = 10.0

	err = product.Disable()

	require.Equal(t, "the product price must be zero", err.Error())
}

func TestProductIsValid(t *testing.T) {
	product := application.Product{
		Id:    uuid.NewV4().String(),
		Name:  "Product 1",
		Status: application.ENABLED,
		Price: 10.0,
	}

	isValid, err := product.IsValid()
	require.True(t, isValid)
	require.Nil(t, err)

	product.Id = "invalid-uuid"
	isValid, err = product.IsValid()
	
	require.False(t, isValid)
	require.Equal(t, "Id: invalid-uuid does not validate as uuidv4", err.Error())

	product.Id = uuid.NewV4().String()
	product.Name = ""

	isValid, err = product.IsValid()

	require.False(t, isValid)
	require.Equal(t, "Name: non zero value required", err.Error())
	
	product.Name = "Product 1"
	product.Status = "invalid-status"

	isValid, err = product.IsValid()

	require.Equal(t, "the product status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	product.Price = -10

	isValid, err = product.IsValid()

	require.False(t, isValid)
	require.Equal(t, "the product price must be greater than zero", err.Error())
}