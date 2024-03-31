package provider

import "strings"

func isStringEmpty(s string) bool { return len(strings.TrimSpace(s)) == 0 }
