package gx

import (
	"github.com/hajimehoshi/ebiten/v2"
	//"fmt"
)

type Shader = ebiten.Shader

var (
	// The shader does not 
	SolidWhiteColorShader = MustNewShader([]byte(`
		package main
		
		func Fragment(p vec4, coord vec2, color vec4) vec4 {
			return vec4(1, 1, 1, 1)
		}
	`))
)

func MustNewShader(src []byte) (*Shader) {
	shader, err := NewShader(src)
	if err != nil {
		panic(err)
	}
	
	return shader
}

func NewShader(src []byte) (*Shader, error) {
	return ebiten.NewShader(src)
}

