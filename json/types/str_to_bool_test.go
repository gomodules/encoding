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

func TestStrToBool_MarshalJSON_True(t *testing.T) {
	assert := assert.New(t)

	var a StrToBool
	a = true
	data, err := a.MarshalJSON()
	assert.Nil(err)
	assert.Equal(`true`, string(data))
}

func TestStrToBool_MarshalJSON_False(t *testing.T) {
	assert := assert.New(t)

	var a StrToBool
	a = false
	data, err := a.MarshalJSON()
	assert.Nil(err)
	assert.Equal(`false`, string(data))
}

func TestStrToBool_MarshalJSON_ZeroValue(t *testing.T) {
	assert := assert.New(t)

	var a StrToBool
	data, err := a.MarshalJSON()
	assert.Nil(err)
	assert.Equal(`false`, string(data))
}

func TestStrToBool_UnmarshalJSON_Empty(t *testing.T) {
	assert := assert.New(t)

	var a StrToBool
	err := a.UnmarshalJSON([]byte(`""`))
	assert.Nil(err)
	assert.Equal(StrToBool(false), a)
}

func TestStrToBool_UnmarshalJSON_NonEmpty(t *testing.T) {
	assert := assert.New(t)

	var a StrToBool
	err := a.UnmarshalJSON([]byte(`"false"`))
	assert.Nil(err)
	assert.Equal(StrToBool(true), a)
}

func TestStrToBool_UnmarshalJSON_True(t *testing.T) {
	assert := assert.New(t)

	var a StrToBool
	err := a.UnmarshalJSON([]byte(`true`))
	assert.Nil(err)
	assert.Equal(StrToBool(true), a)
}

func TestStrToBool_UnmarshalJSON_False(t *testing.T) {
	assert := assert.New(t)

	var a StrToBool
	err := a.UnmarshalJSON([]byte(`false`))
	assert.Nil(err)
	assert.Equal(StrToBool(false), a)
}

func TestStrToBool_UnmarshalJSON_Fail(t *testing.T) {
	assert := assert.New(t)

	var a StrToBool
	err := a.UnmarshalJSON([]byte(`True`))
	assert.NotNil(err)
}
