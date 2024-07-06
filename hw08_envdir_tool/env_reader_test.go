package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("positive case get env map", func(t *testing.T) {
		environments, err := ReadDir("./testdata/env")
		require.NoError(t, err)
		require.Equal(t, Environment{
			"BAR":   EnvValue{Value: "bar", NeedRemove: false},
			"EMPTY": EnvValue{Value: "", NeedRemove: false},
			"FOO":   EnvValue{Value: "   foo\nwith new line", NeedRemove: false},
			"HELLO": EnvValue{Value: "\"hello\"", NeedRemove: false},
			"UNSET": EnvValue{Value: "", NeedRemove: true},
		}, environments)
	})

	t.Run("negative case get env map from uexisting folder", func(t *testing.T) {
		_, err := ReadDir("./testdata/env2")
		require.EqualError(t, err, "error when trying to read the directory: open ./testdata/env2: no such file or directory")
	})
}
