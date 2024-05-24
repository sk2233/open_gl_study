/*
@author: sk
@date: 2023/5/28
*/
package main

import (
	"fmt"
	"openGL/frame3"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func main() {
	frame3.Run(NewStencilGame(), 1280, 720)
}

type StencilGame struct {
	sprite, bg *frame3.Sprite
	x, y       float32
}

func NewStencilGame() *StencilGame {
	return &StencilGame{}
}

func (s *StencilGame) Init(window *glfw.Window) {
	s.sprite = frame3.NewSprite(frame3.NewShader("test21/base"), frame3.NewTexture("test21/test.png"))
	s.bg = frame3.NewSprite(frame3.CreateDefaultShader(), frame3.NewTexture("test21/bg.png"))
	gl.Enable(gl.STENCIL_TEST)
	gl.StencilOp(gl.KEEP, gl.KEEP, gl.REPLACE)
}

var (
	count   = 0
	lastSec = int64(0)
)

func (s *StencilGame) Update(window *glfw.Window) {
	s.x += frame3.GetAxle(window, glfw.KeyA, glfw.KeyD) * 3
	s.y += frame3.GetAxle(window, glfw.KeyW, glfw.KeyS) * 3
	sec := time.Now().Unix()
	count++
	if lastSec != sec {
		fmt.Println(count)
		lastSec = sec
		count = 0
	}
}

func (s *StencilGame) Draw() {
	gl.StencilFunc(gl.ALWAYS, 0, 0xFF)
	gl.StencilMask(0xFF)
	gl.Clear(gl.STENCIL_BUFFER_BIT)

	// 设置为   2233
	gl.StencilFunc(gl.ALWAYS, 2233, 0xFF)
	gl.StencilMask(0xFF)
	s.sprite.Draw(complex(s.x, s.y), 3+3i, 0, frame3.ColorWhite)

	// 设置  为  3333
	gl.StencilFunc(gl.ALWAYS, 3333, 0xFF)
	gl.StencilMask(0xFF)
	s.sprite.Draw(complex(s.x+100, s.y), 3+3i, 0, frame3.ColorWhite)
	//for i := 0; i < 5000; i++ {
	//
	//}

	// 模版值 必须  大于  2233 才行    前面的也会绘制 这里只是遮住了前面的 前面的可以考虑 取消绘制
	gl.StencilFunc(gl.LEQUAL, 2233, 0xFF)
	gl.StencilMask(0x00)
	s.bg.Draw(0, 1+1i, 0, frame3.ColorWhite)
}
