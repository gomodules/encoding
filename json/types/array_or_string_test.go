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

func TestArrayOrString_MarshalJSON_Nil(t *testing.T) {
	assert := assert.New(t)

	var a ArrayOrString
	a = nil
	data, err := a.MarshalJSON()
	assert.Nil(err)
	assert.Equal(`[]`, string(data))
}

func TestArrayOrString_MarshalJSON_Empty(t *testing.T) {
	assert := assert.New(t)

	var a ArrayOrString
	data, err := a.MarshalJSON()
	assert.Nil(err)
	assert.Equal(`[]`, string(data))
}

func TestArrayOrString_MarshalJSON_Single(t *testing.T) {
	assert := assert.New(t)

	var a ArrayOrString
	a = []string{"x"}
	data, err := a.MarshalJSON()
	assert.Nil(err)
	assert.Equal(`"x"`, string(data))
}

func TestArrayOrString_MarshalJSON_Multiple(t *testing.T) {
	assert := assert.New(t)

	var a ArrayOrString
	a = []string{"x", "y"}
	data, err := a.MarshalJSON()
	assert.Nil(err)
	assert.Equal(`["x","y"]`, string(data))
}

func TestArrayOrString_UnmarshalJSON_Empty(t *testing.T) {
	assert := assert.New(t)

	var a ArrayOrString
	err := a.UnmarshalJSON([]byte(`[]`))
	assert.Nil(err)
	assert.Empty(a)
}

func TestArrayOrString_UnmarshalJSON_Single(t *testing.T) {
	assert := assert.New(t)

	var a ArrayOrString
	err := a.UnmarshalJSON([]byte(`"x"`))
	assert.Nil(err)
	assert.Len(a, 1)
	assert.Equal("x", a[0])
}

func TestArrayOrString_UnmarshalJSON_Multiple(t *testing.T) {
	assert := assert.New(t)

	var a ArrayOrString
	err := a.UnmarshalJSON([]byte(`["x","y"]`))
	assert.Nil(err)
	assert.Len(a, 2)
	assert.Equal("x", a[0])
	assert.Equal("y", a[1])
}
