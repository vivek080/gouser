package userdemo

type ProfileManager interface {
	CreateProfile(map[string]interface{}) error
	UpdateProfile(map[string]interface{}) error
}

type User struct {
	P ProfileManager
}

func NewUserService(profileManager ProfileManager) *User {
	u := User{P: profileManager}
	return &u
}

func (u *User) CreateUser(fName string, lName string) string {
	value := "Hello " + fName + " " + lName + ", this function implements code reusability"
	return value
}
