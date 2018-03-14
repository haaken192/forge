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

package input

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"

	"github.com/haakenlabs/forge"
)

func KeyDown(key glfw.Key) bool {
	return forge.GetWindow().KeyDown(key)
}

func KeyUp(key glfw.Key) bool {
	return forge.GetWindow().KeyUp(key)
}

func KeyPressed() bool {
	return forge.GetWindow().KeyPressed()
}

func MouseDown(button glfw.MouseButton) bool {
	return forge.GetWindow().MouseDown(button)
}

func MouseUp(button glfw.MouseButton) bool {
	return forge.GetWindow().MouseUp(button)
}

func MouseWheelX() float64 {
	return forge.GetWindow().MouseWheelX()
}

func MouseWheelY() float64 {
	return forge.GetWindow().MouseWheelY()
}

func MouseWheel() bool {
	return forge.GetWindow().MouseWheel()
}

func MouseMoved() bool {
	return forge.GetWindow().MouseMoved()
}

func MousePressed() bool {
	return forge.GetWindow().MousePressed()
}

func MousePosition() mgl32.Vec2 {
	return forge.GetWindow().MousePosition()
}

func WindowResized() bool {
	return forge.GetWindow().WindowResized()
}

func ShouldClose() bool {
	return forge.GetWindow().ShouldClose()
}

func HandleEvents() {
	forge.GetWindow().HandleEvents()
}

func HasEvents() bool {
	return forge.GetWindow().HasEvents()
}
