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
	"sync"
	"testing"
)

func TestRandomness_RandomString(t *testing.T) {
	var wg sync.WaitGroup

	chn := make(chan string, 1000000)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(c chan<- string) {
			defer wg.Done()

			for j := 0; j < 1000; j++ {
				c <- RandomString(10)
			}
		}(chn)
	}
	wg.Wait()
	close(chn)

	set := map[string]bool{}
	for s := range chn {
		if set[s] {
			t.Errorf("not generating unique string" + s)
		}

		set[s] = true
	}
}

func Test_RandomString(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		len  int
	}{
		{
			name: "Zero length RandomString",
			args: args{
				n: 0,
			},
			len: 0,
		},
		{
			name: "10 length RandomString",
			args: args{
				n: 10,
			},
			len: 10,
		},
		{
			name: "One length RandomString",
			args: args{
				n: 1,
			},
			len: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandomString(tt.args.n); len(got) != tt.len {
				t.Errorf("RandomString() = %v(len=%v), want len = %v", got, len(got), tt.len)
			}
		})
	}
}
