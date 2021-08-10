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

type JsonMap map[string]interface{}

func (j *JsonMap) FromDB(bytes []byte) error {
	*j = make(map[string]interface{})
	return json.Unmarshal(bytes, j)
}

func (j *JsonMap) ToDB() ([]byte, error) {
	if len(*j) == 0 {
		return j.Default(), nil
	}
	return json.Marshal(j)
}

func (j *JsonMap) Default() []byte {
	return []byte("{}")
}

func (j *JsonMap) String() string {
	bytes, err := j.ToDB()
	if err != nil {
		return string(j.Default())
	}
	return string(bytes)
}
