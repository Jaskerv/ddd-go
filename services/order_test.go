package services

import (
	"testing"

	"github.com/Jaskerv/ddd-go/aggregate"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func initProducts(t testing.TB) []aggregate.Product {
	beer, err := aggregate.NewProduct("Beer", "Hoppy beer", 3.4)
	assert.NoError(t, err)

	peanuts, err := aggregate.NewProduct("Peanuts", "Very nutty", 1.99)
	assert.NoError(t, err)

	wine, err := aggregate.NewProduct("Wine", "Fine wine", 7.50)
	assert.NoError(t, err)

	return []aggregate.Product{
		beer,
		peanuts,
		wine,
	}
}

func TestOrder_NewOrderService(t *testing.T) {
	products := initProducts(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)

	assert.NoError(t, err)

	cust, err := aggregate.NewCustomer("John Doe")

	assert.NoError(t, err)

	err = os.customers.Create(cust)
	assert.NoError(t, err)

	order := []uuid.UUID{
		products[0].GetID(),
	}

	_, err = os.CreateOrder(cust.GetID(), order)
	assert.NoError(t, err)
}
