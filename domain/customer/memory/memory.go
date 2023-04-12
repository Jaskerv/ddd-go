// Package memory is a in-memory implementation of Customer Repository.
package memory

import (
	"fmt"
	"sync"

	"github.com/Jaskerv/ddd-go/aggregate"
	"github.com/Jaskerv/ddd-go/domain/customer"
	"github.com/google/uuid"
)

var _ customer.CustomerRepository = (*MemoryRepository)(nil)

type MemoryRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	mu        sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

func (mr *MemoryRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	if customer, ok := mr.customers[id]; ok {
		return customer, nil
	}
	return aggregate.Customer{}, customer.ErrCustomerNotFound
}

func (mr *MemoryRepository) Create(c aggregate.Customer) error {
	if mr.customers == nil {
		mr.mu.Lock()
		mr.customers = make(map[uuid.UUID]aggregate.Customer)
		mr.mu.Unlock()
	}

	// Make sure customer does not exist in repo
	if _, ok := mr.customers[c.GetID()]; ok {
		return fmt.Errorf("customer already exists :%w", customer.ErrFailedToCreateCustomer)
	}

	mr.mu.Lock()
	defer mr.mu.Unlock()

	mr.customers[c.GetID()] = c

	return nil
}

func (mr *MemoryRepository) Update(c aggregate.Customer) error {
	// Make sure customer is exists in repo
	if _, ok := mr.customers[c.GetID()]; !ok {
		return fmt.Errorf("customer does not exists :%w", customer.ErrUpdateCustomer)
	}

	mr.mu.Lock()
	defer mr.mu.Unlock()

	mr.customers[c.GetID()] = c

	return nil
}
