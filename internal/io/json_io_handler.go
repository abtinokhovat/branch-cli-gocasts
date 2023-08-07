package io

import (
	"encoding/json"
	"io"
	"os"
)

type JsonReader[T any] interface {
	Read() ([]T, error)
}

type JsonWriter[T any] interface {
	WriteOne(data T) error
	WriteMany(data []T) error
}

type JsonIOHandler[T any] struct {
	FilePath   string
	serializer Serializer[T]
}

func NewJsonIOHandler[T any](path string, serializer Serializer[T]) *JsonIOHandler[T] {
	return &JsonIOHandler[T]{
		FilePath:   path,
		serializer: serializer,
	}
}

func (h *JsonIOHandler[T]) openFile() (*os.File, error) {
	file, err := os.Open(h.FilePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (h *JsonIOHandler[T]) Read() ([]T, error) {
	file, err := h.openFile()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	deserialized, err := h.serializer.Deserialize(string(content))
	if err != nil {
		return nil, err
	}

	return deserialized, nil
}

func (h *JsonIOHandler[T]) WriteOne(data T, encoder *json.Encoder) error {
	if err := encoder.Encode(data); err != nil {
		return err
	}
	return nil
}

func (h *JsonIOHandler[T]) WriteMany(data []T, encoder *json.Encoder) error {
	for _, item := range data {
		if err := h.WriteOne(item, encoder); err != nil {
			return err
		}
	}
	return nil
}

func (h *JsonIOHandler[T]) WriteOneToFile(data T, file *os.File) error {
	encoder := json.NewEncoder(file)
	return h.WriteOne(data, encoder)
}

func (h *JsonIOHandler[T]) WriteManyToFile(data []T, file *os.File) error {
	encoder := json.NewEncoder(file)
	return h.WriteMany(data, encoder)
}
