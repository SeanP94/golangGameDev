package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(400, 400, "Hello Window")
	rl.SetTargetFPS(60)

	// Canvas for the shader
	blank := rl.GenImageColor(1024, 1024, rl.Blank)
	texture := rl.LoadTextureFromImage(blank)
	shader := drawShader()
	rl.UnloadImage(blank)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Hello World!", 20, 200, 20, rl.Black)
		rl.BeginShaderMode(shader)
		rl.DrawTexture(texture, 0, 0, rl.Blank)

		rl.EndShaderMode()
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
