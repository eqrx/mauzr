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
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	http.Client
}

func NewClient(crtPath, keyPath string) Client {
	config := TLSConfig(crtPath, keyPath)
	client := Client{
		http.Client{
			Transport: &http.Transport{TLSClientConfig: config},
		},
	}
	return client
}

func (c *Client) GetWithContext(ctx context.Context, url string, body interface{}) error {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err == nil {
		var response *http.Response
		response, err = c.Do(request)

		switch {
		case err != nil:
		case response.StatusCode != http.StatusOK:
			err = fmt.Errorf("%v", response.Status)
		default:
			defer response.Body.Close()
			if body != nil {
				if err = json.NewDecoder(response.Body).Decode(&body); err != nil {
					err = fmt.Errorf("could not deserialize JSON - %v", err)
				}
			}
		}
	}
	return err
}

func (c *Client) PostWithContext(ctx context.Context, url string, body io.Reader) error {
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err == nil {
		var response *http.Response
		response, err = c.Do(request)
		switch {
		case err != nil:
		case response.StatusCode != http.StatusOK:
			err = fmt.Errorf("%v", response.Status)
		default:
			response.Body.Close()
		}
	}
	return err
}