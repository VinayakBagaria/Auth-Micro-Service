package security

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptPassword(t *testing.T) {
	pass, err := EncrytPassword("123456789")
	assert.NoError(t, err)
	assert.NotEmpty(t, pass)
	assert.Len(t, pass, 60)
}

func TestVerifyPassword(t *testing.T) {
	pass, err := EncrytPassword("123456789")
	assert.NoError(t, err)
	assert.NotEmpty(t, pass)
	assert.Len(t, pass, 60)

	assert.NoError(t, VerifyPassword(pass, "123456789"))
	assert.Error(t, VerifyPassword(pass, pass))
	assert.Error(t, VerifyPassword("123456789", pass))
}
