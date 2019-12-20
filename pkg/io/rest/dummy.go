/*
Copyright 2019 Alexander Sowitzki.

GNU Affero General Public License version 3 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://opensource.org/licenses/AGPL-3.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package rest

import (
	"context"
	"io"
	"net/http"
)

type Dummy struct {
}

func (d Dummy) Get(context.Context, string, interface{}) error                { return nil }
func (d Dummy) Post(context.Context, string, io.Reader) error                 { return nil }
func (d Dummy) Endpoint(path, form string, queryHandler func(query *Request)) {}
func (d Dummy) Serve(context.Context) error                                   { return nil }
func (d Dummy) HandleFunc(string, func(http.ResponseWriter, *http.Request))   {}
func (d Dummy) AddDefaultResponseHeader(http.Header)                          {}
func (d Dummy) Do(req *http.Request) (*http.Response, error)                  { return nil, nil }