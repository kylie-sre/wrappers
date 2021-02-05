package dataset

import "fmt"

type query struct{}

func (q *query) Read() (Iterator, error) {
	return &iterator{}, nil
}

func NewQuery() Query {
	return &query{}
}

type mockquery struct{}

func (q *mockquery) Read() (Iterator, error) {
	fmt.Println("Query")
	return &mockiterator{}, nil
}

func NewMockQuery() Query {
	return &mockquery{}
}

