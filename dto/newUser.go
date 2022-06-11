package dto

type NewUserDto struct {
	User  UserDto  `json:"user"`
	Token TokenDto `json:"token"`
}
