package shortener

import (
	"crypto/md5"
	"encoding/hex"
)

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func UrlShorten(url string) string {
	// This function will host the URL Shortner logic
	urlShort := getMD5Hash(url)
	return urlShort[:7]
}
