package task3

import "testing"

type names struct{
	name string
	expected []rune
}

func TestSetName(t *testing.T) {
	tests := []names {
		{
			name: "Elena",
			expected: []rune{'E', 'l', 'E', 'n', 'A'},
		},
		{
			name: "nikolay",
			expected: []rune{'n', 'I', 'k', 'O', 'l', 'A', 'Y'},
		},
		{
			name: "yura",
			expected: []rune{'Y', 'U', 'r', 'A'},
		},
		{
			name: "",
			expected: []rune{},
		},
	}
	result := User{}
	for index, value := range tests {
		count := index + 1
		result.SetName(value.name)

		for i, v := range value.expected {
			if v != result.Name()[i] {
				t.Errorf("Test %v failed", count)
				t.Errorf("Index %v Value %q", index, value)
			}
			if len(value.expected) != len(result.Name()) {
				t.Errorf("Expected lenght %v Result %v", len(value.expected), len(result.Name()))
			}
		}
	}
}

type numbers struct {
	phone int
	expected int
	exp error
}
func TestSetPhone(t *testing.T) {
	tests := []numbers {
		{
			phone: 79803086119,
			expected: 79803086119,
		},
		{
			phone: 1234,
			exp:
		},
		{
			phone: 0,
			expected: 0,
		},
		{
			phone: 9999999999999999,
			expected: 0,
		},
	}
	result := User{}
	for index, value := range tests {
		count := index + 1
		result.SetPhone(value.phone)

		if value.expected != result.Phone() {
			t.Errorf("Test %v failed", count)
			t.Errorf("Expected %v Result %v", value.expected, result.Phone())
		}
	}
}

