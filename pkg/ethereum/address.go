package ethereum

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrInvalidAddressLength = errors.New("invalid address length")
	ErrInvalidAddressFormat = errors.New("invalid address format")
)

// ValidateAddress checks if the given string is a valid Ethereum address.
func ValidateAddress(address string) error {
	address = strings.TrimPrefix(address, "0x")

	if len(address) != 40 {
		return ErrInvalidAddressLength
	}

	match, _ := regexp.MatchString("^[0-9a-fA-F]{40}$", address)
	if !match {
		return ErrInvalidAddressFormat
	}

	return nil
}
