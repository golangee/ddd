// Code generated by golangee/eearc; DO NOT EDIT.
//
// Copyright 2021 Torben Schinke
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package supportietyserver

import (
	json "encoding/json"
	flag "flag"
	fmt "fmt"
	chat "github.com/golangee/architecture/testdata/workspace/server/internal/tickets/core/chat"
	usecase "github.com/golangee/architecture/testdata/workspace/server/internal/tickets/usecase"
	os "os"
)

// Configuration contains all aggregated configurations for the entire application and all contained bounded contexts.
type Configuration struct {
	Tickets TicketsConfig
}

// Reset restores this instance to the default state.
func (c *Configuration) Reset() {
	c.Tickets.Reset()
}

// ConfigureFlags configures the flags to be ready to get evaluated.
func (c *Configuration) ConfigureFlags(flags *flag.FlagSet) {
	c.Tickets.ConfigureFlags(flags)
}

// ParseEnv tries to parse the environment variables into this instance.
func (c *Configuration) ParseEnv() error {
	if err := c.Tickets.ParseEnv(); err != nil {
		return fmt.Errorf("cannot parse 'Tickets': %w", err)
	}

	return nil
}

// ParseFile tries to parse a json file into this instance. Only the defined values are overridden.
//
// Example JSON
//
//   {
//     "Tickets": {
//       "Core": {
//         "AnotherConfig": {
//           "BlaFeature": false
//         }
//       },
//       "Usecase": {
//         "MyConfig": {
//           "FancyFeature": false
//         }
//       }
//     }
//   }
func (c *Configuration) ParseFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("cannot open configuration file: %w", err)
	}

	defer file.Close() // intentionally ignoring read-only error on close

	dec := json.NewDecoder(file)
	if err := dec.Decode(c); err != nil {
		return fmt.Errorf("cannot decode json: %w", err)
	}

	return nil
}

// TicketsConfig contains all aggregated configurations for the entire bounded context 'tickets'.
type TicketsConfig struct {
	Core    TicketsCoreConfig
	Usecase TicketsUsecaseConfig
}

// Reset restores this instance to the default state.
func (t *TicketsConfig) Reset() {
	t.Core.Reset()
	t.Usecase.Reset()
}

// ConfigureFlags configures the flags to be ready to get evaluated.
func (t *TicketsConfig) ConfigureFlags(flags *flag.FlagSet) {
	t.Core.ConfigureFlags(flags)
	t.Usecase.ConfigureFlags(flags)
}

// ParseEnv tries to parse the environment variables into this instance.
func (t *TicketsConfig) ParseEnv() error {
	if err := t.Core.ParseEnv(); err != nil {
		return fmt.Errorf("cannot parse 'Core': %w", err)
	}
	if err := t.Usecase.ParseEnv(); err != nil {
		return fmt.Errorf("cannot parse 'Usecase': %w", err)
	}
	return nil

}

// TicketsCoreConfig contains all configurations for the 'core' layer of 'tickets'.
type TicketsCoreConfig struct {
	AnotherConfig chat.AnotherConfig
}

// Reset restores this instance to the default state.
func (t *TicketsCoreConfig) Reset() {
	t.AnotherConfig.Reset()
}

// ParseEnv tries to parse the environment variables into this instance.
func (t *TicketsCoreConfig) ParseEnv() error {
	if err := t.AnotherConfig.ParseEnv(); err != nil {
		return fmt.Errorf("cannot parse 'AnotherConfig': %w", err)
	}

	return nil
}

// ConfigureFlags configures the flags to be ready to get evaluated.
func (t *TicketsCoreConfig) ConfigureFlags(flags *flag.FlagSet) {
	t.AnotherConfig.ConfigureFlags(flags)
}

// TicketsUsecaseConfig contains all configurations for the 'usecase' layer of 'tickets'.
type TicketsUsecaseConfig struct {
	MyConfig usecase.MyConfig
}

// Reset restores this instance to the default state.
func (t *TicketsUsecaseConfig) Reset() {
	t.MyConfig.Reset()
}

// ParseEnv tries to parse the environment variables into this instance.
func (t *TicketsUsecaseConfig) ParseEnv() error {
	if err := t.MyConfig.ParseEnv(); err != nil {
		return fmt.Errorf("cannot parse 'MyConfig': %w", err)
	}

	return nil
}

// ConfigureFlags configures the flags to be ready to get evaluated.
func (t *TicketsUsecaseConfig) ConfigureFlags(flags *flag.FlagSet) {
	t.MyConfig.ConfigureFlags(flags)
}
