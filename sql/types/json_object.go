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

package types

import (
	"gomodules.xyz/encoding/json"
)

type JsonObj []byte

func (j *JsonObj) FromDB(bytes []byte) error {
	*j = JsonObj(bytes)
	return nil
}

func (j *JsonObj) ToDB() ([]byte, error) {
	if j.IsEmpty() {
		return j.Default(), nil
	}
	return []byte(*j), nil
}

func (j *JsonObj) IsEmpty() bool {
	return len([]byte(*j)) == 0
}

func (j *JsonObj) Default() []byte {
	return []byte("{}")
}

func (j *JsonObj) Marshal(v interface{}) (err error) {
	*j, err = json.Marshal(v)
	return
}

func (j *JsonObj) Unmarshal(v interface{}) error {
	return json.Unmarshal([]byte(*j), v)
}

func (j *JsonObj) String() string {
	return string([]byte(*j))
}

func (j *JsonObj) Bytes() []byte {
	return []byte(*j)
}

func NewJsonObj(v interface{}) (JsonObj, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return JsonObj(b), nil
}

func MustJsonObj(v interface{}) JsonObj {
	out, err := NewJsonObj(v)
	if err != nil {
		panic(err)
	}
	return out
}
