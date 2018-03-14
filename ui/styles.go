/*
Copyright (c) 2018 HaakenLabs

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package ui

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/haakenlabs/forge"
)

type StyleSet struct {
	BackgroundColor    forge.Color `json:"background_color"`
	AltBackgroundColor forge.Color `json:"alt_background_color"`
	PrimaryColor       forge.Color `json:"primary_color"`
	PrimaryTextColor   forge.Color `json:"primary_text_color"`
	SecondaryTextColor forge.Color `json:"secondary_text_color"`
	TertiaryTextColor  forge.Color `json:"tertiary_text_color"`
	InverseTextColor   forge.Color `json:"inverse_text_color"`
}

var Styles = StyleSet{
	BackgroundColor:    forge.Color{0.1, 0.1, 0.1, 0.9},
	AltBackgroundColor: forge.Color{0.24, 0.24, 0.24, 0.9},
	PrimaryColor:       forge.Color{0.0, 0.27, 0.68, 0.9},
	PrimaryTextColor:   forge.ColorWhite,
	SecondaryTextColor: forge.ColorYellow,
	TertiaryTextColor:  forge.ColorGreen,
	InverseTextColor:   forge.ColorBlue,
}

func LoadStyle(r io.Reader) error {
	var s StyleSet

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	Styles = s

	return nil
}
