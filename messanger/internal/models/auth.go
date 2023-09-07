package models

type Credentials struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TokenPair struct {
	JWT          string
	RefreshToken string
}
