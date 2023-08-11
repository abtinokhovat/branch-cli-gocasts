package region

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
func (s *Service) GetAllRegions() ([]Region, error) {
	branches, err := s.adp.GetAll()
	if err != nil {
		return nil, err
	}

	return branches, nil
}
func (s *Service) GetRegionDetail(id int) (*Region, error) {
	brn, err := s.adp.GetById(id)
	if err != nil {
		return nil, err
	}

	return brn, nil
}
