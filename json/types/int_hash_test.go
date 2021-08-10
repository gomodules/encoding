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
	"testing"

	"gomodules.xyz/encoding/json"
	. "gomodules.xyz/encoding/json/types"

	"github.com/stretchr/testify/assert"
)

func TestIntHash_Empty(t *testing.T) {
	assert := assert.New(t)

	var x IntHash
	err := x.UnmarshalJSON([]byte(`""`))
	assert.Nil(err)
}

func TestIntHash(t *testing.T) {
	assert := assert.New(t)

	type Example struct {
		A IntHash
		B IntHash
		C IntHash
		D IntHash
		E *IntHash
		F *IntHash `json:",omitempty"`
		G IntHash
	}
	s := `{
		"A": "0$str\\",
		"B": 1,
		"C": "8$xyz",
		"E": null
	}`

	var e Example
	err := json.Unmarshal([]byte(s), &e)
	assert.Nil(err)

	b, err := json.Marshal(&e)
	assert.Nil(err)
	assert.Equal(`{"A":"0$str\\","B":1,"C":"8$xyz","D":0,"E":null,"G":0}`, string(b))
}
