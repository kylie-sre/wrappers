package dataset

type Iterator interface {
	Next(interface{}) error
}

type Query interface {
	Read() (Iterator, error)
}

type Dataset interface {
	Query(string) Query
}
