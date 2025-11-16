package senders

import "net/mail"

// IsValidMail checks if the provided email string is a valid mail address.
func IsValidMail(email string) bool {
	_, err := mail.ParseAddress(email)

	return err == nil
}
