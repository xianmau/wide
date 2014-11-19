// Copyright (c) 2014, B3log
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

package util

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// WebSocket channel.
type WSChannel struct {
	Sid     string          // wide session id
	Conn    *websocket.Conn // websocket connection
	Request *http.Request   // HTTP request related
	Time    time.Time       // the latest use time
}

// Close closed the channel.
func (c *WSChannel) Close() {
	if nil != c.Conn {
		c.Conn.Close()
	}
}

// Refresh refreshes the channel by updating its use time.
func (c *WSChannel) Refresh() {
	c.Time = time.Now()
}
