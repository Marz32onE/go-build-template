package storage

import data "github.com/marz32one/go-build-template/pkg/storage/data"

type ResourceRepository interface {
	GetAll() ([]data.Resource, error)
	GetByName(name string) (*data.Resource, error)
	Create(resource *data.Resource) error
	Update(resource *data.Resource) error
	Delete(name string) error
}
