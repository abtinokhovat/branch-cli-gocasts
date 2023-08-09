package io

import (
	"branches-cli/internal/io"
	io2 "io"
	"os"
	"testing"
)

func TestJsonIOHandler_Read(t *testing.T) {
	// testing data
	expectedData := []testStructTwo{
		{Name: "test1", Value: 42},
		{Name: "test2", Value: 24},
	}
	jsonData := `[{"name": "test1", "value": 42}, {"name": "test2", "value": 24}]`

	// temp file
	tempFile, err := createTempFile(jsonData)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	// act
	serializer := io.NewJsonSerializer[testStructTwo](testStructTwo{})
	handler := io.NewJsonIOHandler[testStructTwo](tempFile.Name(), serializer)

	data, err := handler.Read()
	if err != nil {
		t.Errorf("Error during Read: %v", err)
		return
	}

	// assert
	if len(data) != 2 {
		t.Errorf("Expected 2 items, but got %d", len(data))
	}

	// Additional assertions on the data if needed
	for i, tc := range expectedData {
		if tc != data[i] {
			t.Errorf("Expected %+v, got %+v", tc, data[i])
		}
	}
}

func TestJsonIOHandler_WriteOne(t *testing.T) {
	type testStruct struct {
		name     string
		data     interface{}
		expected string
	}

	testCases := []testStruct{
		{
			name:     "ordinary data",
			data:     testStructTwo{Name: "test1", Value: 42},
			expected: `[{"name":"test1","value":42}]`,
		},
		{
			name:     "empty data",
			data:     testStructTwo{},
			expected: `[{"name":"","value":0}]`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tempFile, err := createTempFile("")
			if err != nil {
				t.Fatal(err)
			}
			defer tempFile.Close()
			defer os.Remove(tempFile.Name())

			// act
			serializer := io.NewJsonSerializer[testStructTwo](testStructTwo{})
			handler := io.NewJsonIOHandler[testStructTwo](tempFile.Name(), serializer)

			err = handler.WriteOne(tc.data.(testStructTwo))
			if err != nil {
				t.Errorf("%v", err)
				return
			}

			// assert
			content, err := io2.ReadAll(tempFile)
			if err != nil {
				t.Errorf("%v", err)
				return
			}

			if string(content) != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, content)
				return
			}
		})
	}
}

func createTempFile(data string) (*os.File, error) {
	file, err := os.CreateTemp("", "test*.json")
	if err != nil {
		return nil, err
	}
	// TODO
	//defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return nil, err
	}

	return file, nil
}
