package user

type UserResponseDto struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func UserDtoFromModel(user *User) UserResponseDto {
	return UserResponseDto{Email: user.Email, Username: user.Username, LastName: user.LastName, FirstName: user.FirstName}
}
