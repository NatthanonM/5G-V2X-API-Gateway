package models

//// REST BODY ////

// AdminRegisterBody ...
type AdminRegisterBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AdminLoginBody ...
type AdminLoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
