package utils

import (
	"math/rand"
	"os"
	"strings"
	"time"
)

//generate random string
func RandString() string {
	n := 6
	rand.Seed(time.Now().UnixNano())
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Int63() % int64(len(letterBytes))]
    }
    return string(b)
}

//check if a file exist
func Exist(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

//check if protocol being used is allowed
func IsProtocolAllowed(url string, protocols []string) bool {
	for _, proto := range protocols {
		if strings.HasPrefix(url, proto) {
			return true
		}
	}

	return false
} 