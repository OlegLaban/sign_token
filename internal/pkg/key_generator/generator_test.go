package keygenerator

import (
	"testing"

	"github.com/OlegLaban/sing_token/internal/domain"
	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	type Test struct {
		Name      string
		Input     domain.Payload
		ActualErr error
	}

	tests := []Test{
		{
			Name:  "Success",
			Input: domain.NewPayload("John Doe", 1771351985),
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			g := NewGenerator()

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
