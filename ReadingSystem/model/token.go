package model

type Tokens struct {
	ID           uint   `json:"ID"`
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}
