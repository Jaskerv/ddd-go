package product

import (
	"errors"

	"github.com/Jaskerv/ddd-go/aggregate"
	"github.com/google/uuid"
)

var (
	ErrProductNotFound = errors.New("product not found")
	ErrCreate          = errors.New("can't create")
)

type ProductRepository interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(uuid.UUID) (aggregate.Product, error)
	Create(aggregate.Product) error
	Update(aggregate.Product) error
	Delete(uuid.UUID) error
}
