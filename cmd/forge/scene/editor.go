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
	"github.com/haakenlabs/forge/scene"
	"github.com/haakenlabs/forge/scene/effects"
)

const NameEditor = "editor"

func NewEditorScene() *forge.Scene {
	s := forge.NewScene(NameEditor)
	s.SetLoadFunc(func() error {
		testObject := forge.NewGameObject("testObject")
		camera := scene.CreateCamera("camera", true, forge.RenderPathDeferred)
		camera.AddComponent(scene.NewControlOrbit())
		tonemapper := effects.NewTonemapper()

		cameraC := forge.CameraComponent(camera)
		cameraC.AddEffect(tonemapper)

		toneControl := scene.NewControlExposure()
		toneControl.SetTonemapper(tonemapper)
		camera.AddComponent(toneControl)

		test := scene.CreateOrb("orb")

		scene.ControlOrbitComponent(camera).Target = test.Transform()

		if err := s.Graph().AddObject(testObject, nil); err != nil {
			return err
		}
		if err := s.Graph().AddObject(camera, nil); err != nil {
			return err
		}
		if err := s.Graph().AddObject(test, nil); err != nil {
			return err
		}

		return nil
	})

	return s
}
