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
	"os"
	"runtime/pprof"
	"time"

	"github.com/robert-zaremba/errstack"
)

// CPUProfile dumps a CPU profile every given `minutes`
func CPUProfile(filename string, minutes int) {
	if minutes <= 0 {
		return
	}
	// logger.Info(fmt.Sprintf("Writing cpu profile every %d minutes to %s", minutes, filename))
	for {
		f, err := os.Create(filename)
		if err != nil {
			logger.Error("Can't create a file for CPU profile dump", err)
		}
		errstack.Log(logger, pprof.StartCPUProfile(f))
		time.Sleep(time.Minute * time.Duration(minutes))
		pprof.StopCPUProfile()
		errstack.Log(logger, f.Close())
	}
}

// MemProfile dumps a memory profile every given `minutes`
func MemProfile(filename string, minutes int) {
	if minutes <= 0 {
		return
	}
	// logger.Info(fmt.Sprintf("Writing memory profile every %d minutes to %s", minutes, filename))
	for range time.NewTicker(time.Minute * time.Duration(minutes)).C {
		f, err := os.Create(filename)
		if err != nil {
			logger.Fatal("Can't create a file for memory profile dump", err)
		}
		errstack.Log(logger, pprof.WriteHeapProfile(f))
		errstack.Log(logger, f.Close())
	}
}
