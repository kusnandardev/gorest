package customer

type InputPort interface {
	Authenticate(authData interface{}) (interface{}, error)
}
