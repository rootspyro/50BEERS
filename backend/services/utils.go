package services

import "strings"

func ParsePublicId(name string) string {
	return strings.ReplaceAll(name, " ", "_")
}
