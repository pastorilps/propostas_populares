package entity

type Receive_Login_Data struct {
	Username string `json:"username" example:"teste@gmail.com"`
	Password string `json:"password" example:"aB@123456"`
}

type Send_User_Data struct {
	UserID      int16
	UserName    string
	UserPicture int16
	UserNews    bool
}

type Auth_Token struct {
	UserID   int16
	Username string `json:"username"`
	Token    string `json:"token"`
	Expires  string `json:"expires"`
}
