// Copyright 2020 Shivam Rathore
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package lib

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	charset    = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "0123456789"
	charsetLen = int64(len(charset)) // 26+26+10 = 62
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomString returns a current time seeded random string
// of provided length (n), using 62 characters [a-zA-Z0-9]
func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[rand.Int63()%charsetLen]
	}
	return string(b)
}

// GenerateSlug generates a slug (a clean part of a URL for a specific
// address) using an initially provided raw `slug` & a list of existing
// existing slugs (if any), to generate a unique slug.
// It appends numbers to the slug to generate a unique slug in `existing`s
func GenerateSlug(slug string, existing ...string) string {
	// remove all empty spaces if any from the initial slug
	slug = strings.ReplaceAll(slug, " ", "")

	// if the initial slug and processing strings `existing` are not
	// passed, return a random string of length 6 - just a random
	if slug == "" && len(existing) == 0 {
		return RandomString(6)
	}

	// if `existing` is empty, return initial slug
	if len(existing) == 0 {
		return slug
	}

	slugN := len(slug)
	mx := 0
	cannot := map[int]bool{}

	for _, str := range existing {
		// if our slug is not a prefix of `str`, it does not affects us
		if len(str) < slugN || slug != str[:slugN] {
			continue
		}
		// else if the reset of the part if not a number, it is again
		// not going to affect us, as we append numbers to our slug
		suf, err := strconv.Atoi(str[slugN:])
		if err != nil {
			continue
		}
		// store our used number, as it cannot be used
		cannot[suf] = true
		if suf > mx {
			mx = suf
		}
	}

	// Now, use the unused number to generate the slug
	for i := 1; i <= mx; i++ {
		if cannot[i] == false {
			return slug + strconv.Itoa(i)
		}
	}
	return slug + strconv.Itoa(mx+1)
}
