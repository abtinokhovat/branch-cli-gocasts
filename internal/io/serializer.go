package io

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strings"
)

type Serializer[T any] interface {
	Serialize(data T) (string, error)
	Deserialize(jsonString string) ([]T, error)
}

type JsonSerializer[T any] struct {
	structure T
}

func NewJsonSerializer[T any](structure T) *JsonSerializer[T] {
	return &JsonSerializer[T]{
		structure: structure,
	}
}

func (s *JsonSerializer[T]) Serialize(data T) (string, error) {
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	// Parse the struct and write to the file
	if err := encoder.Encode(data); err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func (s *JsonSerializer[T]) Deserialize(jsonString string) ([]T, error) {
	// Create a JSON decoder
	reader := strings.NewReader(jsonString)
	decoder := json.NewDecoder(reader)

	// Read the array start token
	_, err := decoder.Token()
	if err != nil {
		return nil, err
	}

	// make a dataSlice for storing the values in the json file
	var dataSlice []T

	// Loop to decode each JSON object in the array
	for decoder.More() {
		// making an empty space with type of s.structPointer to store the iterated data
		var d T = reflect.New(reflect.TypeOf(s.structure).Elem()).Interface().(T)

		// the data will be stored in the d variable
		if err := decoder.Decode(d); err != nil {
			return nil, err
		}
		dataSlice = append(dataSlice, d)
	}

	// Read the array end token
	_, err = decoder.Token()
	if err != nil {
		return nil, err
	}

	return dataSlice, nil
}
