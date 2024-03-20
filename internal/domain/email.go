package domain

import (
	"errors"
	"fmt"
	"net/mail"
	"regexp"
	"strings"
	"unicode/utf8"
)

const emailMaxLength = 100

var (
	invalidEmailChars = regexp.MustCompile(`[^a-zA-Z0-9+.@_~\-]`)
	validEmailSeq     = regexp.MustCompile(`^[a-zA-Z0-9+._~\-]+@[a-zA-Z0-9+._~\-]+(\.[a-zA-Z0-9+._~\-]+)+$`)
)

type Email string

func NewEmail(email string) (Email, error) {
	if strings.TrimSpace(email) == "" {
		return "", errors.New("email cannot be empty")
	}

	if strings.ContainsAny(email, " \t\n\r") {
		return "", errors.New("email cannot contain whitespace")
	}

	if strings.ContainsAny(email, `"'`) {
		return "", errors.New("email cannot contain quotes")
	}

	if rc := utf8.RuneCountInString(email); rc > emailMaxLength {
		return "", fmt.Errorf("email cannot be a over %v characters in length", emailMaxLength)
	}

	addr, err := mail.ParseAddress(email)
	if err != nil {
		return "", err
	}

	if addr.Name != "" {
		return "", errors.New("email doesn't include a name")
	}

	if matches := invalidEmailChars.FindAllString(addr.Address, -1); len(matches) != 0 {
		return "", fmt.Errorf("email cannot contain %v", matches)
	}

	if !validEmailSeq.MatchString(addr.Address) {
		_, end, _ := strings.Cut(addr.Address, "@")

		if !strings.Contains(end, ".") {
			return "", errors.New("email doesn't have top-level domain, e.g. .com, .co.uk, etc")
		}

		return "", errors.New("email must be a valid, e.g. email@example.com")
	}

	return Email(addr.Address), nil
}

func (e Email) String() string {
	return string(e)
}
