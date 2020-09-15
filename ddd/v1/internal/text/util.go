package text

import (
	"strings"
	"unicode"
)

// Safename returns a lowercase name which just contains a..z, nothing else.
func Safename(str string) string {
	str = strings.ToLower(str)
	sb := &strings.Builder{}
	for _, r := range str {
		if r >= 'a' && r <= 'z' {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

// JoinSlashes assembles the path segments and ensures that they have only 1 slash per segment. Leading
// or trailing slashes are purged.
func JoinSlashes(paths ...string) string {
	sb := &strings.Builder{}
	for i, path := range paths {
		path = TrimSlashes(path)
		sb.WriteString(path)
		if i < len(paths)-1 {
			sb.WriteRune('/')
		}
	}

	return sb.String()
}

// TrimSlashes removes leading and trailing slashes
func TrimSlashes(str string) string {
	for strings.HasPrefix(str, "/") {
		str = str[1:]
	}

	for strings.HasSuffix(str, "/") {
		str = str[:len(str)-1]
	}

	return str
}

// MakePublic converts aBc to ABc.
func MakePublic(str string) string {
	if len(str) == 0 {
		return str
	}

	return string(unicode.ToUpper(rune(str[0]))) + str[1:]
}

// CamelCaseToWords converts a text like MyBookLibrary into "my book library"
func CamelCaseToWords(cc string) string {
	sb := &strings.Builder{}
	for i, r := range cc {
		if unicode.IsUpper(r) {
			if i > 0 {
				sb.WriteRune(' ')
			}
			sb.WriteRune(unicode.ToLower(r))
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

// TrimComment removes '...' and any whitespace afterwards.
func TrimComment(str string) string {
	str = strings.TrimSpace(str)
	if strings.HasPrefix(str, "...") {
		str = str[3:]
	}
	return strings.TrimSpace(str)
}

// ParseVerb checks for Get/Post/Head/Patch/Put/Delete verbs and returns
// them in uppercase. If no verb is found, an empty string is returned.
func ParseVerb(text string) string {
	text = strings.ToLower(text)
	if strings.HasPrefix(text, "get") {
		return "GET"
	}

	if strings.HasPrefix(text, "post") {
		return "POST"
	}

	if strings.HasPrefix(text, "delete") {
		return "DELETE"
	}

	if strings.HasPrefix(text, "put") {
		return "PUT"
	}

	if strings.HasPrefix(text, "head") {
		return "head"
	}

	return ""
}
