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

	"github.com/haakenlabs/forge/internal/engine"
)

type StyleSet struct {
	BackgroundColor    engine.Color `json:"background_color"`
	AltBackgroundColor engine.Color `json:"alt_background_color"`
	PrimaryColor       engine.Color `json:"primary_color"`
	PrimaryTextColor   engine.Color `json:"primary_text_color"`
	SecondaryTextColor engine.Color `json:"secondary_text_color"`
	TertiaryTextColor  engine.Color `json:"tertiary_text_color"`
	InverseTextColor   engine.Color `json:"inverse_text_color"`
}

var Styles = StyleSet{
	BackgroundColor:    engine.Color{0.1, 0.1, 0.1, 0.9},
	AltBackgroundColor: engine.Color{0.24, 0.24, 0.24, 0.9},
	PrimaryColor:       engine.Color{0.0, 0.27, 0.68, 0.9},
	PrimaryTextColor:   engine.ColorWhite,
	SecondaryTextColor: engine.ColorYellow,
	TertiaryTextColor:  engine.ColorGreen,
	InverseTextColor:   engine.ColorBlue,
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
