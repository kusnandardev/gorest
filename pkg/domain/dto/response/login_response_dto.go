package response

type LoginResponseDto struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Token    string `json:"token"`
}
