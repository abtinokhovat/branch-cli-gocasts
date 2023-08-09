package branch

import "branches-cli/internal/region"

type Service struct {
	adp StorageAdapter
}

func NewBranchService(adp StorageAdapter) *Service {
	return &Service{
		adp: adp,
	}
}

func (s *Service) ListBranchesInRegion(region region.Region) ([]Branch, error) {
	// Implement the logic to list branches in a specific region
	panic("implement ListBranchesInRegion")
}

func (s *Service) GetBranchDetail(id int) (*Branch, error) {
	// Implement the logic to get branch details by ID
	panic("implement GetBranchDetail")
}

func (s *Service) CreateBranch(branch *Branch) error {
	// Implement the logic to create a new branch
	panic("implement CreateBranch")
}

func (s *Service) EditBranch(branch *Branch) error {
	// Implement the logic to edit a branch
	panic("implement EditBranch")
}
