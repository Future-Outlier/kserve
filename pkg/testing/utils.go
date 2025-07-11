/*
Copyright 2025 The KServe Authors.

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

package testing

import (
	"fmt"
	"os"
	"path/filepath"
)

// ProjectRoot returns the root directory of the project by searching for the go.mod file up from where it is called.
func ProjectRoot() string {
	rootDir := ""

	currentDir, errCurrDir := os.Getwd()
	if errCurrDir != nil {
		panic(fmt.Sprintf("failed to get current working directory: %v", errCurrDir))
	}

	for {
		if _, err := os.Stat(filepath.Join(currentDir, "go.mod")); err == nil {
			rootDir = filepath.FromSlash(currentDir)
			break
		}

		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			break
		}

		currentDir = parentDir
	}

	if rootDir == "" {
		panic("failed to get project root directory. no go.mod file found")
	}

	return rootDir
}
