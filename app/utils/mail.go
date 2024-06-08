package utils

import "net/mail"

func IsValidMail(email string) bool {
	_, err := mail.ParseAddress(email)

	return err == nil
}
