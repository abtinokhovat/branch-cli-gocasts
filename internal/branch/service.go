package branch

import "branches-cli/internal/region"

type Service struct {
	adp StorageAdapter
}

func BuildService() *Service {
	repo := BuildRepository()
	return NewBranchService(repo)
}
func NewBranchService(adp StorageAdapter) *Service {
	return &Service{
		adp: adp,
	}
}

func (s *Service) NewId() (int, error) {
	branches, err := s.adp.GetAll()
	if err != nil {
		return -1, err
	}

	var max int
	// get max id
	for _, brn := range branches {
		if brn.Id > max {
			max = brn.Id
		}
	}

	return max + 1, err
}
func (s *Service) GetAllBranches() ([]Branch, error) {
	branches, err := s.adp.GetAll()
	if err != nil {
		return nil, err
	}

	return branches, nil
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
