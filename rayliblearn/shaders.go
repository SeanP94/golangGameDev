package main

import rl "github.com/gen2brain/raylib-go/raylib"

func drawShader() rl.Shader {
	shader := rl.LoadShader("", "shaders/fs.fs.glsl")
	return shader
}

func drawFlash() rl.Shader {
	shader := rl.LoadShader("", "shaders/basic.glsl")
	return shader
}
