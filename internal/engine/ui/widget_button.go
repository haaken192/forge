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
	"github.com/haakenlabs/forge/internal/engine"
)

var _ Renderer = &Button{}

var defaultButtonSize = mgl32.Vec2{96, 32}

type Button struct {
	BaseComponent
	Appearance

	value string

	onPressedFunc func()

	background *Graphic
	text       *Text
}

func NewButton() *Button {
	w := &Button{
		value:      "Button",
		background: NewGraphic(),
		text:       NewText(),
	}

	w.TextColor = Styles.PrimaryTextColor
	w.TextColorActive = Styles.PrimaryTextColor
	w.TextColorFocus = Styles.PrimaryTextColor
	w.BgColor = Styles.BackgroundColor
	w.BgColorActive = Styles.BackgroundColor
	w.BgColorFocus = Styles.BackgroundColor

	w.SetName("UIButton")
	engine.GetInstance().MustAssign(w)

	w.background.SetColor(w.BgColor)
	w.background.rect.SetSize(defaultButtonSize)

	w.text.SetFontSize(12)
	w.text.SetValue(w.value)
	w.text.SetColor(w.TextColor)

	return w
}

func (w *Button) SetValue(value string) {
	w.value = value
	w.text.SetValue(value)
}

func (w *Button) Value() string {
	return w.value
}

func (w *Button) SetOnPressedFunc(fn func()) {
	w.onPressedFunc = fn
}

func (w *Button) OnMouseEnter() {
	w.background.SetColor(w.BgColorActive)
	w.text.SetColor(w.TextColorActive)
}

func (w *Button) OnMouseLeave() {
	w.background.SetColor(w.BgColor)
	w.text.SetColor(w.TextColor)
}

func (w *Button) OnClick() {
	if w.onPressedFunc != nil {
		w.onPressedFunc()
	}
}

func (w *Button) UIDraw() {
	m := w.GetTransform().ActiveMatrix()

	w.background.Draw(m)
	w.text.Draw(m)
}

func (w *Button) Start() {
	w.Rearrange()
}

func (w *Button) Rearrange() {
	w.text.Refresh()
	textPos := Align(w.text.Rect(), w.RectTransform().Rect(), AlignmentMiddleCenter)
	w.text.SetPosition(textPos)

	w.background.Refresh()
}

func ButtonComponent(g *engine.GameObject) *Button {
	c := g.Components()
	for i := range c {
		if ct, ok := c[i].(*Button); ok {
			return ct
		}
	}

	return nil
}

func CreateButton(name string) *engine.GameObject {
	object := CreateGenericObject(name)
	rt := RectTransformComponent(object)
	rt.SetSize(defaultButtonSize)

	button := NewButton()

	object.AddComponent(button)

	return object
}
