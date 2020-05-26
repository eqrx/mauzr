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

// Package color contains color types.
package color

var (
	// Off turns the pixel off.
	Off = RGBW{}
	// Bright sets the pixel to as bright as possible.
	Bright = RGBW{White: 1.0}
	// Unmanaged indicates that the pixel is not actively managed.
	Unmanaged = RGBW{Red: 1.0}
	// Error indicates that something is wrong.
	Error = Off.MixWith(0.1, RGBW{Red: 1.0})
	// Warning indicates that something requires attention.
	Warning = Off.MixWith(0.1, RGBW{Red: 1.0, Green: 1.0})
	// Good indicated that everything is fine.
	Good = Off.MixWith(0.1, RGBW{Green: 1.0})
)
