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

package window

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"

	"github.com/haakenlabs/forge"
	"github.com/haakenlabs/forge/math"
)

func AspectRatio() float32 {
	return forge.GetWindow().AspectRatio()
}

func CenterWindow() {
	forge.GetWindow().CenterWindow()
}

func ClearBuffers() {
	forge.GetWindow().ClearBuffers()
}

func EnableVsync(enable bool) {
	forge.GetWindow().EnableVsync(enable)
}

func Resolution() math.IVec2 {
	return forge.GetWindow().Resolution()
}

func SetSize(size math.IVec2) {
	forge.GetWindow().SetSize(size)
}

func SwapBuffers() {
	forge.GetWindow().SwapBuffers()
}

func Vsync() bool {
	return forge.GetWindow().Vsync()
}

func GLFWWindow() *glfw.Window {
	return forge.GetWindow().GLFWWindow()
}

func OrthoMatrix() mgl32.Mat4 {
	return forge.GetWindow().OrthoMatrix()
}
