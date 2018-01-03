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

// Package elastic provides a way to send messages to an ElasticSearch instance.
// ES uses the following concepts:
// - index: the name of the index where ElasticSearch
// - DocType: the type of Doc in the index. (Will have to be dropped in ES7
// - The ID of the message
package elastic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"time"

	"github.com/robert-zaremba/errstack"
)

// Feed is the client to send message to ElsasticSearch
type Feed struct {
	client      *http.Client
	host        string
	feedInfo    *feedInfo
	userContext interface{}
}

// NewFeed is the Factory
// -elasticAddress: the url of the instance of ElasticSearch
// -srcName: the name the program feeding ElasticSearch
func NewFeed(elasticAddress string, srcName string) *Feed {
	return &Feed{
		client: &http.Client{},
		host:   elasticAddress,
		feedInfo: &feedInfo{
			Name: srcName,
		},
	}
}

// Send an upsert message to ElasticSearch
// - index: the name of the index in ElasticSearch
// - docType: the doc type in elastic search
// - msgType: the message type that defines the structure of the msgBody
// - msgID: an identifier of the message for ES. No ID, means a system insert. If an ID provided, an Upsert
// - msgBody: the msg payload
// Returns an error matching the outcome of the query
func (r *Feed) Send(index string, docType string, msgType string, msgID string, msgBody interface{}) errstack.E {
	msg := &docMsg{
		Timestamp:   time.Now(),
		Source:      r.feedInfo,
		Type:        msgType,
		Data:        msgBody,
		UserContext: r.userContext,
	}

	buf, err := json.MarshalIndent(msg, "", "\t")
	if err != nil {
		return errstack.WrapAsDomain(err, "can't marshal Feed message")
	}
	endPoint := path.Join(index, docType)
	logger.Debug(string(buf))

	var url, method string
	if msgID == "" {
		url = fmt.Sprintf("%s/%s", r.host, endPoint)
		method = "POST"
	} else {
		url = fmt.Sprintf("%s/%s", r.host, path.Join(endPoint, msgID))
		method = "PUT"
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(buf))
	if err != nil {
		return errstack.WrapAsDomain(err, "Error building request")
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return errstack.WrapAsDomain(err, "Error sending request")
	}
	defer errstack.CallAndLog(logger, resp.Body.Close)

	// StatusCode other than 2XX is an error
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return errstack.NewDomainF("Sending Feed message returned unexpected status code: %d",
			resp.Status)
	}
	return nil
}
