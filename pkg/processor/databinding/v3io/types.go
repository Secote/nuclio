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

package v3io

import (
	"github.com/nuclio/nuclio/pkg/functionconfig"
	"github.com/nuclio/nuclio/pkg/processor/databinding"

	"github.com/mitchellh/mapstructure"
	"github.com/nuclio/errors"
)

type Configuration struct {
	databinding.Configuration
	NumWorkers int
	Username   string
	Password   string
}

func NewConfiguration(id string, databindingConfiguration *functionconfig.DataBinding) (*Configuration, error) {
	newConfiguration := Configuration{}

	// create base
	newConfiguration.Configuration = *databinding.NewConfiguration(id, databindingConfiguration)

	// parse attributes
	if err := mapstructure.Decode(newConfiguration.Configuration.Attributes, &newConfiguration); err != nil {
		return nil, errors.Wrap(err, "Failed to decode attributes")
	}

	// set default num workers
	if newConfiguration.NumWorkers == 0 {
		newConfiguration.NumWorkers = 8
	}

	return &newConfiguration, nil
}
