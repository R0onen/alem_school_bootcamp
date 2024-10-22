package firststruct

func NewUser(name, password string) User {
	u := User{
		Name:     name,
		password: password,
	}
	return u
}
