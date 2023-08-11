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
	branches, err := s.adp.GetAll()
	if err != nil {
		return nil, err
	}

	var filteredBranches []Branch

	for _, brn := range branches {
		if brn.RegionId == region.Id {
			filteredBranches = append(filteredBranches, brn)
		}
	}

	return filteredBranches, nil
}

func (s *Service) GetBranchDetail(id int) (*Branch, error) {
	brn, err := s.adp.GetById(id)
	if err != nil {
		return nil, err
	}

	return brn, nil
}

func (s *Service) CreateBranch(branch *Branch) error {
	err := s.adp.Create(branch)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) EditBranch(branch *Branch) error {
	err := s.adp.UpdateById(branch)
	if err != nil {
		return err
	}

	return nil
}
