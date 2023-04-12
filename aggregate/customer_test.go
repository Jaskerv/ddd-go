package aggregate_test

import (
	"testing"

	"github.com/Jaskerv/ddd-go/aggregate"
	"github.com/stretchr/testify/assert"
)

func TestCustomer_NewCustomer(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Empty name validation",
			name:        "",
			expectedErr: aggregate.ErrInvalidPerson,
		},
		{
			test:        "Valid name",
			name:        "John Doe",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.NewCustomer(tc.name)
			assert.ErrorIs(t, err, tc.expectedErr)
		})
	}
}
