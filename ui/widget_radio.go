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
	"github.com/go-gl/mathgl/mgl32"
	"github.com/haakenlabs/forge"
)

type RadioState int

const (
	RadioStateOff RadioState = iota
	RadioStateMixed
	RadioStateOn
)

type RadioGroup struct {
	BaseComponent

	radios []*Radio
}

var defaultRadioSize = mgl32.Vec2{16, 16}
var defaultRadioCheckSize = mgl32.Vec2{9, 9}

var _ Widget = &Radio{}

type Radio struct {
	BaseComponent

	state      RadioState
	eventState EventType

	BgColor       forge.Color
	BgColorActive forge.Color
	RadioMixColor forge.Color
	RadioOnColor  forge.Color

	backgroundColor forge.Color
	tint            forge.Color

	onChangeFunc func(RadioState)

	background *Graphic
	check      *Graphic
	text       *Text
}

func (w *Radio) Dragging() bool {
	return false
}

func (w *Radio) HandleEvent(event EventType) {
	switch event {
	case EventClick:
		if w.state == RadioStateOn {
			w.state = RadioStateOff
		} else {
			w.state = RadioStateOn
		}
		if w.onChangeFunc != nil {
			w.onChangeFunc(w.state)
		}
	}

	w.eventState = event
}

func (w *Radio) Redraw() {
	switch w.eventState {
	case EventClick:
		fallthrough
	case EventMouseEnter:
		w.background.SetColor(w.BgColorActive)
	default:
		w.background.SetColor(w.BgColor)
	}
	switch w.state {
	case RadioStateMixed:
		w.check.SetColor(w.RadioMixColor)
	case RadioStateOn:
		w.check.SetColor(w.RadioOnColor)
	}

	m := w.RectTransform().ActiveMatrix()

	w.background.Draw(m)
	if w.state == RadioStateOn {
		w.check.Draw(m)
	}
	w.text.Draw(m)
}

func (w *Radio) Raycast(pos mgl32.Vec2) bool {
	bounding := forge.NewRect(
		w.RectTransform().WorldPosition().Add(w.background.Position()),
		w.background.Size(),
	)

	return bounding.Contains(pos)
}

func (w *Radio) Rearrange() {
	w.check.Refresh()
	w.check.SetPosition(Align(w.check.Rect(), w.background.Rect(), AlignmentMiddleCenter))
	w.background.Refresh()

	w.text.Refresh()
	textPos := Align(w.text.Rect(), w.background.Rect(), AlignmentMiddleLeft)
	textPos = textPos.Add(mgl32.Vec2{w.background.Size().X() + 8, 0})
	w.text.SetPosition(textPos)
}

func (w *Radio) SetValue(value string) {
	w.text.SetValue(value)
	w.Rearrange()
}

func (w *Radio) Start() {
	w.Rearrange()
}

func (w *RadioGroup) AddRadio(radio ...*Radio) {
	w.radios = append(w.radios, radio...)
}

func NewRadio() *Radio {
	w := &Radio{}

	w.SetName("UIRadio")
	forge.GetInstance().MustAssign(w)

	return w
}

func NewRadioGroup() *RadioGroup {
	w := &RadioGroup{}

	w.SetName("UIRadioGroup")
	forge.GetInstance().MustAssign(w)

	return w
}

func CreateRadio(name string) *forge.GameObject {
	object := CreateGenericObject(name)

	radio := NewRadio()

	radio.background = NewGraphic()
	radio.background.rect.SetSize(defaultRadioSize)
	radio.check = NewGraphic()
	radio.check.rect.SetSize(defaultRadioCheckSize)
	radio.text = NewText()
	radio.text.SetValue("Radio")

	object.AddComponent(radio)

	return object
}

func CreateRadioGroup(name string, radios ...*Radio) *forge.GameObject {
	object := CreateGenericObject(name)

	group := NewRadioGroup()
	group.AddRadio(radios...)

	object.AddComponent(group)

	return object
}
