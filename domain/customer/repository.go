package customer

import (
	"errors"

	"github.com/Jaskerv/ddd-go/aggregate"
	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound       = errors.New("the customer was not found in the repository")
	ErrFailedToCreateCustomer = errors.New("failed to add the customer")
	ErrUpdateCustomer         = errors.New("failed to update the customer")
)

type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Create(aggregate.Customer) error
	Update(aggregate.Customer) error
}
