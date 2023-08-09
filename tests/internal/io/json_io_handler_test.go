package io

import (
	"branches-cli/internal/io"
	"fmt"
	io2 "io"
	"os"
	"testing"
)

type testStruct struct {
	Name  string
	Age   int
	Likes []int
}

var (
	testData = []testStruct{
		{
			Name:  "David",
			Age:   40,
			Likes: []int{15, 23, 35},
		},
		{
			Name:  "Eve",
			Age:   22,
			Likes: []int{12, 28, 44, 72},
		},
		{
			Name:  "Frank",
			Age:   31,
			Likes: []int{7, 19},
		},
	}
	testDataString = `[{"Name":"David","Age":40,"Likes":[15,23,35]},{"Name":"Eve","Age":22,"Likes":[12,28,44,72]},{"Name":"Frank","Age":31,"Likes":[7,19]}]`
)

type MockSerializer struct {
	io.Serializer[testStruct]
}

func (s MockSerializer) Serialize(data testStruct) (string, error) {
	return fmt.Sprintf("{ Name: %s, Age: %d, Likes: %v }", data.Name, data.Age, data.Likes), nil
}
func (s MockSerializer) Deserialize(jsonString string) ([]testStruct, error) {
	var testData []testStruct

	data := testStruct{
		Name:  jsonString,
		Age:   404,
		Likes: nil,
	}
	testData = append(testData, data)

	return testData, nil
}

func createTempFile(data string) (*os.File, error) {
	file, err := os.CreateTemp("", "test*.json")
	if err != nil {
		return nil, err
	}

	_, err = file.WriteString(data)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func TestJsonIOHandler_Read(t *testing.T) {
	testCases := []struct {
		name string
		data string
	}{
		{name: "ordinary", data: "salam"},
		{name: "empty", data: ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			file, err := createTempFile(tc.data)
			if err != nil {
				return
			}
			defer file.Close()
			defer os.Remove(file.Name())

			serializer := MockSerializer{}
			handler := io.NewJsonIOHandler[testStruct](file.Name(), serializer)

			read, err := handler.Read()
			if err != nil {
				t.Errorf("%v", err)
			}
			if len(read) == 0 {
				t.Errorf("List should not be empty an error may occured")
			}
			if read[0].Name != tc.data {
				t.Errorf("Expected %s, but got %s\nThe whole data:%v", read[0].Name, tc.data, read)
			}
		})
	}

}

func TestJsonIOHandler_DeleteAll(t *testing.T) {
	t.Run("ordinary", func(t *testing.T) {
		// 1. setup
		dummyText := "The data that will be deleted"
		file, err := createTempFile(dummyText)
		if err != nil {
			return
		}
		defer file.Close()
		defer os.Remove(file.Name())

		beforeDelete, err := io2.ReadAll(file)
		if err != nil || string(beforeDelete) == dummyText {
			t.Fatalf("error in the test failed to writing to file\nReadErr: %s,\n data: %v", err, beforeDelete)
		}

		serializer := MockSerializer{}
		handler := io.NewJsonIOHandler[testStruct](file.Name(), serializer)

		// 2. execution
		err = handler.DeleteAll()
		if err != nil {
			t.Fatalf("Error on DeleteAll: %s", err)
		}

		// 3. assertion
		afterDelete, err := io2.ReadAll(file)
		if err != nil || len(afterDelete) != 0 {
			t.Fatalf("Falied to delete the content\nDeleteErr: %s,\n data: %v", err, afterDelete)
		}
	})

	t.Run("not_valid_file", func(t *testing.T) {
		// 1. setup
		serializer := MockSerializer{}
		handler := io.NewJsonIOHandler[testStruct]("non_existent_file.json", serializer)

		// 2. execution
		err := handler.DeleteAll()

		// 3. assertion
		if err == nil {
			t.Errorf("Expected an error, but got none")
		}
	})
}

func TestJsonIOHandler_DeleteAndWrite(t *testing.T) {
	t.Run("ordinary", func(t *testing.T) {
		// 1. setup
		initText := "The init data is here and here"

		file, err := createTempFile(initText)
		if err != nil {
			return
		}
		defer file.Close()
		defer os.Remove(file.Name())

		beforeDeleteAndWrite, err := io2.ReadAll(file)
		if err != nil || string(beforeDeleteAndWrite) == initText {
			t.Fatalf("error in the test failed to writing to file\nReadErr: %s,\n data: %v", err, beforeDeleteAndWrite)
		}

		serializer := MockSerializer{}
		handler := io.NewJsonIOHandler[testStruct](file.Name(), serializer)

		// 2. execution
		err = handler.DeleteAndWrite(testData)
		if err != nil {
			t.Fatalf("Error on DeleteAndWrite: %s", err)
		}

		// 3. assertion
		_, _ = file.Seek(0, 0)
		afterDelete, err := io2.ReadAll(file)
		if string(afterDelete) != testDataString {
			t.Fatalf("Falied to delete and rewrite the content\nDeleteAndWrite: %s,\n data: %v", err, afterDelete)
		}
	})

	t.Run("not_valid_file", func(t *testing.T) {
		// 1. setup
		serializer := MockSerializer{}
		handler := io.NewJsonIOHandler[testStruct]("non_existent_file.json", serializer)

		// 2. execution
		err := handler.DeleteAndWrite(testData)

		// 3. assertion
		if err == nil {
			t.Errorf("Expected an error, but got none")
		}
	})
}
