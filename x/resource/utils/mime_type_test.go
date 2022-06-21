package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateMimeType(t *testing.T) {
	cases := []struct {
		mt    string
		valid bool
	}{
		{"application/json", true},
		{"my/mime", false},
		{"notMine/type", false},
	}

	for _, tc := range cases {
		t.Run(tc.mt, func(t *testing.T) {
			err_ := ValidateMimeType(tc.mt)

			if tc.valid {
				require.NoError(t, err_)
			} else {
				require.Error(t, err_)
			}
		})
	}
}
