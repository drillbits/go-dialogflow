//    Copyright 2017 drillbits
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package dialogflow

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Fulfillment represents data about text response(s), rich messages,
// response received from webhook.
type Fulfillment struct {
	Speech   string     `json:"speech"`
	Messages []*Message `json:"messages"`
}

// FulfillmentRequest represents an HTTP request to be send by Dialogflow.
// See https://dialogflow.com/docs/fulfillment#request.
type FulfillmentRequest struct {
	*QueryResponse
	// TODO: originalRequest
}

// ParseRequest parses the given request and returns FulfillmentRequest,.
func ParseRequest(r *http.Request) (*FulfillmentRequest, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var req FulfillmentRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		return nil, err
	}
	return &req, nil
}
