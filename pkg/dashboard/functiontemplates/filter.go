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

import "strings"

type Filter struct {
	Contains string
}

func (ftf *Filter) functionTemplatePasses(template *FunctionTemplate) bool {
	if ftf.empty() {
		return true
	}

	stringsToSearch := []string{
		template.Name,
	}

	if string(template.serializedTemplate) != "" {
		stringsToSearch = append(stringsToSearch, string(template.serializedTemplate))
	}

	for _, stringToSearch := range stringsToSearch {
		if strings.Contains(stringToSearch, ftf.Contains) {
			return true
		}
	}

	return false
}

func (ftf *Filter) empty() bool {
	return ftf.Contains == ""
}
