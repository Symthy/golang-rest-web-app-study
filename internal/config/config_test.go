package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	wantPort := 3333
	t.Setenv("PORT", fmt.Sprint(wantPort))

	got, err := New()
	require.NoError(t, err)
	require.Equal(t, wantPort, got.Port)

	wantEnv := "dev"
	require.Equal(t, wantEnv, got.Env)
}
