package node

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
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

// Fee calculation function - Used in the M2S Transactions in order to determine Fee based on the following range: [ <10k = .75% ] [ 10k-50k = .5% ] [ >50k = .25% ]
func FeeCalc(x float64) float64 {
	var amount float64

	if x <= 25000 && x > 0 {
		amount = 0.0075 * x
	} else if x >= 25000 || x <= 100000 {
		amount = 0.005 * x
	} else if x > 100000 {
		amount = 0.0025 * x
	}

	return amount
}

func JsonToSeed(data []byte) *Seed {
	var seed Seed
	s := &seed

	json.Unmarshal(data, s)

	return s
}

// Join the Send & Receive transactions into one map for uploading to Redis
func Complete(s Seed, r Seed) map[Seed]Seed {
	m := make(map[Seed]Seed)
	m[s] = r

	return m
}

// Authorize a Tx; Returns true if the difference between the amount sent and the amount received was the total fee charged; If that is not true, it returns false
func Auth(s Seed, r Seed) bool {
	var status bool

	if (math.Round((s.AmountSent-r.AmountReceived)*100) / 100) == s.Fee {
		status = true
	} else {
		status = false
	}

	return status
}

func GenerateURL(s Seed) string {
	if s.AmountReceived == 0 {
		return fmt.Sprintf("https://app.rhizomatiq.dev/transactions/send/%s", s.Hash)
	}

	return fmt.Sprintf("https://app.rhizomatiq.dev/transactions/receive/%s", s.Hash)
}
