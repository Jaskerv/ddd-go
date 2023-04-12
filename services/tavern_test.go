package services

import (
	"testing"

	"github.com/Jaskerv/ddd-go/aggregate"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_Tavern(t *testing.T) {
	products := initProducts(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)

	assert.NoError(t, err)
	tavern, err := NewTavern(WithOrderService(os))
	assert.NoError(t, err)

	cust, err := aggregate.NewCustomer("John Doe")
	assert.NoError(t, err)

	err = os.customers.Create(cust)
	assert.NoError(t, err)

	order := []uuid.UUID{
		products[0].GetID(),
		products[2].GetID(),
	}

	err = tavern.Order(cust.GetID(), order)
	assert.NoError(t, err)
}
