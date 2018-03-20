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

package prefabs

import (
	"fmt"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/haakenlabs/forge"
	"github.com/haakenlabs/forge/system/input"
	"github.com/haakenlabs/forge/ui"
)

type Debug struct {
	forge.BaseScriptComponent

	labelTitle *ui.Label
}

func (d *Debug) LateUpdate() {
	if input.KeyDown(glfw.KeyF2) {
		fmt.Printf("this: %v parent: %v\n",
			d.labelTitle.RectTransform().Rect(),
			ui.RectTransformComponent(d.labelTitle.GameObject().Parent()).Rect())
	}
}

func NewDebug(name string) *forge.GameObject {
	o := ui.CreateController(name + "object")

	p := ui.CreatePanel(name + "-panel")
	ui.ImageComponent(p).RectTransform().SetSize(mgl32.Vec2{320, 512})
	ui.ImageComponent(p).SetColor(ui.Styles.BackgroundColor)

	s0 := ui.CreatePanel(name + "-s0")
	ui.ImageComponent(s0).RectTransform().SetSize(mgl32.Vec2{320, 18})
	ui.ImageComponent(s0).SetColor(ui.Styles.WidgetColor)

	s0Title := ui.CreateLabel(name + "-s0-title")
	ui.LabelComponent(s0Title).SetValue("Forge Debugger")
	ui.RectTransformComponent(s0Title).SetPosition2D(mgl32.Vec2{4, 0})
	ui.RectTransformComponent(s0Title).SetPresets(ui.AnchorMiddleLeft, ui.PivotMiddleLeft)

	s0.AddChild(s0Title)
	p.AddChild(s0)

	d := &Debug{
		labelTitle: ui.LabelComponent(s0Title),
	}

	d.SetName(name + "-debug")
	forge.GetInstance().MustAssign(d)

	o.AddChild(p)
	o.AddComponent(d)

	return o
}
