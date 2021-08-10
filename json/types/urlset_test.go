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

	. "gomodules.xyz/encoding/json/types"

	"github.com/stretchr/testify/assert"
	"gomodules.xyz/sets"
)

func TestURLSet_MarshalJSON_Nil(t *testing.T) {
	assert := assert.New(t)

	var a *URLSet
	a = nil
	data, err := a.MarshalJSON()
	assert.Nil(err)
	assert.Equal(`""`, string(data))
}

func TestURLSet_MarshalJSON_Empty(t *testing.T) {
	assert := assert.New(t)

	var a URLSet
	data, err := a.MarshalJSON()
	assert.Nil(err)
	assert.Equal(`""`, string(data))
}

func TestURLSet_MarshalJSON_Single(t *testing.T) {
	assert := assert.New(t)

	a := NewURLSet("https", 2380)
	a.Insert("127.0.0.1")
	data, err := a.MarshalJSON()
	assert.Nil(err)
	assert.Equal(`"https://127.0.0.1:2380"`, string(data))
}

func TestURLSet_MarshalJSON_Multiple(t *testing.T) {
	assert := assert.New(t)

	a := NewURLSet("https", 2380)
	a.Insert("127.0.0.1", "127.0.0.2")
	data, err := a.MarshalJSON()
	assert.Nil(err)
	assert.Equal(`"https://127.0.0.1:2380,https://127.0.0.2:2380"`, string(data))
}

func TestURLSet_UnmarshalJSON_Empty(t *testing.T) {
	assert := assert.New(t)

	var a URLSet
	err := a.UnmarshalJSON([]byte(`""`))
	assert.Nil(err)
	assert.True(a.Equal(URLSet{
		Scheme: "",
		Hosts:  sets.NewString(),
		Port:   0,
	}))
}

func TestURLSet_UnmarshalJSON_Single(t *testing.T) {
	assert := assert.New(t)

	var a URLSet
	err := a.UnmarshalJSON([]byte(`"https://127.0.0.1:2380"`))
	assert.Nil(err)
	assert.True(a.Equal(URLSet{
		Scheme: "https",
		Hosts:  sets.NewString("127.0.0.1"),
		Port:   2380,
	}))
}

func TestURLSet_UnmarshalJSON_Multiple(t *testing.T) {
	assert := assert.New(t)

	var a URLSet
	err := a.UnmarshalJSON([]byte(`"https://127.0.0.1:2380,https://127.0.0.2:2380"`))
	assert.Nil(err)
	assert.True(a.Equal(URLSet{
		Scheme: "https",
		Hosts:  sets.NewString("127.0.0.1", "127.0.0.2"),
		Port:   2380,
	}))
}
