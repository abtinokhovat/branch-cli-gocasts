package branch

const filePath = "storage/branch.json"

type Branch struct {
	Id                int    `json:"id,omitempty"`
	Name              string `json:"name,omitempty"`
	Phone             string `json:"phone,omitempty"`
	CreateDate        string `json:"createDate,omitempty"`
	NumberOfEmployees int    `json:"numberOfEmployees,omitempty"`
}

func NewBranch(id int, name string, phone string, createDate string, numberOfEmployees int) *Branch {
	return &Branch{
		Id:                id,
		Name:              name,
		Phone:             phone,
		CreateDate:        createDate,
		NumberOfEmployees: numberOfEmployees,
	}
}
