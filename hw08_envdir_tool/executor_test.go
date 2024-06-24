package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("positive test run command", func(t *testing.T) {
		exitCode := RunCmd([]string{"pwd"}, Environment{})
		require.Equal(t, 0, exitCode)
	})

	t.Run("negative test run command", func(t *testing.T) {
		exitCode := RunCmd([]string{}, Environment{})
		require.Equal(t, 1, exitCode)
	})
}
