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
	"bytes"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func Test_ParseRequest(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    *FulfillmentRequest
		wantErr bool
	}{
		{
			name: "without originalRequest",
			args: args{
				r: httptest.NewRequest("POST", "/", bytes.NewBufferString(`{
					"id": "8d86775e-7128-423d-b166-a641a533b12b",
					"timestamp": "2017-10-11T02:22:05.208Z",
					"lang": "ja",
					"result": {
					  "source": "agent",
					  "resolvedQuery": "hi",
					  "speech": "",
					  "action": "input.unknown",
					  "actionIncomplete": false,
					  "parameters": {},
					  "contexts": [],
					  "metadata": {
						"intentId": "e2793c3e-7fe4-4a5e-ba55-b0bf18bb41b0",
						"webhookUsed": "true",
						"webhookForSlotFillingUsed": "false",
						"intentName": "Default Fallback Intent"
					  },
					  "fulfillment": {
						"speech": "Sorry, what was that?",
						"messages": [
						  {
							"type": 0,
							"id": "80c2c352-2f92-45f8-aaf3-aebc4779cedf",
							"speech": "Sorry, can you say that again?"
						  }
						]
					  },
					  "score": 1.0
					},
					"status": {
					  "code": 200,
					  "errorType": "success"
					},
					"sessionId": "c0438949-3918-4646-a079-a8e7c1644dfa"
				  }`)),
			},
			want: &FulfillmentRequest{
				QueryResponse: &QueryResponse{
					ID:        "8d86775e-7128-423d-b166-a641a533b12b",
					Timestamp: timeMustParse(time.RFC3339, "2017-10-11T02:22:05.208Z"),
					Lang:      "ja",
					Result: &Result{
						Source:           "agent",
						ResolvedQuery:    "hi",
						Speech:           "",
						Action:           "input.unknown",
						ActionIncomplete: false,
						Parameters:       make(map[string]string),
						Contexts:         make([]*Context, 0),
						Metadata: &Metadata{
							IntentID:                  "e2793c3e-7fe4-4a5e-ba55-b0bf18bb41b0",
							WebhookUsed:               "true",
							WebhookForSlotFillingUsed: "false",
							IntentName:                "Default Fallback Intent",
						},
						Fulfillment: &Fulfillment{
							Speech: "Sorry, what was that?",
							Messages: []*Message{
								&Message{
									Type:   0,
									ID:     "80c2c352-2f92-45f8-aaf3-aebc4779cedf",
									Speech: "Sorry, can you say that again?",
								},
							},
						},
						Score: 1.0,
					},
					Status: &Status{
						Code:      200,
						ErrorType: "success",
					},
					SessionID: "c0438949-3918-4646-a079-a8e7c1644dfa",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseRequest(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
