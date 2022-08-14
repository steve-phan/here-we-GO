package interator

type UserIterator struct {
	index int
	Users []*User
}

func (u *UserIterator) HasNext() bool {
	return u.index < len(u.Users)
}

func (u *UserIterator) GetNext() *User {
	if u.HasNext() {
		User := u.Users[u.index]
		u.index++
		return User
	}
	return nil
}
