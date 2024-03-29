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

package restful

import (
	"encoding/json"
	"net/http"
)

//
// Encoder
//

type jsonEncoder struct {
	jsonEncoder    *json.Encoder
	responseWriter http.ResponseWriter
	resourceType   string
}

// encode a single resource
func (je *jsonEncoder) EncodeResource(resourceID string, resourceAttributes Attributes) {
	je.jsonEncoder.Encode(&resourceAttributes) // nolint: errcheck
}

// encode multiple resources
func (je *jsonEncoder) EncodeResources(resources map[string]Attributes) {
	var resourceIDList []string

	// if attributes is nil, we return a list
	for resourceID, resourceAttributes := range resources {

		// if there's attributes, don't return as a list
		if resourceAttributes != nil {
			break
		}

		resourceIDList = append(resourceIDList, resourceID)
	}

	// if we populated a list, return it as a simple list, otherwise as a map
	if len(resourceIDList) != 0 {
		je.jsonEncoder.Encode(&resourceIDList) // nolint: errcheck
	} else {
		je.jsonEncoder.Encode(&resources) // nolint: errcheck
	}
}

//
// Factory
//

type JSONEncoderFactory struct{}

func (jef *JSONEncoderFactory) NewEncoder(responseWriter http.ResponseWriter, resourceType string) Encoder {

	// set content type
	responseWriter.Header().Set("Content-Type", "application/json")

	return &jsonEncoder{
		jsonEncoder:    json.NewEncoder(responseWriter),
		responseWriter: responseWriter,
		resourceType:   resourceType,
	}
}
