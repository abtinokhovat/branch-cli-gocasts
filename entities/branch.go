package pkg

import (
	"branches-cli/internal/io"
)

const filePath = "storage/branch.json"

type Branch struct {
	Id                int    `json:"Id,omitempty"`
	Name              string `json:"Name,omitempty"`
	Phone             string `json:"Phone,omitempty"`
	CreateDate        string `json:"CreateDate,omitempty"`
	NumberOfEmployees int    `json:"NumberOfEmployees,omitempty"`
	manipulator       *io.JsonIOHandler[Branch]
}

func NewBranch(id int, name string, phone string, createDate string, numberOfEmployees int) *Branch {

	serializer := io.NewJsonSerializer[Branch](Branch{})
	manipulator := io.NewJsonIOHandler[Branch](filePath, serializer)

	return &Branch{
		Id:                id,
		Name:              name,
		Phone:             phone,
		CreateDate:        createDate,
		NumberOfEmployees: numberOfEmployees,
		manipulator:       manipulator,
	}
}

//func GetBranches() []*Branch {
//	handler := data.NewJsonIOHandler(filePath, &Branch{})
//	branches, err := handler.Read()
//	if err != nil {
//		fmt.Errorf(err.Error())
//	}
//	return branches
//}

//func (b *Branch) Create() *Branch {
//	// data Helper
//	//b.manipulator = &config
//
//	// set id to max id + 1
//	id, _ := b.manipulator.MaxId()
//	b.Id = id + 1
//
//	slice := []interface{}{b}
//	_ = b.manipulator.Write(slice)
//
//	return b
//}
//
//func (b *Branch) Edit(name string, phone string, date string, numOfEmp int) *Branch {
//	if name != "" {
//		b.Name = name
//	}
//	if phone != "" {
//		b.Phone = phone
//	}
//	if date != "" {
//		b.CreateDate = date
//	}
//	if numOfEmp != b.NumberOfEmployees {
//		b.NumberOfEmployees = numOfEmp
//	}
//	_ = b.manipulator.UpdateById(b.Id, &b)
//	return b
//}
