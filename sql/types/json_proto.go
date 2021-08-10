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
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

type JsonProto []byte

func (j *JsonProto) FromDB(bytes []byte) error {
	*j = JsonProto(bytes)
	return nil
}

func (j *JsonProto) ToDB() ([]byte, error) {
	if j.IsEmpty() {
		return j.Default(), nil
	}
	return []byte(*j), nil
}

func (j *JsonProto) IsEmpty() bool {
	return len([]byte(*j)) == 0
}

func (j *JsonProto) Default() []byte {
	return []byte("{}")
}

func (j *JsonProto) Marshal(pb proto.Message) (err error) {
	m := jsonpb.Marshaler{}
	s, err := m.MarshalToString(pb)
	*j = []byte(s)
	return
}

func (j *JsonProto) Unmarshal(pb proto.Message) error {
	return jsonpb.UnmarshalString(j.String(), pb)
}

func (j *JsonProto) String() string {
	return string([]byte(*j))
}

func (j *JsonProto) Bytes() []byte {
	return []byte(*j)
}

func NewJsonpb(pb proto.Message) (JsonProto, error) {
	m := jsonpb.Marshaler{}
	s, err := m.MarshalToString(pb)
	if err != nil {
		return nil, err
	}
	return JsonProto([]byte(s)), nil
}

func MustJsonpb(pb proto.Message) JsonProto {
	out, err := NewJsonpb(pb)
	if err != nil {
		panic(err)
	}
	return out
}
