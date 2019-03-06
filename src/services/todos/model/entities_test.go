package model

import (
	"services/utils/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestName(t *testing.T) {

	valid := rand.String(100)

	tcs := []struct {
		Description    string
		GivenRaw       string
		ExpectedResult Name
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
		{
			Description:    "valid",
			GivenRaw:       valid,
			ExpectedResult: Name(valid),
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Description, func(t *testing.T) {
			n, err := NewName(tc.GivenRaw)

			if tc.ExpectedError != nil {
				require.EqualError(t, err, tc.ExpectedError.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.ExpectedResult, n)
			}
		})
	}
}

func TestDescription(t *testing.T) {
	valid := rand.String(300)

	tcs := []struct {
		Description    string
		GivenRaw       string
		ExpectedResult Description
		ExpectedError  error
	}{
		{
			Description:   "invalid/too long",
			GivenRaw:      rand.String(301),
			ExpectedError: ErrDescriptionTooLong,
		},
		{
			Description:    "valid",
			GivenRaw:       valid,
			ExpectedResult: Description(valid),
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Description, func(t *testing.T) {
			d, err := NewDescription(tc.GivenRaw)

			if tc.ExpectedError != nil {
				require.EqualError(t, err, tc.ExpectedError.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.ExpectedResult, d)
			}
		})
	}
}
