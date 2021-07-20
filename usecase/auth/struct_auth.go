package structAuth

type InputLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type InputRegister struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,gte=8"`
}
