// Copyright (c) 2017 Sweetbridge Inc.
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

package setup

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/kardianos/osext"
)

// Root returns an absolute path to the project root
// It panics if root can't be located
func Root() string {
	root, callerErr := findRootFromRuntimeCaller()
	if callerErr == nil {
		return root
	}
	root, executableErr := findRootFromExecutable()
	if executableErr == nil {
		return root
	}
	logger.Fatal("Can't locate project root", callerErr, executableErr)
	return root
}

func findRootFromRuntimeCaller() (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("Can't retrive runtime caller info")
	}
	root, ok := traverseUpForRoot(filename, 5)
	if !ok {
		return "", fmt.Errorf("Unable to find project root for runtime caller position %s", filename)
	}
	return root, nil
}

func findRootFromExecutable() (string, error) {
	filename, err := osext.Executable()
	if err != nil {
		return "", fmt.Errorf("Can't retrive runtime path to executable: %s", err.Error())
	}
	root, ok := traverseUpForRoot(filename, 5)
	if !ok {
		return "", fmt.Errorf("Unable to find project root for runtime executable position %s", filename)
	}
	return root, nil
}

func traverseUpForRoot(p string, levels int) (string, bool) {
	p = filepath.Dir(p)
	for i := 0; i < levels; i++ {
		if isRootPath(p) {
			return p, true
		}
		p = filepath.Dir(p)
	}
	return "", false
}

func isRootPath(p string) bool {
	return exists(filepath.Join(p, "LICENSE")) && exists(filepath.Join(p, ".git"))
}

func exists(name string) bool {
	_, err := os.Stat(name)
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}
	return err == nil || !os.IsNotExist(err)
}
