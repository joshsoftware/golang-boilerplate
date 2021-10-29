package service

import (
	"errors"
	"joshsoftware/golang-boilerplate/db"
	"testing"
)

func TestValidateUserAge(t *testing.T) {
	t.Run("Valid age data", func(t *testing.T) {
		testCases := []struct {
			user     db.User
			expected error
		}{
			{
				user: db.User{
					Name: "U1",
					Age:  12,
				},
				expected: nil,
			},
			{
				user: db.User{
					Name: "U2",
					Age:  99,
				},
				expected: nil,
			},
		}
		for _, ip := range testCases {
			actual := ValidateUserAge(ip.user)
			if ip.expected != actual {
				t.Fail()
			}

		}
	})
	t.Run("Invalid age data", func(t *testing.T) {
		testCases := []struct {
			user     db.User
			expected error
		}{
			{
				user: db.User{
					Name: "U3",
					Age:  0,
				},
				expected: errors.New("age cannot be less than equal to 0"),
			},
			{
				user: db.User{
					Name: "U3",
					Age:  -1,
				},
				expected: errors.New("age cannot be less than equal to 0"),
			},
		}
		for _, tc := range testCases {
			actual := ValidateUserAge(tc.user)
			if actual.Error() != tc.expected.Error() {
				t.Fail()
			}
		}
	})

}

func BenchmarkValidateUserAge(b *testing.B) {
	user := db.User{
		Name: "U1",
		Age:  12,
	}
	for i := 0; i < b.N; i++ {
		ValidateUserAge(user)
	}
}
