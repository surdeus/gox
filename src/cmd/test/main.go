package main

import (
	"github.com/surdeus/gox/src/gx"
)

func main() {
	e := gx.New(&gx.WindowConfig{
		Title: "Test title",
		Width: 480,
		Height: 320,
	})

	e.Run()
}
