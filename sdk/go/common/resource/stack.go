// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package resource

import (
	"github.com/pulumi/pulumi/sdk/go/common/tokens"
)

// RootStackType is the type name that will be used for the root component in the Pulumi resource tree.
const RootStackType tokens.Type = "pulumi:pulumi:Stack"

// DefaultRootStackURN constructs a default root stack URN for the given stack and project.
func DefaultRootStackURN(stack tokens.QName, proj tokens.PackageName) URN {
	return NewURN(stack, proj, "", RootStackType, tokens.QName(string(proj)+"-"+string(stack)))
}
