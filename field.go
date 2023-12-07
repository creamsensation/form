package form

type Field[T any] struct {
	Type     string
	DataType string
	Id       string
	Name     string
	Label    string
	Value    T
	Errors   []error
	Required bool
	Multiple bool
}
