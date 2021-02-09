package models

// Admin ...
type Admin struct {
	Username string `json:"username"`
}

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

//// REST RESPONSE ////

// WebAuthProfileResponse ...
type WebAuthProfileResponse struct {
	BaseResponse
	Data *Admin `json:"data"`
}
