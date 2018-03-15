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
	"math"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/haakenlabs/forge"
)

const defaultProgressHeight = float32(12)

var _ Widget = &Progress{}

type Progress struct {
	BaseComponent

	progress float64

	WidgetColor       forge.Color
	WidgetColorActive forge.Color

	onChangeFunc func(float64)

	background  *Graphic
	activeTrack *Graphic
}

func (w *Progress) SetProgress(value float64) {
	w.progress = mgl64.Clamp(value, 0.0, 1.0)
	w.Rearrange()

	if w.onChangeFunc != nil {
		w.onChangeFunc(w.progress)
	}
}

func (w *Progress) Rearrange() {
	width := w.RectTransform().Size().X()
	//w.RectTransform().SetSize(mgl32.Vec2{width, defaultProgressHeight})

	activeWidth := float32(math.Floor(float64(width) * w.progress))
	bgWidth := float32(math.Ceil(float64(width) * (1.0 - w.progress)))

	w.activeTrack.SetSize(mgl32.Vec2{bgWidth, defaultProgressHeight})
	w.background.SetSize(mgl32.Vec2{activeWidth, defaultProgressHeight})
	w.background.SetPosition(mgl32.Vec2{bgWidth, 0})

	w.background.Refresh()
	w.activeTrack.Refresh()
}

func (w *Progress) Redraw() {
	m := w.RectTransform().ActiveMatrix()

	w.activeTrack.SetColor(w.WidgetColorActive)
	w.background.SetColor(w.WidgetColor)

	w.activeTrack.Draw(m)
	w.background.Draw(m)
}

func (w *Progress) Raycast(pos mgl32.Vec2) bool {
	return false
}

func (w *Progress) Dragging() bool {
	return false
}

func (w *Progress) HandleEvent(event EventType) {}

func NewProgress() *Progress {
	w := &Progress{
		progress: 0.0,
	}

	w.WidgetColor = Styles.WidgetColor
	w.WidgetColorActive = Styles.WidgetColorPrimary

	w.SetName("UIProgress")
	forge.GetInstance().MustAssign(w)

	return w
}

func ProgressComponent(g *forge.GameObject) *Progress {
	c := g.Components()
	for i := range c {
		if ct, ok := c[i].(*Progress); ok {
			return ct
		}
	}

	return nil
}

func CreateProgress(name string) *forge.GameObject {
	object := CreateGenericObject(name)

	progress := NewProgress()

	progress.background = NewGraphic()
	progress.activeTrack = NewGraphic()

	object.AddComponent(progress)

	return object
}
