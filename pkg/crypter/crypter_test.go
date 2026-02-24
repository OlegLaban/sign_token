package crypter

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncryptAndDecrypt(t *testing.T) {
	type Test struct {
		Name          string
		Input         string
		CustomDecrypt string
		ActualErr     error
		DecryptErr    error
	}
	var key = "some-vsecret-key"

	tests := []Test{
		{
			Name:  "Success",
			Input: "some-insensitive-info",
		},
		{
			Name:          "Decrypt error",
			Input:         "some-insensitive-info",
			CustomDecrypt: "invalid",
			DecryptErr:    ErrTextToShort,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			crypt := New(key)
			res, err := crypt.Encrypt(tt.Input)
			if tt.ActualErr != nil {
				require.ErrorIs(t, err, tt.ActualErr)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, res)
				if tt.CustomDecrypt != "" {
					res = []byte(tt.CustomDecrypt)
				}
				res, err := crypt.Decrypt(res)
				if tt.DecryptErr != nil {
					require.ErrorIs(t, err, tt.DecryptErr)
				} else {
					require.Equal(t, res, []byte(tt.Input))
				}
			}
		})
	}
}
