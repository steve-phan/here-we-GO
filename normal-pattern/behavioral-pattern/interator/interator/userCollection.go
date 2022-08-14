package interator

type UserCollection struct {
	Users []*User
}

func (u *UserCollection) CreateInterator() Iterator {
	return &UserIterator{
		Users: u.Users,
	}
}
