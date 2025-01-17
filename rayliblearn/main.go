package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var SCREEN_WIDTH int32 = 800
var SCREEN_HEIGHT int32 = 800

func main() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "Hello Window")
	rl.SetTargetFPS(60)

	// Canvas for the shader
	blank := rl.GenImageColor(1024, 1024, rl.Blank)
	texture := rl.LoadTextureFromImage(blank)
	shader := drawShader("crt.fs")
	rl.UnloadImage(blank)

	curve := rl.GetShaderLocation(shader, "curvature")

	rl.SetShaderValue(shader, curve, []float32{3.0}, rl.ShaderUniformFloat)

	if !rl.IsShaderValid(shader) {
		println("Shader failed to compile!")
		rl.SetTraceLogLevel(rl.LogAll)
		return
	}

	for !rl.WindowShouldClose() {
		updateCoreShader(&shader)
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.BeginShaderMode(shader)

		{
			rl.DrawTexture(texture, 0, 0, rl.White)
		}

		rl.EndShaderMode()

		rl.DrawText("X: "+strconv.FormatFloat(float64(rl.GetMousePosition().X), 'f', 2, 64), 0, 0, 20, rl.Black)
		rl.DrawText("Y: "+strconv.FormatFloat(float64(rl.GetMousePosition().Y), 'f', 2, 64), 100, 0, 20, rl.Black)
		rl.EndDrawing()
	}
	rl.UnloadShader(shader)
	rl.UnloadTexture(texture)
	rl.CloseWindow()
}
