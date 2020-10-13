package userservice

func New() IUserService {
	return UserService{}
}