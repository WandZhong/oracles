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

package elastic

import (
	"time"

	"github.com/ethereum/go-ethereum/log"
)

var logger = log.Root()

/*
docMsg represents the message to be sent to ElasticSearch.
- Type: a type name that corresponds the structure of the payload (Data field)
- Timestamp: the time of the sending of the message
- Source: A description of the source
- UserContext: user security context
- Data: the data specific (payload). The data structure is determined by the field Type.
- Tags: a set of tags
*/
type docMsg struct {
	Type        string      `json:"type"`
	Timestamp   time.Time   `json:"msgTimeStamp"`
	Source      *feedInfo   `json:"source"`
	UserContext interface{} `json:"userContext,omitempty"`
	Data        interface{} `json:"data"`
	Tags        []string    `json:"tags,omitempty"`
}

// feedInfo is the information to identify the source in ElasticSearch
// - Name: is an identifier of the source.
// - Extra: allows any free form of information as a complement to the source information
type feedInfo struct {
	Name  string      `json:"name"`
	Extra interface{} `json:"extra,omitempty"`
}
