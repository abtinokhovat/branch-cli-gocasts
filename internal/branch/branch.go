package branch

import "fmt"

const StoragePath = "storage/branch.json"

type Branch struct {
	Id                int    `json:"id"`
	Name              string `json:"name,omitempty"`
	Phone             string `json:"phone,omitempty"`
	CreateDate        string `json:"createDate,omitempty"`
	NumberOfEmployees int    `json:"numberOfEmployees,omitempty"`
	RegionId          int    `json:"regionId"`
}

func New(id int, name string, phone string, createDate string, numberOfEmployees int, regionId int) *Branch {
	return &Branch{
		Id:                id,
		Name:              name,
		Phone:             phone,
		CreateDate:        createDate,
		NumberOfEmployees: numberOfEmployees,
		RegionId:          regionId,
	}
}

func (b *Branch) String() string {
	result := fmt.Sprintf("\033[33m#️⃣%d-%s\033[0m\n---------------\n☎️:%s\n🚻:%d\n🗻:%d\n📅:%s",
		b.Id, b.Name,
		b.Phone,
		b.NumberOfEmployees,
		b.RegionId,
		b.CreateDate)
	return result
}
