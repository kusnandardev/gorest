package request

type LoginRequestDto struct {
	Username string `json:"username" valid:"Required"`
	Password string `json:"password" valid:"Required"`
}
