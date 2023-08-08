package io

import (
	"bytes"
	"encoding/json"
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
	var objects []T
	if len(jsonString) == 0 {
		return objects, nil
	}
	err := json.Unmarshal([]byte(jsonString), &objects)
	if err != nil {
		return nil, err
	}
	return objects, nil
}
