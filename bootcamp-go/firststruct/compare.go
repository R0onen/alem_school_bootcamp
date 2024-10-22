package firststruct

type User struct {
	Name     string
	password string
}

func (u User) Compare(v User) bool {
	identical := false
	if u.Name == v.Name {
		identical = true
	} else {
		return false
	}
	if u.password == v.password {
		identical = true
	} else {
		return false
	}
	return identical
}
