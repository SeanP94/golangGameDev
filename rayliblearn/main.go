package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(400, 400, "Hello Window")
	rl.SetTargetFPS(60)

	// Canvas for the shader
	blank := rl.GenImageColor(1024, 1024, rl.Blank)
	texture := rl.LoadTextureFromImage(blank)
	shader := drawFlash()
	rl.UnloadImage(blank)

	loc := rl.GetShaderLocation(shader, "u_time")

	currTime := []float32{float32(rl.GetTime())}

	fmt.Println(loc)
	rl.SetShaderValue(shader, loc, currTime, rl.ShaderUniformFloat)
	// red := rl.GetShaderLocation(shader, "re")
	// green := rl.GetShaderLocation(shader, "gree")
	// blue := rl.GetShaderLocation(shader, "blu")

	// r := float32(1.0)
	// g := float32(1.0)
	// b := float32(1.0)

	// rl.SetShaderValue(shader, red, []float32{r}, rl.ShaderUniformFloat)
	// rl.SetShaderValue(shader, green, []float32{g}, rl.ShaderUniformFloat)
	// rl.SetShaderValue(shader, blue, []float32{b}, rl.ShaderUniformFloat)

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
