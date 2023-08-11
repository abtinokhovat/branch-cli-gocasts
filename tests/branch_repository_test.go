package tests

import (
	"branches-cli/internal/branch"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testData = []branch.Branch{
	{
		Id:                10,
		Name:              "Tajrish",
		Phone:             "02122299978",
		CreateDate:        "2023-10-11",
		NumberOfEmployees: 23,
		RegionId:          1,
	},
	{
		Id:                11,
		Name:              "Heravi",
		Phone:             "02122299999",
		CreateDate:        "2023-09-11",
		NumberOfEmployees: 100,
		RegionId:          1,
	},
	{
		Id:                12,
		Name:              "Gheytarie",
		Phone:             "02122299990",
		CreateDate:        "2023-09-13",
		NumberOfEmployees: 10,
		RegionId:          1,
	},
	{
		Id:                13,
		Name:              "Shoosh",
		Phone:             "02188299990",
		CreateDate:        "2023-01-13",
		NumberOfEmployees: 10,
		RegionId:          1,
	},
	{
		Id:                14,
		Name:              "Esfahan",
		Phone:             "07188299990",
		CreateDate:        "2022-01-13",
		NumberOfEmployees: 19,
		RegionId:          2,
	},
	{
		Id:                15,
		Name:              "Esfahan_Meidoon",
		Phone:             "07189299990",
		CreateDate:        "2023-05-13",
		NumberOfEmployees: 65,
		RegionId:          2,
	},
}

type MockIOHandler struct {
}

func (h *MockIOHandler) Read() ([]branch.Branch, error) {
	return testData, nil
}
func (h *MockIOHandler) WriteOne(data branch.Branch) error {
	testData = append(testData, data)
	return nil
}
func (h *MockIOHandler) DeleteAndWrite(data []branch.Branch) error {
	testData = data
	return nil
}
func (h *MockIOHandler) DeleteAll() error {
	testData = nil
	return nil
}

func TestRepository_GetById(t *testing.T) {
	for _, tc := range testData {
		t.Run(tc.Name, func(t *testing.T) {
			// 1. setup
			handler := MockIOHandler{}
			repo := branch.NewRepository(&handler)

			// 2. execution
			result, err := repo.GetById(tc.Id)
			if err != nil {
				t.Fatalf("Could not execute GetById: %s", err)
			}

			// 3. assertion
			assert.Equal(t, *result, tc, fmt.Sprintf("Expected %v, but got %v", tc, result))
		})
	}
}

func TestRepository_GetAll(t *testing.T) {
	t.Run("ordinary", func(t *testing.T) {
		// 1. setup
		handler := MockIOHandler{}
		repo := branch.NewRepository(&handler)

		// 2. execution
		result, err := repo.GetAll()
		if err != nil {
			t.Fatalf("Could not execute GetAll: %s", err)
		}

		// 3. assertion
		assert.Equal(t, testData, result, fmt.Sprintf("Expected %v, but got %v", testData, result))
	})
}

func TestRepository_Create(t *testing.T) {
	for _, tc := range testData {
		t.Run(tc.Name, func(t *testing.T) {
			// 1. setup
			handler := MockIOHandler{}
			repo := branch.NewRepository(&handler)

			// 2. execution
			err := repo.Create(&tc)
			if err != nil {
				t.Fatalf("Could not execute Create: %s", err)
			}

			// 3. assertion
			result, err := repo.GetById(tc.Id)
			if err != nil {
				t.Fatalf("Could not execute GetById: %s", err)
			}
			assert.Equal(t, tc, *result, fmt.Sprintf("Expected %v, but got %v", tc, result))
		})
	}
}

func TestRepository_UpdateById(t *testing.T) {
	for _, tc := range testData {
		t.Run(tc.Name, func(t *testing.T) {
			// 1. setup
			handler := MockIOHandler{}
			repo := branch.NewRepository(&handler)

			// 2. execution
			newData := tc
			newData.NumberOfEmployees += 5 // Modifying some data for update test
			err := repo.UpdateById(&newData)
			if err != nil {
				t.Fatalf("Could not execute UpdateById: %s", err)
			}

			// 3. assertion
			result, err := repo.GetById(tc.Id)
			if err != nil {
				t.Fatalf("Could not execute GetById: %s", err)
			}
			assert.Equal(t, newData, *result, fmt.Sprintf("Expected %v, but got %v", newData, result))
		})
	}
}
