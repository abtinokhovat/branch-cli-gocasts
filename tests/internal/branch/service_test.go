package branch_test

import (
	"branches-cli/internal/branch"
	"errors"
)

var storage = []branch.Branch{
	{},
}

type MockBranchRepository struct {
	branch.StorageAdapter
}

func (r *MockBranchRepository) GetById(id int) (*branch.Branch, error) {
	for _, br := range storage {
		if br.Id == id {
			return &br, nil
		}
	}
	return nil, errors.New("branch not found")
}
func (r *MockBranchRepository) GetAll() ([]branch.Branch, error) {
	return storage, nil
}
func (r *MockBranchRepository) Create(brn *branch.Branch) error {
	storage = append(storage, *brn)
	return nil
}
func (r *MockBranchRepository) UpdateById(brn *branch.Branch) error {
	old, err := r.GetById(brn.Id)
	if err != nil {
		return err
	}

	// Update data
	old.Name = brn.Name
	old.Phone = brn.Phone
	old.CreateDate = brn.CreateDate
	old.NumberOfEmployees = brn.NumberOfEmployees

	return nil
}
