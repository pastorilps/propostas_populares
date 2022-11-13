package entity

type Users struct {
	ID         int16  `json:"id" example:"1"`
	Name       string `json:"name" example:"User Name"`
	Email      string `json:"email" example:"test@test.com"`
	Password   string `json:"password" example:"b3f8b6283fce62d85c5b6334c8ee9a611aed144c3d93d11ef2759f6baabdc3b0"`
	Picture    int16  `json:"picture" example:"1"`
	Newsletter bool   `json:"newsletter" example:"true"`
}

type Send_User struct {
	ID         int16
	Name       string `json:"name" validate:"required" example:"User Name"`
	Email      string `json:"email" validate:"required" example:"test@test.com"`
	Password   string `json:"password" validate:"required" example:"aB@123456"`
	Picture    int16  `json:"picture" validate:"required" example:"1"`
	Newsletter bool   `json:"newsletter" validate:"required" example:"true"`
}

type Receive_User struct {
	ID         int16
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Picture    int16  `json:"picture"`
	Newsletter bool   `json:"newsletter"`
	Token      string `json:"token"`
	UserID     int16  `json:"userid"`
}

type Delete_User struct {
	ID int16 `json:"id" validate:"required" example:"1"`
}
