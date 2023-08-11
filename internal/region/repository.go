package region

import (
	"branches-cli/internal/io"
	"errors"
)

type StorageAdapter interface {
	GetAll() ([]Region, error)
	GetById(id int) (*Region, error)
}

type Repository struct {
	manipulator io.FileIOHandler[Region]
}

func BuildRepository() *Repository {
	serializer := io.NewJsonSerializer[Region](Region{})
	handler := io.NewJsonIOHandler[Region](StoragePath, serializer)
	return NewRepository(handler)
}
func NewRepository(handler io.FileIOHandler[Region]) *Repository {
	return &Repository{
		manipulator: handler,
	}
}

func (r *Repository) GetAll() ([]Region, error) {
	regions, err := r.manipulator.Read()
	if err != nil {
		return nil, err
	}

	return regions, nil
}
func (r *Repository) GetById(id int) (*Region, error) {
	branches, err := r.GetAll()
	if err != nil {
		return nil, err
	}

	for _, branch := range branches {
		if branch.Id == id {
			return &branch, nil
		}
	}

	return nil, errors.New("branch Not Found")
}
