package fasthttp

import (
	"path"
	"strings"
)

func getPathParts(uri string) []string {
	uri = path.Clean(uri)

	parts := strings.Split(uri, "/")

	if len(parts) > 0 && parts[0] == "" {
		parts = parts[1:]
	}

	return parts
}
