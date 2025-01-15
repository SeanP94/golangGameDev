package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(400, 400, "Hello Window")
	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Hello World!", 20, 200, 20, rl.Black)

		shader := drawShader()

		rl.BeginShaderMode(shader)

		rl.EndShaderMode()
		rl.EndDrawing()
	}

	rl.CloseWindow()

}
