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

import "math"

// pow10tab a copy of math.pow10tab for int64 type
// that stores the pre-computed values 10**i for i <= 18.
var pow10tab = [...]int64{
	1e00, 1e01, 1e02, 1e03, 1e04, 1e05, 1e06, 1e07, 1e08, 1e09,
	1e10, 1e11, 1e12, 1e13, 1e14, 1e15, 1e16, 1e17, 1e18,
}

// Pow10 returns Int64 type value of 10**n, the base-10 exponential
// of n in Int64 type.
//
// Special cases are:
//	Pow10(n) =  0 for n < 0
//	Pow10(n) = -1 for n > 18
func Pow10(n int) int64 {
	if n > 18 {
		return -1
	}
	if n < 0 {
		return 0
	}
	return pow10tab[n]
}

// RoundUp will cut short the number of digits after decimal points in `n`
// to only `pre` number of digits.
// `float64` will round up and change the digits when using larger precisions
// `pre` for `n` with more than 10 digits after decimal point.
func RoundUp(n float64, pre int) float64 {
	if pre < 0 {
		return n
	}

	// for keeping same type, math.Pow10 is used
	mul := math.Pow10(pre)
	return float64(int(n*mul)) / mul
}

// Pow wraps the `math.Pow(float64, float64) float64` function to a more
// direct approach for int64 type.
//
// Special case: If a**b comes out to be larger than the `MaxInt64` the
// value will overflow and will not be correct for calculations
func Pow(a, b int64) int64 {
	return int64(math.Pow(float64(a), float64(b)))
}
