package data

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type IDataManipulator interface {
	Read() ([]interface{}, error)
	Write(data []interface{}) error
	MaxId() (int, error)
	UpdateById(id int, data interface{}) error
}

// JSONDataManipulator is a data manipulator that reads and writes data to a JSON file.
type JSONDataManipulator struct {
	IDataManipulator
	FilePath  string
	StructPtr interface{}
}

// Read reads data from the JSON file and populates the provided struct slice.
func (j *JSONDataManipulator) Read() ([]interface{}, error) {
	file, err := os.Open(j.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	results := make([]interface{}, 0)

	if len(data) > 0 {
		err = json.Unmarshal(data, &results)
		if err != nil {
			return nil, err
		}
	}

	return results, nil
}

// Write writes data from the provided struct slice to the JSON file.
func (j *JSONDataManipulator) Write(data []interface{}) error {
	file, err := os.Create(j.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encodedData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(encodedData)
	return err
}

// MaxId finds the maximum ID in the JSON file.
func (j *JSONDataManipulator) MaxId() (int, error) {
	data, err := j.Read()
	if err != nil {
		return 0, err
	}

	maxID := 0
	for _, item := range data {
		idField := item.(map[string]interface{})["Id"]
		if id, ok := idField.(float64); ok {
			if int(id) > maxID {
				maxID = int(id)
			}
		}
	}

	return maxID, nil
}

// UpdateById updates data in the JSON file for the given ID.
func (j *JSONDataManipulator) UpdateById(id int, newData interface{}) error {
	data, err := j.Read()
	if err != nil {
		return err
	}

	idFieldName := "Id" // Modify this field name based on your struct definition.

	// Find the struct instance to update based on the ID field.
	for i, item := range data {
		itemData := item.(map[string]interface{})
		if idField, ok := itemData[idFieldName]; ok {
			if idVal, ok := idField.(float64); ok {
				if int(idVal) == id {
					// Found the matching item, update its data with newData.
					data[i] = newData
					break
				}
			}
		}
	}

	// Write the updated data back to the JSON file.
	return j.Write(data)
}
