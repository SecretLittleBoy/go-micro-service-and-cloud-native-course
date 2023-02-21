// Code generated by goctl. DO NOT EDIT.
package types

type SignupRequest struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
	Gender     int    `json:"gender,options=0|1|2,default=0"`
}

type SignupResponse struct {
	Message string `json:"message"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message      string `json:"message"`
	AccessToken  string `json:"accessToken"`
	AccessExpire int    `json:"accessExpire"`
	RefreshAfter int    `json:"refreshAfter"`
}

type DetailRequest struct {
	UserID int64 `form:"userID"`
}

type DetailResponse struct {
	Username string `json:"userName"`
	Gender   int    `json:"gender"`
}
