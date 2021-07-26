package task3

import (
	"errors"
	"testing"
)

type names struct{
	nameTest string
	name string
	expected []rune
}

func TestSetName(t *testing.T) {
	tests := []names {
		{
			nameTest: "changing the case of the letters E and A",
			name: "Elena",
			expected: []rune{'E', 'l', 'E', 'n', 'A'},
		},
		{
			nameTest: "changing the case of the letters I, O, A and Y",
			name: "nikolay",
			expected: []rune{'n', 'I', 'k', 'O', 'l', 'A', 'Y'},
		},
		{
			nameTest: "changing the case of the letters Y, U and A",
			name: "yura",
			expected: []rune{'Y', 'U', 'r', 'A'},
		},
		{
			nameTest: "empty name",
			name: "",
			expected: []rune{},
		},
	}
	result := User{}
	for _, value := range tests {
		result.SetName(value.name)

		for index, letter := range value.expected {
			if letter != result.Name()[index] {
				t.Log("BAD TEST", value.nameTest)
			}
			if len(value.expected) != len(result.Name()) {
				t.Log("BAD TEST", value.nameTest)
			}
		}
	}
}

type numbers struct {
	nameTest string
	phone int
	failError error
	mustFail bool
}
func TestSetPhone(t *testing.T) {
	tests := []numbers {
		{
			nameTest: "normal number",
			phone: 79803086119,
			failError: nil,
			mustFail: false,
		},
		{
			nameTest: "number is shorter than ten digits",
			phone: 1234,
			failError: errors.New("invalid number"),
			mustFail: true,
		},
		{
			nameTest: "number is zero",
			phone: 0,
			failError: errors.New("invalid number"),
			mustFail: true,
		},
		{
			nameTest: "number is longer than 15 digits",
			phone: 9999999999999999,
			failError: errors.New("invalid number"),
			mustFail: true,
		},
		{
			nameTest: "normal number two",
			phone: 79803086115,
			failError: nil,
			mustFail: false,
		},
	}
	result := User{}
	for _, value := range tests {
		result.SetPhone(value.phone)

		if value.mustFail == false && value.failError == nil {
			t.Log("OK", value.nameTest)
		} else if value.mustFail == true && value.failError == errors.New("invalid number") {
			t.Log("OK", value.nameTest)
		} else if value.mustFail == true && value.failError == nil {
			t.Log("BAD TEST", value.nameTest)
		} else if value.mustFail == false && value.failError == errors.New("invalid number") {
			t.Log("BAD TEST", value.nameTest)
		}
	}
}

