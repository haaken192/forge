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

type CheckState int

const (
	CheckStateOff CheckState = iota
	CheckStateMixed
	CheckStateOn
)

var _ Widget = &Checkbox{}

var defaultCheckboxSize = mgl32.Vec2{16, 16}
var defaultCheckboxCheckSize = mgl32.Vec2{9, 9}

type CheckboxGroup struct {
	BaseComponent

	checkboxes []*Checkbox
}

type Checkbox struct {
	BaseComponent

	state      CheckState
	eventState EventType

	BgColor       forge.Color
	BgColorActive forge.Color
	CheckMixColor forge.Color
	CheckOnColor  forge.Color

	onChangeFunc func(CheckState)

	background *Graphic
	check      *Graphic
	text       *Text
}

func (w *Checkbox) SetOnChangeFunc(fn func(CheckState)) {
	w.onChangeFunc = fn
}

func (w *Checkbox) Dragging() bool {
	return false
}

func (w *Checkbox) HandleEvent(event EventType) {
	switch event {
	case EventClick:
		if w.state == CheckStateOn {
			w.state = CheckStateOff
		} else {
			w.state = CheckStateOn
		}
		if w.onChangeFunc != nil {
			w.onChangeFunc(w.state)
		}
	}

	w.eventState = event
}

func (w *Checkbox) Redraw() {
	switch w.eventState {
	case EventClick:
		fallthrough
	case EventMouseEnter:
		w.background.SetColor(w.BgColorActive)
	default:
		w.background.SetColor(w.BgColor)
	}
	switch w.state {
	case CheckStateMixed:
		w.check.SetColor(w.CheckMixColor)
	case CheckStateOn:
		w.check.SetColor(w.CheckOnColor)
	}

	m := w.RectTransform().ActiveMatrix()

	w.background.Draw(m)
	if w.state == CheckStateOn {
		w.check.Draw(m)
	}
	w.text.Draw(m)
}

func (w *Checkbox) Raycast(pos mgl32.Vec2) bool {
	bounding := forge.NewRect(
		w.RectTransform().WorldPosition().Add(w.background.Position()),
		w.background.Size(),
	)

	return bounding.Contains(pos)
}

func (w *Checkbox) Start() {
	w.Rearrange()
}

func (w *Checkbox) Rearrange() {
	w.check.Refresh()
	w.check.SetPosition(Align(w.check.Rect(), w.background.Rect(), AlignmentMiddleCenter))
	w.background.Refresh()

	w.text.Refresh()
	textPos := Align(w.text.Rect(), w.background.Rect(), AlignmentMiddleLeft)
	textPos = textPos.Add(mgl32.Vec2{w.background.Size().X() + 8, 0})
	w.text.SetPosition(textPos)
}

func (w *Checkbox) SetValue(value string) {
	w.text.SetValue(value)
	w.Rearrange()
}

func (w *CheckboxGroup) AddCheckbox(checkbox ...*Checkbox) {
	w.checkboxes = append(w.checkboxes, checkbox...)
}

func NewCheckbox() *Checkbox {
	w := &Checkbox{}

	w.BgColor = Styles.WidgetColor
	w.BgColorActive = Styles.WidgetColorActive
	w.CheckMixColor = Styles.WidgetColor
	w.CheckOnColor = Styles.WidgetColorPrimary

	w.SetName("UICheckbox")
	forge.GetInstance().MustAssign(w)

	return w
}

func NewCheckboxGroup() *CheckboxGroup {
	w := &CheckboxGroup{}

	w.SetName("UICheckboxGroup")
	forge.GetInstance().MustAssign(w)

	return w
}

func CheckboxComponent(g *forge.GameObject) *Checkbox {
	c := g.Components()
	for i := range c {
		if ct, ok := c[i].(*Checkbox); ok {
			return ct
		}
	}

	return nil
}

func CreateCheckbox(name string) *forge.GameObject {
	object := CreateGenericObject(name)

	checkbox := NewCheckbox()

	checkbox.background = NewGraphic()
	checkbox.background.rect.SetSize(defaultCheckboxSize)
	checkbox.check = NewGraphic()
	checkbox.check.rect.SetSize(defaultCheckboxCheckSize)
	checkbox.text = NewText()
	checkbox.text.SetValue("Checkbox")

	object.AddComponent(checkbox)

	return object
}

func CreateCheckboxGroup(name string, checkboxes ...*Checkbox) *forge.GameObject {
	object := CreateGenericObject(name)

	group := NewCheckboxGroup()
	group.AddCheckbox(checkboxes...)

	object.AddComponent(group)

	return object
}
