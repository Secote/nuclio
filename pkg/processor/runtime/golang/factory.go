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

package golang

import (
	"os"

	"github.com/nuclio/nuclio/pkg/processor/runtime"

	"github.com/nuclio/logger"
)

type factory struct{}

func (f *factory) Create(parentLogger logger.Logger,
	runtimeConfiguration *runtime.Configuration) (runtime.Runtime, error) {

	// temporarily, for backwards compatibility until this is injected from builder
	runtimeConfiguration.Spec.Build.Path = f.handlerPluginPath()

	return NewRuntime(parentLogger.GetChild("golang"),
		runtimeConfiguration,
		&pluginHandlerLoader{
			abstractHandler{
				logger: parentLogger,
			},
		})
}

func (f *factory) handlerPluginPath() string {
	handlerDLL := os.Getenv("NUCLIO_GOLANG_HANDLER_PLUGIN_PATH")
	if handlerDLL != "" {
		return handlerDLL
	}

	return "/opt/nuclio/handler.so"
}

// register factory
func init() {
	runtime.RegistrySingleton.Register("golang", &factory{})
}
