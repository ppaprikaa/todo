package e

import (
	"fmt"
	"strings"
)

// Wraps given error with given msg
// error: "msg: err"
func WrapS(msg string, err error) error {
	var (
		msgNotEmpty = strings.TrimSpace(msg) != ""
		errNotEmpty = err != nil
	)

	if errNotEmpty {
		if strings.TrimSpace(err.Error()) == "" {
			errNotEmpty = false
		}
	}

	if errNotEmpty && msgNotEmpty {
		return fmt.Errorf("%s: %w", msg, err)
	}

	return nil
}

// Wraps given lower error with given upper error
// error: "upper: lower"
func WrapE(upper error, lower error) error {
	var (
		upperNotEmpty = upper != nil
		lowerNotEmpty = lower != nil
	)

	if upperNotEmpty {
		if strings.TrimSpace(upper.Error()) == "" {
			upperNotEmpty = false
		}
	}

	if lowerNotEmpty {
		if strings.TrimSpace(lower.Error()) == "" {
			lowerNotEmpty = false
		}
	}

	if upperNotEmpty && lowerNotEmpty {
		return fmt.Errorf("%w: %w", upper, lower)
	}

	return nil
}
