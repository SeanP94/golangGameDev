package main

import rl "github.com/gen2brain/raylib-go/raylib"

func drawShader() rl.Shader {
	shader := rl.LoadShader("shaders/vs.vs.glsl", "shaders/fs.fs.glsl")
	return shader
}

func drawFlash() rl.Shader {
	shader := rl.LoadShader("shaders/vs.vs.glsl", "shaders/basic.glsl")
	return shader
}

func drawBOS() rl.Shader {
	shader := rl.LoadShader("shaders/bos.vs", "shaders/bos.fs")
	return shader
}

func updateCoreShader(shader *rl.Shader) {
	/*
		TODO: This should be in the Shader object class you create later, where each shader saves its u_time, u_res etc...
	*/
	currTime := float32(rl.GetTime())
	time_loc := rl.GetShaderLocation(*shader, "u_time")
	resolutionLoc := rl.GetShaderLocation(*shader, "u_resolution")

	currSize := []float32{float32(SCREEN_WIDTH), float32(SCREEN_HEIGHT)}

	rl.SetShaderValue(*shader, resolutionLoc, currSize, rl.ShaderUniformVec2)
	mouse_loc := rl.GetShaderLocation(*shader, "u_mouse")
	rl.SetShaderValue(*shader, time_loc, []float32{currTime}, rl.ShaderUniformFloat)
	rl.SetShaderValue(*shader, mouse_loc, []float32{float32(rl.GetMousePosition().X), float32(rl.GetMousePosition().Y)}, rl.ShaderUniformVec2)
}
