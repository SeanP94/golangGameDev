package main

import (
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
	shader := drawBOS()
	rl.UnloadImage(blank)

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

		rl.DrawText("Hello World!", 20, 200, 20, rl.Black)
		rl.EndDrawing()
	}
	rl.UnloadShader(shader)
	rl.UnloadTexture(texture)
	rl.CloseWindow()
}
