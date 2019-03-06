package model

import (
	"services/utils/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTodo(t *testing.T) {
	validName := rand.String(100)
	validDescription := rand.String(300)
	createdAt := time.Now().UTC()
	until := createdAt.Add(time.Hour)

	tcs := []struct {
		Description      string
		GivenName        string
		GivenDescription string
		ExpectedResult   *Todo
		ExpectedError    error
	}{
		{
			Description:      "invalid/empty name",
			GivenName:        "",
			GivenDescription: validDescription,
			ExpectedError:    ErrEmptyName,
		},
		{
			Description:      "invalid/name too long",
			GivenName:        rand.String(101),
			GivenDescription: validDescription,
			ExpectedError:    ErrNameTooLong,
		},
		{
			Description:      "invalid/description too long",
			GivenName:        validName,
			GivenDescription: rand.String(301),
			ExpectedError:    ErrDescriptionTooLong,
		},
		{
			Description:      "valid",
			GivenName:        validName,
			GivenDescription: validDescription,
			ExpectedResult: &Todo{
				Name:        Name(validName),
				Description: Description(validDescription),
				CreatedAt:   createdAt,
				Until:       until,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Description, func(t *testing.T) {
			n, err := NewTodo(tc.GivenName, tc.GivenDescription, createdAt, until)

			if tc.ExpectedError != nil {
				require.EqualError(t, err, tc.ExpectedError.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.ExpectedResult, n)
			}
		})
	}
}
