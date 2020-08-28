package validation

import (
	"fmt"
	"unicode"
)

func isPublicGoIdentifier(str string) bool {
	if !isGoIdentifier(str) {
		return false
	}

	for _, r := range str {
		return unicode.IsUpper(r)
	}

	panic("illegal state")
}

func isPrivateGoIdentifier(str string) bool {
	if !isGoIdentifier(str) {
		return false
	}

	for _, r := range str {
		return unicode.IsLower(r)
	}

	panic("illegal state")
}

func isGoPackageName(str string) bool {
	return isPrivateGoIdentifier(str)
}

// isGoIdentifier returns true for things like aBc or Abc or Abc1 but false for 1abc or a_bc.
func isGoIdentifier(str string) bool {
	if len(str) == 0 {
		return false
	}

	for i, r := range str {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (i > 0 && r >= '0' && r <= '1')) {
			return false
		}
	}

	return true
}

func startsUppercase(str string) bool {
	for _, r := range str {
		return unicode.IsUpper(r)
	}

	return false
}

func startsLowercase(str string) bool {
	for _, r := range str {
		return unicode.IsLower(r)
	}

	return false
}

func buildErr(property, value, msg string, p withPos) error {
	if len(value) > maxValueCitateLength {
		value = value[0:maxValueCitateLength-3] + "(...)"
	}
	return fmt.Errorf("%s.%s '%s': %s: %s:%d", p.Pos().Name, property, value, msg, p.Pos().File, p.Pos().Line)
}
