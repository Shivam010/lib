package lib

import (
	"math/rand"
	"strconv"
	"time"
)

const (
	charset    = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "0123456789"
	charsetLen = int64(len(charset))
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[rand.Int63()%charsetLen]
	}
	return string(b)
}

func GenerateSlug(slug string, using ...string) string {
	// initial slug and processing strings `using` are not passed
	// so, returning random string of length 4
	if slug == "" && len(using) == 0 {
		return RandomString(4)
	}

	slugN := len(slug)
	mx := 0
	mp := map[int]bool{}

	for _, str := range using {
		if len(str) < slugN || slug != str[:slugN] {
			continue
		}
		suf, err := strconv.Atoi(str[slugN:])
		if err != nil {
			continue
		}
		mp[suf] = true
		if suf > mx {
			mx = suf
		}
	}
	for i := 1; i <= mx; i++ {
		if mp[i] == false {
			return slug + strconv.Itoa(i)
		}
	}
	return slug + strconv.Itoa(mx+1)
}
