package interfaces

type Validator interface {
	Struct(s interface{}) error
}
