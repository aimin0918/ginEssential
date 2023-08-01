package dto

import "oceanlearn.teach/ginessential/model"

type UserDto struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Id:        int(user.ID),
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
