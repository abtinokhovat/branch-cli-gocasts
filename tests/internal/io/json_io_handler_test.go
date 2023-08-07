package io

import (
	"branches-cli/internal/io"
	"encoding/json"
	"os"
)

type MockJsonSerializer[T any] struct {
	io.Serializer[T]
}

func (s *MockJsonSerializer[T]) Serialize(data T) (string, error) {
	// Serialize the `data` value to a `string`.
	s2, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// Return the serialized `string` and a nil `error`.
	return string(s2), nil
}

func (s *MockJsonSerializer[T]) Deserialize(s2 string) ([]T, error) {
	//TODO implement me
	panic("implement me")
}

//func TestJsonIOHandler_Read(t *testing.T) {
//	jsonData := `[{"name": "test1", "value": 42}, {"name": "test2", "value": 24}]`
//	tempFile, err := createTempFile(jsonData)
//	if err != nil {
//		t.Fatal(err)
//	}
//	defer os.Remove(tempFile.Name())
//
//	serializer := MockJsonSerializer[TestData]{}
//	handler := io.NewJsonIOHandler[TestData](tempFile.Name(), serializer)
//
//	data, err := handler.Read()
//	if err != nil {
//		t.Errorf("Error during Read: %v", err)
//		return
//	}
//
//	if len(data) != 2 {
//		t.Errorf("Expected 2 items, but got %d", len(data))
//	}
//
//	// Additional assertions on the data if needed
//}

func createTempFile(data string) (*os.File, error) {
	file, err := os.CreateTemp("", "test")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return nil, err
	}

	return file, nil
}
