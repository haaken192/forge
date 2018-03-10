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

import "github.com/haakenlabs/forge/internal/engine"

type Label struct {
	BaseComponent
	Appearance

	text *Text
}

func NewLabel() *Label {
	w := &Label{
		text: NewText(),
	}

	w.TextColor = Styles.PrimaryTextColor

	w.SetName("UILabel")
	engine.GetInstance().MustAssign(w)

	w.text.SetValue("Label")
	w.text.SetFontSize(12)
	w.text.SetColor(w.TextColor)

	return w
}

func (w *Label) UIDraw() {
	m := w.RectTransform().ActiveMatrix()

	w.text.Draw(m)
}

func (w *Label) SetValue(value string) {
	w.text.SetValue(value)
	w.Rearrange()
}

func (w *Label) Value() string {
	return w.text.Value()
}

func (w *Label) SetFontSize(size int32) {
	w.text.SetFontSize(size)
	w.Rearrange()
}

func (w *Label) FontSize(size int32) int32 {
	return w.text.fontSize
}

func (w *Label) Rearrange() {
	w.text.Refresh()
}

func (w *Label) OnActivate() {
	w.Rearrange()
}

func (w *Label) OnTransformChanged() {
	w.Rearrange()
}

func (w *Label) Start() {
	w.Rearrange()
}

func LabelComponent(g *engine.GameObject) *Label {
	c := g.Components()
	for i := range c {
		if ct, ok := c[i].(*Label); ok {
			return ct
		}
	}

	return nil
}

func CreateLabel(name string) *engine.GameObject {
	object := CreateGenericObject(name)

	object.AddComponent(NewLabel())

	return object
}
