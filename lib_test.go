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

import "testing"

func TestGenerateSlug(t *testing.T) {
	type args struct {
		slug     string
		existing []string
	}
	tests := []struct {
		name   string
		args   args
		want   string
		random bool
	}{
		{
			name:   "Nil or empty request",
			args:   args{},
			want:   "SomeRandomString",
			random: true,
		},
		{
			name: "No slug provided",
			args: args{
				existing: []string{"lib", "library", "lib1"},
			},
			want:   "SomeRandomString",
			random: true,
		},
		{
			name: "Only Slug provided",
			args: args{
				slug: "lib",
			},
			want:   "lib",
			random: false,
		},
		{
			name: "Normal",
			args: args{
				slug:     "lib",
				existing: []string{"lib", "library", "lib23", "rary"},
			},
			want:   "lib1",
			random: false,
		},
		{
			name: "Normal",
			args: args{
				slug:     "lib",
				existing: []string{"lib", "lib1", "lib23", "lib4", "library"},
			},
			want:   "lib2",
			random: false,
		},
		{
			name: "No matching existing",
			args: args{
				slug:     "lib",
				existing: []string{"school", "play", "ground", "code"},
			},
			want:   "lib1",
			random: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateSlug(tt.args.slug, tt.args.existing...)
			if (!tt.random && got != tt.want) || (tt.random && got == "") {
				t.Errorf("GenerateSlug() = %v, want %v, random %v", got, tt.want, tt.random)
			}
		})
	}
}
