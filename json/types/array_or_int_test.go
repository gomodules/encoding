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

func TestArrayOrInt_MarshalJSON_Nil(t *testing.T) {
	assert := assert.New(t)

	var a ArrayOrInt
	a = nil
	data, err := a.MarshalJSON()
	assert.Nil(err)
	assert.Equal(`[]`, string(data))
}

func TestArrayOrInt_MarshalJSON_Empty(t *testing.T) {
	assert := assert.New(t)

	var a ArrayOrInt
	data, err := a.MarshalJSON()
	assert.Nil(err)
	assert.Equal(`[]`, string(data))
}

func TestArrayOrInt_MarshalJSON_Single(t *testing.T) {
	assert := assert.New(t)

	var a ArrayOrInt
	a = []int{1}
	data, err := a.MarshalJSON()
	assert.Nil(err)
	assert.Equal(`1`, string(data))
}

func TestArrayOrInt_MarshalJSON_Multiple(t *testing.T) {
	assert := assert.New(t)

	var a ArrayOrInt
	a = []int{1, 2}
	data, err := a.MarshalJSON()
	assert.Nil(err)
	assert.Equal(`[1,2]`, string(data))
}

func TestArrayOrInt_UnmarshalJSON_Empty(t *testing.T) {
	assert := assert.New(t)

	var a ArrayOrInt
	err := a.UnmarshalJSON([]byte(`[]`))
	assert.Nil(err)
	assert.Empty(a)
}

func TestArrayOrInt_UnmarshalJSON_Single(t *testing.T) {
	assert := assert.New(t)

	var a ArrayOrInt
	err := a.UnmarshalJSON([]byte(`1`))
	assert.Nil(err)
	assert.Len(a, 1)
	assert.Equal(1, a[0])
}

func TestArrayOrInt_UnmarshalJSON_Multiple(t *testing.T) {
	assert := assert.New(t)

	var a ArrayOrInt
	err := a.UnmarshalJSON([]byte(`[1,2]`))
	assert.Nil(err)
	assert.Len(a, 2)
	assert.Equal(1, a[0])
	assert.Equal(2, a[1])
}
