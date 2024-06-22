package transport

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCP(t *testing.T) {
	listenAddr := ":13000"
	transport := NewTCP(listenAddr)

	assert.Equal(t, transport.address, listenAddr)

	assert.Nil(t, transport.ListenAndAccept())
}
