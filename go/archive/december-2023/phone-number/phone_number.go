// Package phone number cleans up user-entered phone numbers so that they can be sent SMS messages.
package phonenumber

import (
	"errors"
	"fmt"
	"strings"
)

// removeNonNumerics returns an input string after removing all non-numeric characters.
func removeNonNumerics(phoneNumber string) string {
	var result strings.Builder
	for i := 0; i < len(phoneNumber); i++ {
		b := phoneNumber[i]
		if '0' <= b && b <= '9' {
			result.WriteByte(b)
		}
	}

	return result.String()
}

// Number cleans up differently formatted telephone numbers by removing punctuation and the country
// code (1) if present.
func Number(phoneNumber string) (string, error) {
	raw := removeNonNumerics(phoneNumber)

	if len(raw) > 11 || len(raw) < 10 {
		return "", errors.New("not a valid phone number length")
	} else if len(raw) == 11 { // A length of 11 means a phone number + country code
		if raw[0] != '1' {
			return "", errors.New("not a valid phone number country code")
		} else {
			// Remove the country code from the phone number
			raw = raw[1:]
		}
	}

	if raw[0] < '2' || raw[3] < '2' {
		return "", errors.New("area codes and exchange codes must be digits between 2 and 9")
	}

	return raw, nil
}

// AreaCode returns only th area code of the input phone number
func AreaCode(phoneNumber string) (string, error) {
	raw, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}

	return raw[0:3], nil
}

// Format produces the output like "(613) 995-0253"
func Format(phoneNumber string) (string, error) {
	raw, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}

	area := raw[0:3]
	exchange := raw[3:6]
	local := raw[6:]

	return fmt.Sprintf("(%s) %s-%s", area, exchange, local), nil
}
