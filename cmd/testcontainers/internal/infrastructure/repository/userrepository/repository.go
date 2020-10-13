package userrepository

func New() IUserRepository {
	return UserRepository{}
}