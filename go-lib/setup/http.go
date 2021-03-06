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

import "net/http"

// HTTPServer starts an HTTP server
func HTTPServer(name, port string, router http.Handler) {
	logger.Info(name+" listening at", "port", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		logger.Error("Can't initiate HTTP service", err)
	}
}
