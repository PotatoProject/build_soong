// Copyright 2023 Google Inc. All rights reserved.
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

package rust

import (
	"android/soong/android"
	"fmt"
	"strings"
	"testing"
)

func TestAfdoEnabled(t *testing.T) {
	bp := `
	rust_binary {
		name: "foo",
		srcs: ["foo.rs"],
		afdo: true,
	}
`
	result := android.GroupFixturePreparers(
		prepareForRustTest,
		android.FixtureAddTextFile("toolchain/pgo-profiles/sampling/foo.afdo", ""),
		rustMockedFiles.AddToFixture(),
	).RunTestWithBp(t, bp)

	foo := result.ModuleForTests("foo", "android_arm64_armv8-a").Rule("rustc")

	expectedCFlag := fmt.Sprintf(afdoFlagFormat, "toolchain/pgo-profiles/sampling/foo.afdo")

	if !strings.Contains(foo.Args["rustcFlags"], expectedCFlag) {
		t.Errorf("Expected 'foo' to enable afdo, but did not find %q in cflags %q", expectedCFlag, foo.Args["rustcFlags"])
	}
}

func TestAfdoEnabledWithMultiArchs(t *testing.T) {
	bp := `
	rust_binary {
		name: "foo",
		srcs: ["foo.rs"],
		afdo: true,
		compile_multilib: "both",
	}
`
	result := android.GroupFixturePreparers(
		prepareForRustTest,
		android.FixtureAddTextFile("toolchain/pgo-profiles/sampling/foo_arm.afdo", ""),
		android.FixtureAddTextFile("toolchain/pgo-profiles/sampling/foo_arm64.afdo", ""),
		rustMockedFiles.AddToFixture(),
	).RunTestWithBp(t, bp)

	fooArm := result.ModuleForTests("foo", "android_arm_armv7-a-neon").Rule("rustc")
	fooArm64 := result.ModuleForTests("foo", "android_arm64_armv8-a").Rule("rustc")

	expectedCFlagArm := fmt.Sprintf(afdoFlagFormat, "toolchain/pgo-profiles/sampling/foo_arm.afdo")
	expectedCFlagArm64 := fmt.Sprintf(afdoFlagFormat, "toolchain/pgo-profiles/sampling/foo_arm64.afdo")

	if !strings.Contains(fooArm.Args["rustcFlags"], expectedCFlagArm) {
		t.Errorf("Expected 'fooArm' to enable afdo, but did not find %q in cflags %q", expectedCFlagArm, fooArm.Args["rustcFlags"])
	}

	if !strings.Contains(fooArm64.Args["rustcFlags"], expectedCFlagArm64) {
		t.Errorf("Expected 'fooArm64' to enable afdo, but did not find %q in cflags %q", expectedCFlagArm64, fooArm64.Args["rustcFlags"])
	}
}
