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

type Progress struct {
	BaseComponent

	progress float64

	backgroundColor engine.Color
	tint            engine.Color

	onChangeFunc func(float64)

	background  *Graphic
	activeTrack *Graphic
}

func (w *Progress) UIDraw() {
	m := w.RectTransform().ActiveMatrix()

	w.background.Draw(m)
	w.activeTrack.Draw(m)
}

func NewProgress() *Progress {
	w := &Progress{
		progress: 0.0,
	}

	w.SetName("UIProgress")
	engine.GetInstance().MustAssign(w)

	return w
}

func ProgressComponent(g *engine.GameObject) *Progress {
	c := g.Components()
	for i := range c {
		if ct, ok := c[i].(*Progress); ok {
			return ct
		}
	}

	return nil
}

func CreateProgress(name string) *engine.GameObject {
	object := CreateGenericObject(name)

	progress := NewProgress()

	progress.background = NewGraphic()
	progress.activeTrack = NewGraphic()

	object.AddComponent(progress)

	return object
}