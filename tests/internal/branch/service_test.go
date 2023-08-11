package branch_test

import (
	"branches-cli/internal/branch"
	"branches-cli/internal/region"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var storage = []branch.Branch{
	{Id: 1, Name: "Branch 1", RegionId: 1},
	{Id: 2, Name: "Branch 2", RegionId: 2},
	{Id: 3, Name: "Branch 3", RegionId: 1},
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
	i, err := r.findIndexById(brn.Id)
	if err != nil {
		return err
	}

	// Update data in storage
	storage[i].Name = brn.Name
	storage[i].Phone = brn.Phone
	storage[i].CreateDate = brn.CreateDate
	storage[i].NumberOfEmployees = brn.NumberOfEmployees
	storage[i].RegionId = brn.RegionId

	return nil
}
func (r *MockBranchRepository) findIndexById(id int) (int, error) {
	for i, br := range storage {
		if br.Id == id {
			return i, nil
		}
	}
	return -1, errors.New("branch not found")
}

func TestService_ListBranchesInRegion(t *testing.T) {
	t.Run("ordinary", func(t *testing.T) {
		// 1. setup
		mockRepo := &MockBranchRepository{}
		service := branch.NewBranchService(mockRepo)

		// 2. execution
		branches, err := service.ListBranchesInRegion(region.Region{Id: 1})
		assert.NoError(t, err, "Error listing branches in region")

		// 3. assertion
		assert.Len(t, branches, 2, "Unexpected number of branches")
	})
}

func TestService_GetBranchDetail(t *testing.T) {
	t.Run("ordinary", func(t *testing.T) {
		// 1. setup
		mockRepo := &MockBranchRepository{}
		service := branch.NewBranchService(mockRepo)

		// 2. execution
		brn, err := service.GetBranchDetail(2)
		assert.NoError(t, err, "Error getting branch detail")

		// 3. assertion
		assert.Equal(t, "Branch 2", brn.Name, "Unexpected branch name")
	})
}

func TestService_CreateBranch(t *testing.T) {
	t.Run("ordinary", func(t *testing.T) {
		// 1. setup
		mockRepo := &MockBranchRepository{}
		service := branch.NewBranchService(mockRepo)

		newBranch := &branch.Branch{Id: 4, Name: "New Branch", RegionId: 3}

		// 2. execution
		err := service.CreateBranch(newBranch)
		assert.NoError(t, err, "Error creating branch")

		// 3. assertion
		assert.Len(t, storage, 4, "Unexpected number of branches in storage")
	})
}

func TestService_EditBranch(t *testing.T) {
	testStruct := []struct {
		name string
		data *branch.Branch
	}{
		{
			name: "Update Name",
			data: &branch.Branch{
				Id:       2,
				Name:     "Edited Branch",
				RegionId: 2,
			},
		},
		{
			name: "Update Region",
			data: &branch.Branch{
				Id:       1,
				Name:     "Branch 1",
				RegionId: 3,
			},
		},
	}
	for _, tc := range testStruct {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			mockRepo := &MockBranchRepository{}
			service := branch.NewBranchService(mockRepo)

			// 2. execution
			err := service.EditBranch(tc.data)
			assert.NoError(t, err, "Error editing branch")

			// 3. assertion
			assert.Equal(t, storage[tc.data.Id-1], *tc.data, fmt.Sprintf("%s failed", tc.name))
		})
	}

}
