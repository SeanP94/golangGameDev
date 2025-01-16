package main

import (
	"fmt"

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
	shader := drawFlash()
	rl.UnloadImage(blank)

	loc := rl.GetShaderLocation(shader, "u_time")
	loc_size := rl.GetShaderLocationAttrib(shader, "u_resolution")

	currTime := []float32{float32(rl.GetTime())}
	currSize := []float32{float32(SCREEN_WIDTH), float32(SCREEN_HEIGHT)}
	fmt.Println(loc)
	fmt.Println(loc_size)
	rl.SetShaderValue(shader, loc, currTime, rl.ShaderUniformFloat)
	rl.SetShaderValue(shader, loc_size, currSize, rl.ShaderUniformVec2)

	for !rl.WindowShouldClose() {
		currTime := float32(rl.GetTime())
		rl.SetShaderValue(shader, loc, []float32{currTime}, rl.ShaderUniformFloat)

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
