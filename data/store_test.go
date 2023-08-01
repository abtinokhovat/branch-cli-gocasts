package data

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

func TestJSONDataManipulator_ReadAndWrite(t *testing.T) {
	// Create test data
	testData := []interface{}{
		map[string]interface{}{"Id": 1, "Name": "Item 1"},
		map[string]interface{}{"Id": 2, "Name": "Item 2"},
	}

	// Create a temporary test file
	file, err := ioutil.TempFile("", "test_data.json")
	if err != nil {
		t.Fatalf("Failed to create temporary test file: %v", err)
	}
	defer os.Remove(file.Name())

	// Create JSONDataManipulator
	jsonManipulator := &JSONDataManipulator{
		FilePath:  file.Name(),
		StructPtr: &testData,
	}

	// Write test data to the JSON file
	err = jsonManipulator.Write(testData)
	if err != nil {
		t.Fatalf("Failed to write data to JSON file: %v", err)
	}

	// Read data from the JSON file
	readData, err := jsonManipulator.Read()
	if err != nil {
		t.Fatalf("Failed to read data from JSON file: %v", err)
	}

	// Compare the original data with the read data
	if len(testData) != len(readData) {
		t.Fatalf("Read data length mismatch. Expected: %d, Got: %d", len(testData), len(readData))
	}

	for i := range testData {
		original, _ := json.Marshal(testData[i])
		read, _ := json.Marshal(readData[i])
		if string(original) != string(read) {
			t.Errorf("Data mismatch for entry %d. Expected: %s, Got: %s", i+1, original, read)
		}
	}
}

func TestJSONDataManipulator_MaxId(t *testing.T) {
	// Create test data
	testData := []interface{}{
		map[string]interface{}{"Id": 1, "Name": "Item 1"},
		map[string]interface{}{"Id": 2, "Name": "Item 2"},
	}

	// Create a temporary test file
	file, err := ioutil.TempFile("", "test_data.json")
	if err != nil {
		t.Fatalf("Failed to create temporary test file: %v", err)
	}
	defer os.Remove(file.Name())

	// Create JSONDataManipulator
	jsonManipulator := &JSONDataManipulator{
		FilePath:  file.Name(),
		StructPtr: &testData,
	}

	// Write test data to the JSON file
	err = jsonManipulator.Write(testData)
	if err != nil {
		t.Fatalf("Failed to write data to JSON file: %v", err)
	}

	// Get the maximum ID
	maxID, err := jsonManipulator.MaxId()
	if err != nil {
		t.Fatalf("Failed to get maximum ID: %v", err)
	}

	// Verify the maximum ID
	expectedMaxID := 2
	if maxID != expectedMaxID {
		t.Errorf("Maximum ID mismatch. Expected: %d, Got: %d", expectedMaxID, maxID)
	}
}

func TestJSONDataManipulator_UpdateById(t *testing.T) {
	// Create test data
	testData := []interface{}{
		map[string]interface{}{"Id": 1, "Name": "Item 1"},
		map[string]interface{}{"Id": 2, "Name": "Item 2"},
	}

	// Create a temporary test file
	file, err := ioutil.TempFile("", "test_data.json")
	if err != nil {
		t.Fatalf("Failed to create temporary test file: %v", err)
	}
	defer os.Remove(file.Name())

	// Create JSONDataManipulator
	jsonManipulator := &JSONDataManipulator{
		FilePath:  file.Name(),
		StructPtr: &testData,
	}

	// Write test data to the JSON file
	err = jsonManipulator.Write(testData)
	if err != nil {
		t.Fatalf("Failed to write data to JSON file: %v", err)
	}

	// Update data by ID
	updateData := map[string]interface{}{"Id": 2, "Name": "Updated Item 2"}
	err = jsonManipulator.UpdateById(2, updateData)
	if err != nil {
		t.Fatalf("Failed to update data by ID: %v", err)
	}

	// Read data from the JSON file
	readData, err := jsonManipulator.Read()
	if err != nil {
		t.Fatalf("Failed to read data from JSON file: %v", err)
	}

	// Find the updated item
	var updatedItem interface{}
	for _, item := range readData {
		itemData := item.(map[string]interface{})
		if idVal, ok := itemData["Id"]; ok {
			if int(idVal.(float64)) == 2 {
				updatedItem = item
				break
			}
		}
	}

	if updatedItem == nil {
		t.Fatal("Updated item not found")
	}

	// Verify the updated item's data
	expectedUpdatedData := `{"Id":2,"Name":"Updated Item 2"}`
	updatedData, _ := json.Marshal(updatedItem)
	if string(updatedData) != expectedUpdatedData {
		t.Errorf("Updated data mismatch. Expected: %s, Got: %s", expectedUpdatedData, updatedData)
	}
}
