package model

type UserLoginResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
	UserId     int    `json:"user_id,omitempty"`
	Token      string `json:"token,omitempty"`
}
