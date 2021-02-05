package dataset

import (
	"fmt"
)

type iterator struct{}

func (i *iterator) Next(dst interface{}) error {
	return nil
}

func NewIterator() Iterator {
	return &iterator{}
}

type mockiterator struct{}

func (i *mockiterator) Next(dst interface{}) error {
	fmt.Println(dst)
	return nil
}

func NewMockIterator() Iterator {
	return &mockiterator{}
}
