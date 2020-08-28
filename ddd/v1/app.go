// Copyright 2020 Torben Schinke
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ddd

// AppSpec contains the information about the entire application and all its bounded contexts.
type AppSpec struct {
	name            string
	description     string
	boundedContexts []*BoundedContextSpec
	pos             Pos
}

// Application is a factory to create an AppSpec.
func Application(name, description string, domains []*BoundedContextSpec) *AppSpec {
	return &AppSpec{
		pos:             capturePos("Application", 1),
		name:            name,
		boundedContexts: domains,
		description:     description,
	}
}

// Name returns the applications name
func (a *AppSpec) Name() string {
	return a.name
}

// BoundedContexts returns the splitted domains contexts.
func (a *AppSpec) BoundedContexts() []*BoundedContextSpec {
	return a.boundedContexts
}

// Pos returns debug position.
func (a *AppSpec) Pos() Pos {
	return a.pos
}

// Description tells a story about the application.
func (a *AppSpec) Description() string {
	return a.description
}

func (a *AppSpec) Walk(f func(obj interface{}) error) error {
	if err := f(a); err != nil {
		return err
	}

	for _, context := range a.boundedContexts {
		if err := context.Walk(f); err != nil {
			return err
		}
	}

	return nil
}
