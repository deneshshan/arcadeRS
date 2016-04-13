package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
	"time"
)

const WindowTitle = "arcadeRS"
const WindowHeight = 600
const WindowWidth = 800

func main() {
	window, err := sdl.CreateWindow(WindowTitle, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, WindowWidth,
		WindowHeight, sdl.WINDOW_OPENGL)
	HandleErrorAndExit("Error creating sdl window: {}", err)

	renderer, err := sdl.CreateRenderer(window, -1, 0)
	HandleErrorAndExit("Error creating renderer: {}", err)

	err = renderer.SetDrawColor(0, 0, 0, 0)
	err = renderer.Clear()
	renderer.Present()
	HandleErrorAndExit("Error when working with renderer: {}", err)

	time.Sleep(5 * time.Second)

	fmt.Println("BYE!!!")
	os.Exit(0)
}

func HandleErrorAndExit(custom_message string, err error) {
	if err != nil {
		fmt.Println(custom_message, err)
		os.Exit(1)
	}
}
