// Package memory is a in-memory implementation of Customer Repository.
package memory

import (
	"fmt"
	"sync"

	"github.com/Jaskerv/ddd-go/aggregate"
	"github.com/Jaskerv/ddd-go/domain/product"
	"github.com/google/uuid"
)

var _ product.ProductRepository = (*MemoryRepository)(nil)

type MemoryRepository struct {
	products map[uuid.UUID]aggregate.Product
	mu       sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		products: make(map[uuid.UUID]aggregate.Product),
	}
}

func (mr *MemoryRepository) GetAll() ([]aggregate.Product, error) {
	var products []aggregate.Product

	for _, product := range mr.products {
		products = append(products, product)
	}

	return products, nil
}

func (mr *MemoryRepository) GetByID(id uuid.UUID) (aggregate.Product, error) {
	if p, ok := mr.products[id]; ok {
		return p, nil
	}

	return aggregate.Product{}, product.ErrProductNotFound
}

func (mr *MemoryRepository) Create(p aggregate.Product) error {
	mr.mu.Lock()
	defer mr.mu.Unlock()

	if _, ok := mr.products[p.GetID()]; ok {
		return fmt.Errorf("product does already exists :%w", product.ErrCreate)
	}

	mr.products[p.GetID()] = p
	return nil
}

func (mr *MemoryRepository) Update(p aggregate.Product) error {
	mr.mu.Lock()
	defer mr.mu.Unlock()
	// Make sure product is exists in repo
	if _, ok := mr.products[p.GetID()]; !ok {
		return fmt.Errorf("product does not exists :%w", product.ErrProductNotFound)
	}

	mr.products[p.GetID()] = p
	return nil
}

func (mr *MemoryRepository) Delete(id uuid.UUID) error {
	mr.mu.Lock()
	defer mr.mu.Unlock()
	// Make sure product is exists in repo
	if _, ok := mr.products[id]; !ok {
		return fmt.Errorf("product does not exists :%w", product.ErrProductNotFound)
	}

	delete(mr.products, id)
	return nil
}
