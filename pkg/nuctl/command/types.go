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

package command

import (
	"strings"

	"github.com/nuclio/nuclio/pkg/functionconfig"
	"github.com/nuclio/nuclio/pkg/platform"
)

type stringSliceFlag []string

func (ssf *stringSliceFlag) String() string {
	return strings.Join(*ssf, ", ")
}

func (ssf *stringSliceFlag) Set(value string) error {
	*ssf = append(*ssf, value)
	return nil
}

func (ssf *stringSliceFlag) Type() string {
	return "String"
}

type ProjectImportConfig struct {
	Project        *platform.ProjectConfig
	Functions      map[string]*functionconfig.Config
	FunctionEvents map[string]*platform.FunctionEventConfig
	APIGateways    map[string]*platform.APIGatewayConfig
}

type ProjectImportOptions struct {
	projectImportConfig *ProjectImportConfig
}
