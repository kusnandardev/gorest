package customer

type Interactor struct {
}

func NewCustomerInteractor() *Interactor {
	return &Interactor{}
}

func (i *Interactor) Authenticate(authData interface{}) (interface{}, error) {

	return nil, nil
}
