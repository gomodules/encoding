/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package json

import (
	"testing"
)

func Test_extract(t *testing.T) {
	tests := []struct {
		name       string
		arg        string
		wantTag    string
		wantInline bool
		wantExists bool
	}{
		{
			name:       "No Tag",
			arg:        "",
			wantTag:    "",
			wantInline: false,
			wantExists: false,
		},
		{
			name:       "Regular tag without omitempty",
			arg:        "inline",
			wantTag:    "inline",
			wantInline: false,
			wantExists: true,
		},
		{
			name:       "Regular Tag with omitempty",
			arg:        "inline,omitempty",
			wantTag:    "inline",
			wantInline: false,
			wantExists: true,
		},
		{
			name:       "Inline tag without omitempty",
			arg:        ",inline",
			wantTag:    "",
			wantInline: true,
			wantExists: true,
		},
		{
			name:       "Inline tag with inline,omitempty",
			arg:        ",inline,omitempty",
			wantTag:    "",
			wantInline: true,
			wantExists: true,
		},
		{
			name:       "Inline tag with omitempty,inline",
			arg:        ",omitempty,inline",
			wantTag:    "",
			wantInline: true,
			wantExists: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTag, gotInline, gotExists := extract(tt.arg)
			if gotTag != tt.wantTag {
				t.Errorf("extract() gotTag = %v, want %v", gotTag, tt.wantTag)
			}
			if gotInline != tt.wantInline {
				t.Errorf("extract() gotInline = %v, want %v", gotInline, tt.wantInline)
			}
			if gotExists != tt.wantExists {
				t.Errorf("extract() gotExists = %v, want %v", gotExists, tt.wantExists)
			}
		})
	}
}
