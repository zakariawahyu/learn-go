package services

type User struct {
	ID   int
	Name string
}

type UserServices struct {
	db []*User
}

type UserInterface interface {
	Register(user *User) string
	GetUser() []*User
	GetUserByID(i int) *User
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

func (u *UserServices) GetUserByID(i int) *User {
	for _, val := range u.db {
		if val.ID == i {
			return val
		}
	}
	return nil
}
