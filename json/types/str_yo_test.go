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

package types_test

import (
	"fmt"
	"testing"

	"gomodules.xyz/encoding/json"
	. "gomodules.xyz/encoding/json/types"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func TestStrYo(t *testing.T) {
	assert := assert.New(t)

	type Example struct {
		A StrYo
		B StrYo
		C StrYo
		D StrYo
		E StrYo
		F StrYo `json:",omitempty"`
		G StrYo
	}
	s := `{
		"A": "str\\",
		"B": 1,
		"C": 2.5,
		"D": false,
		"E": true,
		"F": null,
		"G": "8.0"
	}`

	var e Example
	err := json.Unmarshal([]byte(s), &e)

	assert.Nil(err)
	b, err := json.Marshal(&e)
	fmt.Println(string(b))
	assert.Equal(`{"A":"str\\","B":"1","C":"2.5","D":"false","E":"true","G":"8.0"}`, string(b))
}

func TestStrYo2(t *testing.T) {
	assert := assert.New(t)
	json := jsoniter.ConfigCompatibleWithStandardLibrary

	type Example struct {
		A StrYo
		B StrYo
		C StrYo
		D StrYo
		E StrYo
		F StrYo `json:",omitempty"`
		G StrYo
	}
	s := `{
		"A": "str\\",
		"B": 1,
		"C": 2.5,
		"D": false,
		"E": true,
		"F": null,
		"G": "8.0"
	}`

	var e Example
	err := json.Unmarshal([]byte(s), &e)

	assert.Nil(err)
	b, err := json.Marshal(&e)
	fmt.Println(string(b))
	assert.Equal(`{"A":"str\\","B":"1","C":"2.5","D":"false","E":"true","G":"8.0"}`, string(b))
}
