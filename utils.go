package node

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// Valid timestamps for this library must be in this format - 20221225-123045
func Now(t time.Time) string {
	str := fmt.Sprintf("%d%d%d-%d%d%d",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second())

	return str
}

// General hashing function for this library uses Sha 256
func Hash(in string) string {
	h := sha256.New()
	h.Write([]byte(in))

	return string(hex.EncodeToString(h.Sum(nil)))
}
