package request

var SignUpBody struct {
	Name     string
	Email    string
	Password string
}

var SignInBody struct {
	Email    string
	Password string
}
