// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package validate

import (
	"fmt"
	"regexp"
)

func ApplicationTemplateName(v interface{}, k string) (warnings []string, errors []error) {
	value := v.(string)

	if matched := regexp.MustCompile(`^.{1,50}$`).Match([]byte(value)); !matched {
		errors = append(errors, fmt.Errorf("test: %s, %q length should between 1~50", k, v))
	}
	return warnings, errors
}
