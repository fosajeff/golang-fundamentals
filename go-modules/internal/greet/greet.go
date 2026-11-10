package greet

import "strings"

// this is a reusable logic pacakge not an executable package

// exported function starts with capital letter
// so other packages can call it
func Hello(name string) string {
	clean := normalizeName(name)

	return "Hello, " + clean
}

func normalizeName(name string) string {

	n := strings.TrimSpace(name)

	if n == "" {
		return "Guest"
	}

	return strings.ToUpper(n)
}
