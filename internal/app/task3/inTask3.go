package task3

import "errors"

type User struct {
	name []rune
	phone int
}

//check username, changes small vowel letters to large ones and add it to the struct
//TODO: refactor
func(u *User) SetName(name string) {
	var newName = make([]rune, len(name))
	var letter rune

	for i := 0; i < len(name); i++ {
		if name[i] == 97{
			letter = 'A'
			newName[i] = letter
		} else if name[i] == 101 {
			letter = 'E'
			newName[i] = letter
		} else if name[i] == 105 {
			letter = 'I'
			newName[i] = letter
		} else if name[i] == 111 {
			letter = 'O'
			newName[i] = letter
		} else if name[i] == 117 {
			letter = 'U'
			newName[i] = letter
		} else if name[i] == 121 {
			letter = 'Y'
			newName[i] = letter
		}	else {
			letter = rune(name[i])
			newName[i] = letter
		}
	}
	u.name = newName
}
//check the user's phone number and add it to the struct
func(u *User) SetPhone(phone int) error {
	if phone < 1000000000 || phone > 999999999999999 {
		return errors.New("invalid number")
	}
	u.phone = phone
	return nil
}

//returns the username from the struct
func(u *User) Name() []rune {
	return u.name
}
//returns the user's phone number from the struct
func(u *User) Phone() int {
	return u.phone
}