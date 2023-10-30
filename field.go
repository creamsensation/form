package form

type Field[T any] struct {
	Type     string
	DataType string
	Id       string
	Name     string
	Value    T
	Errors   []error
	Multiple bool
}
