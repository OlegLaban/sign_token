package keygenerator

import (
	"errors"
	"testing"

	"github.com/OlegLaban/sing_token/internal/domain"
	"github.com/OlegLaban/sing_token/internal/pkg/key_generator/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	type Test struct {
		Name        string
		Input       domain.Payload
		EncryptCall func(string) (string, error)
		DebugItems  int
		ActualErr   error
	}

	tests := []Test{
		{
			Name:       "Success",
			Input:      domain.NewPayload("John Doe", 1771351985),
			DebugItems: 3,
			EncryptCall: func(key string) (string, error) {
				return "some-token", nil
			},
		},
		{
			Name:       "Encrypt error",
			Input:      domain.NewPayload("John Doe", 1771351985),
			DebugItems: 2,
			EncryptCall: func(key string) (string, error) {
				return "", errors.New("some error")
			},
			ActualErr: ErrCantEncryptKey,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			c := mocks.NewMockCrypter(ctrl)
			c.EXPECT().Encrypt(gomock.Any()).DoAndReturn(tt.EncryptCall)
			l := mocks.NewMockLogger(ctrl)
			l.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
			l.EXPECT().Debugf(gomock.Any(), gomock.Any()).Times(tt.DebugItems)
			g := NewGenerator(c, l)
			token, err := g.Generate(tt.Input)
			if tt.ActualErr != nil {
				require.ErrorIs(t, err, tt.ActualErr)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, token)
			}	
		})
	}
}
