// Copyright 2023 syzkaller project authors. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

package linux

type customRules struct {
	// The mapping between a Linux subsystem name and its system calls.
	subsystemCalls map[string][]string
}

var (
	linuxSubsystemRules = &customRules{}
)