package greet

import (
	"errors"
	"fmt"
	"strings"
)

var ErrEmptyName = errors.New("greet: name is empty")

func Hello(name string) (string, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return "", ErrEmptyName
	}

	return fmt.Sprintf("Hello, %s!", name), nil
}
