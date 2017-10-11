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

import "time"

// QueryResponse represents an HTTP response of the query endpoint.
// https://dialogflow.com/docs/reference/agent/query#response.
type QueryResponse struct {
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Lang      string    `json:"lang"`
	Result    *Result   `json:"result"`
	Status    *Status   `json:"status"`
	SessionID string    `json:"sessionId"`
}

// Result represents an result of the query.
type Result struct {
	Source           string            `json:"source"`
	ResolvedQuery    string            `json:"resolvedQuery"`
	Speech           string            `json:"speech,omitempty"` // TODO: undocumented
	Action           string            `json:"action"`
	ActionIncomplete bool              `json:"actionIncomplete"`
	Parameters       map[string]string `json:"parameters"`
	Contexts         []*Context        `json:"contexts,omitempty"`
	Fulfillment      *Fulfillment      `json:"fulfillment"`
	Score            float64           `json:"score"`
	Metadata         *Metadata         `json:"metadata"`
}

// Message represents a text response.
type Message struct {
	Type   int    `json:"type"`
	ID     string `json:"id,omitempty"` // TODO: undocumented
	Speech string `json:"speech"`
}

// Metadata contains data on intents and contexts.
type Metadata struct {
	IntentID                  string `json:"intentId"`
	WebhookUsed               string `json:"webhookUsed"`
	WebhookForSlotFillingUsed string `json:"webhookForSlotFillingUsed"`
	WebhookResponseTime       int    `json:"webhookResponseTime"`
	IntentName                string `json:"intentName"`
}

// Status represents how the request succeeded or failed.
// See https://docs.api.ai/docs/status-object
type Status struct {
	Code         int    `json:"code"`
	ErrorType    string `json:"errorType"`
	ErrorID      string `json:"errorId,omitempty"`
	ErrorDetails string `json:"errorDetails,omitempty"`
}

// Context represents a current context of a user’s request.
type Context struct {
	Name       string            `json:"name,omitempty"`
	Parameters map[string]string `json:"parameters,omitempty"`
	Lifespan   int               `json:"lifespan,omitempty"`
}
