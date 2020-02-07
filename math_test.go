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
	"testing"
)

func TestPow10(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Normal Check",
			args: args{n: 10},
			want: 10000000000,
		},
		{
			name: "Maximum value of n",
			args: args{n: 18},
			want: 1000000000000000000,
		},
		{
			name: "n == 0",
			args: args{},
			want: 1,
		},
		{
			name: "n < 0",
			args: args{n: -112},
			want: 0,
		},
		{
			name: "n > 18",
			args: args{n: 19},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pow10(tt.args.n); got != tt.want {
				t.Errorf("Pow10() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoundUp(t *testing.T) {
	type args struct {
		n   float64
		pre int
	}
	tests := []struct {
		name string
		args args
		want float64
		fail bool
	}{
		{
			name: "Normal Check",
			args: args{
				n:   132.987,
				pre: 2,
			},
			want: 132.98,
		},
		{
			name: "Normal Check",
			args: args{
				n:   0.987,
				pre: -2,
			},
			want: 0.987,
		},
		{
			name: "Success in More than 10 Digits after Decimals",
			args: args{
				n:   643.3987657283324,
				pre: 12,
			},
			want: 643.398765728332,
			fail: false,
		},
		{
			name: "Failure in More than 10 Digits after Decimals",
			args: args{
				n:   643.9876527398439,
				pre: 12,
			},
			want: 643.987652739843,
			fail: true, // output will be 643.987652739844
		},
		{
			name: "Failure in More than 10 Digits after Decimals with -ve precision",
			args: args{
				n:   643.9876527398439,
				pre: -2,
			},
			want: 643.9876527398439,
			fail: true, // output will be 643.987652739844
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RoundUp(tt.args.n, tt.args.pre); !tt.fail && got != tt.want {
				t.Errorf("RoundUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPow(t *testing.T) {
	type args struct {
		a int64
		b int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Normal Check",
			args: args{
				a: 2,
				b: 10,
			},
			want: 1024,
		},
		{
			name: "Int64 overflow",
			args: args{
				a: 2,
				b: 100,
			},
			want: -9223372036854775808, // output will be something like this
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pow(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Pow() = %v, want %v", got, tt.want)
			}
		})
	}
}
