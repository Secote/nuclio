/*
Copyright 2023 The Nuclio Authors.

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

package functiontemplates

import "github.com/nuclio/nuclio/pkg/functionconfig"

type FunctionTemplate struct {
	Name                   string
	DisplayName            string
	SourceCode             string
	FunctionConfigTemplate string
	FunctionConfigValues   map[string]interface{}
	FunctionConfig         *functionconfig.Config
	serializedTemplate     []byte
}

type generatedFunctionTemplate struct {
	Name               string
	DisplayName        string
	Configuration      functionconfig.Config
	SourceCode         string
	serializedTemplate []byte
}

// keeps the file contents of a function template
type FunctionTemplateFileContents struct {
	Code     string
	Template string
	Values   string
}
