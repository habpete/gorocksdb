package gorocksdb

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_constructor(t *testing.T) {
	inst, err := New()
	require.NoError(t, err)
	require.NotNil(t, inst)
}
