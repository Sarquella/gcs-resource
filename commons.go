package gcsresource

import (
	"strings"
)

func DynamicPath(path, replaceKey, replaceValue string) string{
	return strings.Replace(path, replaceKey, replaceValue, -1)
}