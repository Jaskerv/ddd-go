package memory

import (
	"testing"

	"github.com/Jaskerv/ddd-go/aggregate"
	"github.com/Jaskerv/ddd-go/domain/customer"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestMemory_GetCustomer(t *testing.T) {
	type testCase struct {
		test        string
		id          uuid.UUID
		expectedErr error
	}

	cust, err := aggregate.NewCustomer("John Doe")
	assert.NoError(t, err)

	id := cust.GetID()

	repo := MemoryRepository{
		customers: map[uuid.UUID]aggregate.Customer{
			id: cust,
		},
	}

	testCases := []testCase{
		{
			test:        "no customer by id",
			id:          uuid.MustParse("1f1dee6b-28d9-460e-9079-bbbc43071bfb"),
			expectedErr: customer.ErrCustomerNotFound,
		},
		{
			test:        "customer by id",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := repo.Get(tc.id)
			assert.ErrorIs(t, err, tc.expectedErr)
		})
	}
}
