/*
Copyright 2016 The Kubernetes Authors.

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

package apimachinery

import (
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/sets"
)

// GroupMeta stores the metadata of a group.
type GroupMeta struct {
	// GroupVersions is Group + all versions in that group.
	GroupVersions []schema.GroupVersion

	RootScopedKinds sets.String

	// RESTMapper provides the default mapping between REST paths and the objects declared in a Scheme and all known
	// versions.
	RESTMapper meta.RESTMapper
}
