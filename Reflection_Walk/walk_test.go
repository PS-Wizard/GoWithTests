package reflectionwalk

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			Name:          "Struct with one string field",
			Input:         struct{ Name string }{"Chris"},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "Struct with two string field",
			Input: struct {
				Name string
				City string
			}{"Chris", "London"},
			ExpectedCalls: []string{"Chris", "London"},
		},

		{
			Name: "Struct with a string and an int field",
			Input: struct {
				Name string
				Age  int
			}{"Chris", 19},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "Nested fields",
			Input: Person{
				"Chris",
				Profile{33, "London"},
			},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "pointers to things",
			Input: &Person{
				"Chris",
				Profile{33, "London"},
			},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "slices",
			Input: []Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			ExpectedCalls: []string{"London", "Reykjavík"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			Walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("Got %q, want %q", got, test.ExpectedCalls)
			}
		})
	}
}
