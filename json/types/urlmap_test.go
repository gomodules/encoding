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
)

func TestURLMap_MarshalJSON_Nil(t *testing.T) {
	assert := assert.New(t)

	var a *URLMap
	a = nil
	data, err := a.MarshalJSON()
	assert.Nil(err)
	assert.Equal(`""`, string(data))
}

func TestURLMap_MarshalJSON_Empty(t *testing.T) {
	assert := assert.New(t)

	var a URLMap
	data, err := a.MarshalJSON()
	assert.Nil(err)
	assert.Equal(`""`, string(data))
}

func TestURLMap_MarshalJSON_Single(t *testing.T) {
	assert := assert.New(t)

	a := NewURLMap("https", 2380)
	a.Insert("s1", "127.0.0.1")
	data, err := a.MarshalJSON()
	assert.Nil(err)
	assert.Equal(`"s1=https://127.0.0.1:2380"`, string(data))
}

func TestURLMap_MarshalJSON_Multiple(t *testing.T) {
	assert := assert.New(t)

	a := NewURLMap("https", 2380)
	a.Insert("s1", "127.0.0.1")
	a.Insert("s2", "127.0.0.2")
	data, err := a.MarshalJSON()
	assert.Nil(err)
	assert.Equal(`"s1=https://127.0.0.1:2380,s2=https://127.0.0.2:2380"`, string(data))
}

func TestURLMap_UnmarshalJSON_Empty(t *testing.T) {
	assert := assert.New(t)

	var a URLMap
	err := a.UnmarshalJSON([]byte(`""`))
	assert.Nil(err)
	assert.True(a.Equal(URLMap{
		Scheme: "",
		Hosts:  nil,
		Port:   0,
	}))
}

func TestURLMap_UnmarshalJSON_Single(t *testing.T) {
	assert := assert.New(t)

	var a URLMap
	err := a.UnmarshalJSON([]byte(`"s1=https://127.0.0.1:2380"`))
	assert.Nil(err)
	assert.True(a.Equal(URLMap{
		Scheme: "https",
		Hosts:  map[string]string{"s1": "127.0.0.1"},
		Port:   2380,
	}))
}

func TestURLMap_UnmarshalJSON_Multiple(t *testing.T) {
	assert := assert.New(t)

	var a URLMap
	err := a.UnmarshalJSON([]byte(`"s1=https://127.0.0.1:2380,s2=https://127.0.0.2:2380"`))
	assert.Nil(err)
	assert.True(a.Equal(URLMap{
		Scheme: "https",
		Hosts: map[string]string{
			"s1": "127.0.0.1",
			"s2": "127.0.0.2",
		},
		Port: 2380,
	}))
}
