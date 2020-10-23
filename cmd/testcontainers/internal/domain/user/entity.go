package user

func (user *Entity) TableName() string {
	return "users"
}

func (user Entity) ChangeEmail(email string) {
	user.Email = email
}
