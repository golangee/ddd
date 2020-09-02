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

// A UseCaseLayerSpec represents a stereotyped USECASE Layer.
type UseCaseLayerSpec struct {
	name        string
	description string
	cases       []*UseCaseSpec
}

// Walks loops over self and use cases.
func (u *UseCaseLayerSpec) Walk(f func(obj interface{}) error) error {
	if err := f(u); err != nil {
		return err
	}

	for _, obj := range u.cases {
		if err := obj.Walk(f); err != nil {
			return err
		}
	}

	return nil
}

// UseCases returns the contained use cases.
func (u *UseCaseLayerSpec) UseCases() []*UseCaseSpec {
	return u.cases
}

// Name of the layer.
func (u *UseCaseLayerSpec) Name() string {
	return u.name
}

// Description of the layer.
func (u *UseCaseLayerSpec) Description() string {
	return u.description
}

// Stereotype of the layer.
func (u *UseCaseLayerSpec) Stereotype() Stereotype {
	return USECASE
}

// UseCases is a factory for a UseCaseLayerSpec. A use case can only ever import Core API. Besides various instances
// of UseCase, it may also define package specific data models.
func UseCases(cases ...*UseCaseSpec) *UseCaseLayerSpec {
	return &UseCaseLayerSpec{
		name: "usecases",
		description: "Package usecase contains all domain specific use cases for the current bounded context.\n" +
			"It contains an exposed public API to be imported by PRESENTATION layers. \n" +
			"It provides a private implementation of the use cases accessible by factory functions.",
		cases: cases,
	}
}

// An UseCaseSpec contains different actors to accomplish a larger goal. This may be also called a use case.
type UseCaseSpec struct {
	// name of the epic or use case, e.g. "BookLoaning" or "BookSearching"
	name    string
	comment string
	stories []*UserStorySpec
}

// UseCase declares a use case consisting of user stories. The name must be a valid public identifier, because
// it will be used as the enclosing interface for all contained stories.
func UseCase(name, comment string, stories ...*UserStorySpec) *UseCaseSpec {
	return &UseCaseSpec{
		name:    name,
		comment: comment,
		stories: stories,
	}
}

// Name returns current name.
func (u *UseCaseSpec) Name() string {
	return u.name
}

// Comment of the use case.
func (u *UseCaseSpec) Comment() string {
	return u.comment
}

// Stories returns all contained user stories.
func (u *UseCaseSpec) Stories() []*UserStorySpec {
	return u.stories
}

// Walks loops over self and stories.
func (u *UseCaseSpec) Walk(f func(obj interface{}) error) error {
	if err := f(u); err != nil {
		return err
	}

	for _, obj := range u.stories {
		if err := obj.Walk(f); err != nil {
			return err
		}
	}

	return nil
}

// UserStorySpec contains a story and the representing calling endpoint and potential types, if not already defined
// globally.
type UserStorySpec struct {
	story   string
	fun     *FuncSpec
	structs []*StructSpec
	pos     Pos
}

// Story declares a user story and the according method and potential extra types.
// A story represents a scenario in a use case or epic and has the format
//  As a < role >, I want to < a goal > so that < a reason >.
// Make a good story and avoid the following pitfalls:
//  * unclear or subjective formulation, which refers to non-measurable values
//  * avoid eventualities and define the context of things (e.g. as a user I want to see >all or relevant< stuff...)
//  * do not use multiple main verbs for the subject (e.g. as a platform user I want to import and archive...)
//  * explain the reason why an actor wants/needs/must do something.
//
// See also Mike Cohns suggestions at https://www.mountaingoatsoftware.com/agile/user-stories and
// https://www.mountaingoatsoftware.com/blog/advantages-of-the-as-a-user-i-want-user-story-template
func Story(story string, fun *FuncSpec, structs ...*StructSpec) *UserStorySpec {
	return &UserStorySpec{
		story:   story,
		fun:     fun,
		structs: structs,
		pos:     capturePos("Story", 1),
	}
}

// Pos returns the debugging position
func (u *UserStorySpec) Pos() Pos {
	return u.pos
}

// Story returns the actual user story.
func (u *UserStorySpec) Story() string {
	return u.story
}

// Func represents the entry point.
func (u *UserStorySpec) Func() *FuncSpec {
	return u.fun
}

// Structs returns the specific models, but they are public to all other stories as well.
func (u *UserStorySpec) Structs() []*StructSpec {
	return u.structs
}

// Walks loops over self and funcs and structs.
func (u *UserStorySpec) Walk(f func(obj interface{}) error) error {
	if err := f(u); err != nil {
		return err
	}

	for _, obj := range u.structs {
		if err := obj.Walk(f); err != nil {
			return err
		}
	}

	return u.fun.Walk(f)
}
