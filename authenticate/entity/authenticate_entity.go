package entity

type Receive_Login_Data struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Send_User_Data struct {
	UserID      int16
	UserName    string
	UserPicture int16
	UserNews    bool
}

type Auth_Token struct {
	Username string `json:"username"`
	Token    string `json:"token"`
	Expires  string `json:"expires"`
}
