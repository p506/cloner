/*
Copyright © 2019 Marcello de Sales <marcello.desales@gmail.com>

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
package cmd

import (
	"testing"
)

/**
 * Testing Golang CLI
 * https://stackoverflow.com/questions/59709345/how-to-implement-unit-tests-for-cli-commands-in-go/59714127#59714127
 */
func TestGitCloneWithWrongURLs(t *testing.T) {

	wrongUrlTests := []CliTests{
		{
			title:              "Given an empty URL to clone",
			input:              TestInput{url: "", forceClone: false},
			exitCode:           1,
			exactErrorMessages: map[string]bool{"git URL invalid: you must provide the repo URL": true},
		},
		{
			title:                "Given an invalid git URL to clone",
			input:                TestInput{url: "incorrect-url", forceClone: false},
			exitCode:             1,
			errorMessageContains: []string{"git URL invalid:", "did not find any match", "incorrect-url"},
		},
	}

	// Execute the test cases
	ExecuteCliTestsStrategy(t, wrongUrlTests)
}

func TestGitCloneSuccessfullyForcing(t *testing.T) {
	existingDirTests := []CliTests{
		{
			title: "Given existing cloned repo",
			input: TestInput{
				url:        "https://github.com/comsysto/redis-locks-with-grafana",
				forceClone: true},
			exitCode: 0,
		},
	}

	// Execute the test cases
	ExecuteCliTestsStrategy(t, existingDirTests)
}

func TestGitCloneExistingDirs(t *testing.T) {

	existingDirTests := []CliTests{
		{
			title:                "Given existing cloned repo",
			input:                TestInput{url: "https://github.com/comsysto/redis-locks-with-grafana", forceClone: false},
			exitCode:             2,
			errorMessageContains: []string{"can't clone repo: clone location", "exists and it's not empty"},
		},
	}

	// Execute the test cases
	ExecuteCliTestsStrategy(t, existingDirTests)
}
