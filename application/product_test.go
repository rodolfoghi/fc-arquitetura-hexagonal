package application_test

import (
	"github.com/codeedu/go-hexagonal/application"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Product Test"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()

	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()

	require.Equal(t, "The price must be greater than zero to enable the product", err.Error())
}
