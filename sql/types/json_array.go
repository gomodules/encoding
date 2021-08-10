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

type JsonArr []byte

func (j *JsonArr) FromDB(bytes []byte) error {
	*j = JsonArr(bytes)
	return nil
}

func (j *JsonArr) ToDB() ([]byte, error) {
	if j.IsEmpty() {
		return j.Default(), nil
	}
	return []byte(*j), nil
}

func (j *JsonArr) IsEmpty() bool {
	return len([]byte(*j)) == 0
}

func (j *JsonArr) Default() []byte {
	return []byte("[]")
}

func (j *JsonArr) Marshal(v interface{}) (err error) {
	*j, err = json.Marshal(v)
	return
}

func (j *JsonArr) Unmarshal(v interface{}) error {
	return json.Unmarshal([]byte(*j), v)
}

func (j *JsonArr) String() string {
	return string([]byte(*j))
}

func (j *JsonArr) Bytes() []byte {
	return []byte(*j)
}

func NewJsonArr(v interface{}) (JsonArr, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return JsonArr(b), nil
}

func MustJsonArr(v interface{}) JsonArr {
	out, err := NewJsonArr(v)
	if err != nil {
		panic(err)
	}
	return out
}
