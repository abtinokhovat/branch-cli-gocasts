package region

import (
	"errors"

	io "github.com/abtinokhovat/file-handler-go"
)

var defaultBranch = Region{1, "Tehran"}

type StorageAdapter interface {
	GetAll() ([]Region, error)
	GetById(id int) (*Region, error)
}

type Repository struct {
	manipulator io.FileIOHandler[Region]
}

func BuildRepository() *Repository {
	serializer := io.NewJsonSerializer[Region]()
	handler := io.NewJsonIOHandler[Region](StoragePath, serializer)
	return NewRepository(handler)
}
func NewRepository(handler io.FileIOHandler[Region]) *Repository {
	// handling empty region list
	data, _ := handler.Read()
	if len(data) == 0 {
		_ = handler.WriteOne(defaultBranch)
	}

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

	return nil, errors.New("region not found")
}
