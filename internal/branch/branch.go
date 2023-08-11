package branch

import "branches-cli/internal/region"

const filePath = "storage/branch.json"

type Branch struct {
	Id                int    `json:"id"`
	Name              string `json:"name,omitempty"`
	Phone             string `json:"phone,omitempty"`
	CreateDate        string `json:"createDate,omitempty"`
	NumberOfEmployees int    `json:"numberOfEmployees,omitempty"`
	RegionId          int    `json:"regionId"`
}

func NewBranch(id int, name string, phone string, createDate string, numberOfEmployees int, region region.Region) *Branch {
	return &Branch{
		Id:                id,
		Name:              name,
		Phone:             phone,
		CreateDate:        createDate,
		NumberOfEmployees: numberOfEmployees,
		RegionId:          region.Id,
	}
}
