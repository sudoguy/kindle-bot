package senders

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidMail_ValidEmail(t *testing.T) {
	email := "test@example.com"

	isValid := IsValidMail(email)

	assert.True(t, isValid, "Expected %s to be a valid email", email)
}

func TestIsValidMail_InvalidEmail(t *testing.T) {
	email := "invalid"

	isValid := IsValidMail(email)

	assert.False(t, isValid, "Expected %s to be an invalid email", email)
}
