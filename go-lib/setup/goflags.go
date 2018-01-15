// Copyright (c) 2017 Sweetbridge Stiftung (Sweetbridge Foundation)
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
	"github.com/jessevdk/go-flags"
	"github.com/robert-zaremba/errstack"
	"os"
)

// ConfigFileOption is a struct that enables the usage of a configuration files to read options.
// This type is to be embedded into the the options structure.
// the file crawlerOpts.go gives an example
type ConfigFileOption struct {
	Config string `short:"c" long:"config" no-ini:"true" required:"true"`
}

// ConfigFilepath returns the path of the config file.
func (r *ConfigFileOption) ConfigFilepath() string {
	return r.Config
}

// ConfigFilePathGetter is an Interface that allows reading the config file path
type ConfigFilePathGetter interface {
	ConfigFilepath() string
}

// ReadOpts reads command line params using the go-flags library
func ReadOpts(opts interface{}) errstack.E {
	_, err := readOpts(opts)
	return err
}

/*
// ReadOptsUsingIniFile reads options from an ini file, using go-flags library
func ReadOptsUsingIniFile(opts ConfigFilePathGetter) errstack.E {
	p, err := readOpts(opts)
	if err != nil {
		return err
	}

	// Loads from the config file
	cfgFileName := opts.ConfigFilepath()
	if cfgFileName != "" {
		inip := flags.NewIniParser(p)
		inip.ParseAsDefaults = true
		if err := inip.ParseFile(cfgFileName); err != nil {
			return errstack.WrapAsReqF(err, "Failed to read the options from the file '%s'", cfgFileName)
		}
	}
	return nil
}
*/

// Utility function that initialises a parser, to load options from commmand line. It returns a parser, if no error.
// This function determines if a help option is present and prints the built-in help.
func readOpts(opts interface{}) (*flags.Parser, errstack.E) {
	p := flags.NewParser(opts, flags.Default|flags.IgnoreUnknown)
	if _, err := p.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		}
		return nil, errstack.NewReq(err.Error())
	}
	return p, nil
}
