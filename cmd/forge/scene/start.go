/*
Copyright (c) 2017 HaakenLabs

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

package scene

import (
	"github.com/haakenlabs/forge"
	"github.com/haakenlabs/forge/system/asset/image"
	"github.com/haakenlabs/forge/ui"
	"github.com/haakenlabs/forge/ui/widget"
)

const NameStart = "start"

func makeSplash() *forge.GameObject {
	o := ui.CreateController("splash_controller")

	p := widget.CreatePanel("splash-background")
	widget.ImageComponent(p).SetColor(ui.Styles.BackgroundColor)
	ui.RectTransformComponent(p).SetAnchorPreset(ui.StretchAnchorAll)
	ui.RectTransformComponent(p).SetSize(forge.GetWindow().Resolution().Vec2())

	l := widget.CreateImage("splash-logo")
	widget.ImageComponent(l).SetTexture(image.MustGet("arc-logo.png"))
	ui.RectTransformComponent(l).SetPresets(ui.AnchorMiddleCenter, ui.PivotMiddleCenter)

	p.AddChild(l)
	o.AddChild(p)

	return o
}

func NewStartScene() *forge.Scene {
	s := forge.NewScene(NameStart)
	s.SetLoadFunc(func() error {

		if err := s.Graph().AddObject(makeSplash(), nil); err != nil {
			return err
		}

		return nil
	})

	return s
}
