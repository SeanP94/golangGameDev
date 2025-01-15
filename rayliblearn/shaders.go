package main

import rl "github.com/gen2brain/raylib-go/raylib"

func drawShader() rl.Shader {
	shader := rl.LoadShader("", "shaders/basic.fs")
	return shader
}
