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
	"fmt"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/haakenlabs/forge"
	"github.com/haakenlabs/forge/particle"
	"github.com/haakenlabs/forge/scene"
	"github.com/haakenlabs/forge/scene/effects"
	"github.com/haakenlabs/forge/system/input"
	"github.com/haakenlabs/forge/ui"
	"github.com/haakenlabs/forge/ui/prefabs"
	"github.com/haakenlabs/forge/ui/widget"
)

const NameStart = "start"

type Inspector struct {
	forge.BaseScriptComponent

	psys *particle.System

	uiObject           *forge.GameObject
	labelStartLifetime *widget.Label
	labelPlaybackSpeed *widget.Label
	labelEmissionRate  *widget.Label
	labelMaxParticles  *widget.Label
	labelCurParticles  *widget.Label

	show bool
}

func (i *Inspector) LateUpdate() {
	i.labelStartLifetime.SetValue(fmt.Sprintf("Start Lifetime: %.0f", i.psys.Core.StartLifetime))
	i.labelPlaybackSpeed.SetValue(fmt.Sprintf("Playback Speed: %.2f", i.psys.Core.PlaybackSpeed))
	i.labelEmissionRate.SetValue(fmt.Sprintf("Emission Rate: %.0f", i.psys.Emission.Rate))
	i.labelMaxParticles.SetValue(fmt.Sprintf("Max Particles: %d", i.psys.Core.MaxParticles()))
	i.labelCurParticles.SetValue(fmt.Sprintf("Particle Count: %d", i.psys.Core.ParticleCount()))

	if input.KeyDown(glfw.KeyF1) {
		if i.show {
			i.uiObject.SetActive(false)
			i.show = false
		} else {
			i.uiObject.SetActive(true)
			i.show = true
		}
	}
}

func (i *Inspector) ToggleParticleSystem() {
	i.psys.Emission.Rate = 0
}

func (i *Inspector) PauseParticles(state widget.CheckState) {
	if state == widget.CheckStateOff {
		i.psys.Core.PlaybackSpeed = 0.0
	} else {
		i.psys.Core.PlaybackSpeed = 1.0
	}
}

func (i *Inspector) SetEmissionRate(rate float64) {
	i.psys.Emission.Rate = float32(rate)
}

func makeUI(psys *particle.System) *forge.GameObject {
	controller := ui.CreateController("ui_controller")

	panel := widget.CreatePanel("test_panel")
	//panel.SetActive(false)

	button := widget.CreateButton("test_button")
	rt := ui.RectTransformComponent(button)
	rt.SetPosition2D(mgl32.Vec2{8, 128})
	widget.ButtonComponent(button).SetValue("Reset")

	checkbox := widget.CreateCheckbox("test_checkbox")
	rt = ui.RectTransformComponent(checkbox)
	rt.SetPosition2D(mgl32.Vec2{8, 192})

	progress := widget.CreateProgress("test_progress")
	rt = ui.RectTransformComponent(progress)
	rt.SetPosition2D(mgl32.Vec2{8, 220})
	rt.SetSize(mgl32.Vec2{256, 16})
	widget.ProgressComponent(progress).SetProgress(0.6)

	slider := widget.CreateSlider("test_slider")
	rt = ui.RectTransformComponent(slider)
	rt.SetPosition2D(mgl32.Vec2{8, 240})
	rt.SetSize(mgl32.Vec2{256, 16})

	inspector := &Inspector{
		psys:     psys,
		uiObject: panel,
	}
	widget.ButtonComponent(button).SetOnPressedFunc(inspector.ToggleParticleSystem)
	widget.CheckboxComponent(checkbox).SetOnChangeFunc(inspector.PauseParticles)
	widget.SliderComponent(slider).SetOnChangeFunc(inspector.SetEmissionRate)
	widget.SliderComponent(slider).SetMaxValue(100000)
	widget.SliderComponent(slider).SetValue(1000)

	labelStartLifetime := widget.CreateLabel("label_startlifetime")
	{
		ui.RectTransformComponent(labelStartLifetime).SetPosition2D(mgl32.Vec2{8, 8})
		lc := widget.LabelComponent(labelStartLifetime)
		lc.SetValue("Start Lifetime: -")
		inspector.labelStartLifetime = lc
	}

	labelPlaybackSpeed := widget.CreateLabel("label_playbackspeed")
	{
		ui.RectTransformComponent(labelPlaybackSpeed).SetPosition2D(mgl32.Vec2{8, 24})
		lc := widget.LabelComponent(labelPlaybackSpeed)
		lc.SetValue("Playback Speed: -")
		inspector.labelPlaybackSpeed = lc
	}

	labelEmissionRate := widget.CreateLabel("label_emissionrate")
	{
		ui.RectTransformComponent(labelEmissionRate).SetPosition2D(mgl32.Vec2{8, 40})
		lc := widget.LabelComponent(labelEmissionRate)
		lc.SetValue("Emission Rate: -")
		inspector.labelEmissionRate = lc
	}

	labelMaxParticles := widget.CreateLabel("label_maxparticles")
	{
		ui.RectTransformComponent(labelMaxParticles).SetPosition2D(mgl32.Vec2{8, 56})
		lc := widget.LabelComponent(labelMaxParticles)
		lc.SetValue("Max Particles: -")
		inspector.labelMaxParticles = lc
	}

	labelCurParticles := widget.CreateLabel("label_particlecount")
	{
		ui.RectTransformComponent(labelCurParticles).SetPosition2D(mgl32.Vec2{8, 72})
		lc := widget.LabelComponent(labelCurParticles)
		lc.SetValue("Particle Count: -")
		inspector.labelCurParticles = lc
	}

	ui.RectTransformComponent(panel).SetPosition2D(mgl32.Vec2{16, 16})

	panel.AddChild(labelStartLifetime)
	panel.AddChild(labelPlaybackSpeed)
	panel.AddChild(labelEmissionRate)
	panel.AddChild(labelMaxParticles)
	panel.AddChild(labelCurParticles)
	panel.AddChild(button)
	panel.AddChild(checkbox)
	panel.AddChild(progress)
	panel.AddChild(slider)

	controller.AddChild(panel)
	controller.AddComponent(inspector)

	return controller
}

func NewStartScene() *forge.Scene {
	s := forge.NewScene(NameStart)
	s.SetLoadFunc(func() error {
		camera := scene.CreateCamera("camera", true, forge.RenderPathForward)
		camera.AddComponent(scene.NewControlOrbit())
		tonemapper := effects.NewTonemapper()

		cameraC := forge.CameraComponent(camera)
		cameraC.AddEffect(tonemapper)
		cameraC.SetClearMode(forge.ClearModeColor)

		toneControl := scene.NewControlExposure()
		toneControl.SetTonemapper(tonemapper)
		camera.AddComponent(toneControl)

		target := forge.NewGameObject("target")

		psys := particle.NewParticleSystem(1000000)
		psys.Emission.Rate = 1000
		psys.Core.StartLifetime = 5
		psys.Core.PlaybackSpeed = 1.0

		target.AddComponent(psys)

		scene.ControlOrbitComponent(camera).Target = target.Transform()

		if err := s.Graph().AddObject(target, nil); err != nil {
			return err
		}
		if err := s.Graph().AddObject(camera, nil); err != nil {
			return err
		}
		if err := s.Graph().AddObject(prefabs.NewDebug("debugger"), nil); err != nil {
			return err
		}

		return nil
	})

	return s
}
