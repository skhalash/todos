package model

import (
	"services/utils/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestName(t *testing.T) {
	tcs := []struct {
		Description    string
		GivenRaw       string
		ExpectedResult *Name
		ExpectedError  error
	}{
		{
			Description:   "invalid/empty",
			GivenRaw:      "",
			ExpectedError: ErrEmptyName,
		},
		{
			Description:   "invalid/too long",
			GivenRaw:      rand.String(101),
			ExpectedError: ErrNameTooLong,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Description, func(t *testing.T) {
			name, err := NewName(tc.GivenRaw)

			if tc.ExpectedError != nil {
				require.EqualError(t, err, tc.ExpectedError.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.ExpectedResult, name)
			}
		})
	}
}
