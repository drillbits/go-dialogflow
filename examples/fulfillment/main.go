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

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/drillbits/go-dialogflow/dialogflow"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello")
	})
	http.HandleFunc("/fulfillment", func(w http.ResponseWriter, r *http.Request) {
		req, err := dialogflow.ParseRequest(r)
		if err != nil {
			log.Printf("failed to parse: %s", err)
			return
		}
		b, err := json.MarshalIndent(req, "", "  ")
		log.Print(string(b))
	})

	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("http server started on %s", l.Addr())
	err = http.Serve(l, nil)
	if err != nil {
		log.Fatal(err)
	}
}
