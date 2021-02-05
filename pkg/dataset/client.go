package dataset

import "fmt"

type client struct{}

func (c *client) Query(q string) Query {
	return &query{}
}

func New() Dataset {
	return &client{}
}

type mockclient struct{}

func (c *mockclient) Query(q string) Query {
	fmt.Println(q)
	return &query{}
}

func NewMock() Dataset {
	return &mockclient{}
}
