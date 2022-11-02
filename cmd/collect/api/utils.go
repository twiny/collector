package api

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

// hashURL
func hashURL(u string) string {
	sum := sha256.Sum224([]byte(u))
	return hex.EncodeToString(sum[:])
}

// pageTitle
func pageTitle(u string) string {
	//
	return strings.ReplaceAll(u, " ", "_")
}
