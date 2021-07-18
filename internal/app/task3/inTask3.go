package task3

import "errors"

type User struct {
	name string
	phone int
}

func(u *User) SetName(name string) {
	u.name = name
}
func(u *User) SetPhone(phone int) error {
	if phone < 1000000000 || phone > 999999999999999 {
		return errors.New("invalid number")
	}
	u.phone = phone
	return nil
}

func(u *User) Name() string {
	return u.name
}
func(u *User) Phone() int {
	return u.phone
}