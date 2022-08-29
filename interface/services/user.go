package services

type User struct {
	Name string
}

type UserServices struct {
	db []*User
}

type UserInterface interface {
	Register(user *User) string
	GetUser() []*User
}

func NewUserServices(user []*User) UserServices {
	return UserServices{
		db: user,
	}
}

func (u *UserServices) Register(user *User) string {
	u.db = append(u.db, user)
	return user.Name + " Berhasil dibuat"
}

func (u *UserServices) GetUser() []*User {
	return u.db
}
