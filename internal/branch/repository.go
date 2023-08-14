package branch

import (
	"branches-cli/internal/io"
	"errors"
)

type StorageAdapter interface {
	GetById(id int) (*Branch, error)
	GetAll() ([]Branch, error)
	Create(*Branch) error
	UpdateById(*Branch) error
}
type Repository struct {
	manipulator io.FileIOHandler[Branch]
}

func BuildRepository() *Repository {
	serializer := io.NewJsonSerializer[Branch]()
	handler := io.NewJsonIOHandler[Branch](StoragePath, serializer)
	return NewRepository(handler)
}
func NewRepository(handler io.FileIOHandler[Branch]) *Repository {
	return &Repository{
		manipulator: handler,
	}
}

func (r *Repository) GetById(id int) (*Branch, error) {
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
func (r *Repository) GetAll() ([]Branch, error) {
	branches, err := r.manipulator.Read()
	if err != nil {
		return nil, err
	}

	return branches, nil
}
func (r *Repository) Create(branch *Branch) error {
	err := r.manipulator.WriteOne(*branch)
	if err != nil {
		return err
	}
	return nil
}
func (r *Repository) UpdateById(branch *Branch) error {
	values, err := r.GetAll()

	if err != nil {
		return err
	}

	for i, value := range values {
		if value.Id == branch.Id {
			if branch.Name != "" {
				values[i].Name = branch.Name
			}
			if branch.Phone != "" {
				values[i].Phone = branch.Phone
			}
			if branch.NumberOfEmployees > 0 {
				values[i].NumberOfEmployees = branch.NumberOfEmployees
			}
			if branch.RegionId > 0 {
				values[i].RegionId = branch.RegionId
			}
			//values[i] = *branch
		}
	}

	// rewrite all data again to file
	err = r.manipulator.DeleteAndWrite(values)
	if err != nil {
		return err
	}

	return nil
}
