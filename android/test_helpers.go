// Copyright 2021 Google Inc. All rights reserved.
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

package android

import (
	"reflect"
	"strings"
	"testing"
)

// Provides general test support.
type TestHelper struct {
	*testing.T
}

// AssertBoolEquals checks if the expected and actual values are equal and if they are not then it
// reports an error prefixed with the supplied message and including a reason for why it failed.
func (h *TestHelper) AssertBoolEquals(message string, expected bool, actual bool) {
	h.Helper()
	if actual != expected {
		h.Errorf("%s: expected %t, actual %t", message, expected, actual)
	}
}

// AssertStringEquals checks if the expected and actual values are equal and if they are not then
// it reports an error prefixed with the supplied message and including a reason for why it failed.
func (h *TestHelper) AssertStringEquals(message string, expected string, actual string) {
	h.Helper()
	if actual != expected {
		h.Errorf("%s: expected %s, actual %s", message, expected, actual)
	}
}

// AssertErrorMessageEquals checks if the error is not nil and has the expected message. If it does
// not then this reports an error prefixed with the supplied message and including a reason for why
// it failed.
func (h *TestHelper) AssertErrorMessageEquals(message string, expected string, actual error) {
	h.Helper()
	if actual == nil {
		h.Errorf("Expected error but was nil")
	} else if actual.Error() != expected {
		h.Errorf("%s: expected %s, actual %s", message, expected, actual.Error())
	}
}

// AssertTrimmedStringEquals checks if the expected and actual values are the same after trimming
// leading and trailing spaces from them both. If they are not then it reports an error prefixed
// with the supplied message and including a reason for why it failed.
func (h *TestHelper) AssertTrimmedStringEquals(message string, expected string, actual string) {
	h.Helper()
	h.AssertStringEquals(message, strings.TrimSpace(expected), strings.TrimSpace(actual))
}

// AssertStringDoesContain checks if the string contains the expected substring. If it does not
// then it reports an error prefixed with the supplied message and including a reason for why it
// failed.
func (h *TestHelper) AssertStringDoesContain(message string, s string, expectedSubstring string) {
	h.Helper()
	if !strings.Contains(s, expectedSubstring) {
		h.Errorf("%s: could not find %q within %q", message, expectedSubstring, s)
	}
}

// AssertStringDoesNotContain checks if the string contains the expected substring. If it does then
// it reports an error prefixed with the supplied message and including a reason for why it failed.
func (h *TestHelper) AssertStringDoesNotContain(message string, s string, unexpectedSubstring string) {
	h.Helper()
	if strings.Contains(s, unexpectedSubstring) {
		h.Errorf("%s: unexpectedly found %q within %q", message, unexpectedSubstring, s)
	}
}

// AssertStringListContains checks if the list of strings contains the expected string. If it does
// not then it reports an error prefixed with the supplied message and including a reason for why it
// failed.
func (h *TestHelper) AssertStringListContains(message string, list []string, expected string) {
	h.Helper()
	if !InList(expected, list) {
		h.Errorf("%s: could not find %q within %q", message, expected, list)
	}
}

// AssertArrayString checks if the expected and actual values are equal and if they are not then it
// reports an error prefixed with the supplied message and including a reason for why it failed.
func (h *TestHelper) AssertArrayString(message string, expected, actual []string) {
	h.Helper()
	if len(actual) != len(expected) {
		h.Errorf("%s: expected %d (%q), actual (%d) %q", message, len(expected), expected, len(actual), actual)
		return
	}
	for i := range actual {
		if actual[i] != expected[i] {
			h.Errorf("%s: expected %d-th, %q (%q), actual %q (%q)",
				message, i, expected[i], expected, actual[i], actual)
			return
		}
	}
}

// AssertDeepEquals checks if the expected and actual values are equal using reflect.DeepEqual and
// if they are not then it reports an error prefixed with the supplied message and including a
// reason for why it failed.
func (h *TestHelper) AssertDeepEquals(message string, expected interface{}, actual interface{}) {
	h.Helper()
	if !reflect.DeepEqual(actual, expected) {
		h.Errorf("%s: expected:\n  %#v\n got:\n  %#v", message, expected, actual)
	}
}

// AssertPanic checks that the supplied function panics as expected.
func (h *TestHelper) AssertPanic(message string, funcThatShouldPanic func()) {
	h.Helper()
	panicked := false
	func() {
		defer func() {
			if x := recover(); x != nil {
				panicked = true
			}
		}()
		funcThatShouldPanic()
	}()
	if !panicked {
		h.Error(message)
	}
}
