/*
@author: sk
@date: 2023/5/27
*/
package main

import (
	"openGL/frame3"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type PostEffect struct {
	fbo      uint32
	rbo      uint32
	textures []uint32
	shader   *frame3.Shader
}

func (e *PostEffect) initData(width, height int32) {
	// Framebuffer
	gl.GenFramebuffers(1, &e.fbo)
	gl.BindFramebuffer(gl.FRAMEBUFFER, e.fbo)
	// textures
	gl.GenTextures(2, &e.textures[0])
	for i := 0; i < 2; i++ {
		// textures[i]
		gl.BindTexture(gl.TEXTURE_2D, e.textures[i])
		// 默认支持 HDR
		gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGB16F, width, height, 0, gl.RGB,
			gl.FLOAT, nil)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
		gl.FramebufferTexture2D(gl.FRAMEBUFFER, uint32(gl.COLOR_ATTACHMENT0+i), gl.TEXTURE_2D, e.textures[i], 0)
		if gl.CheckFramebufferStatus(gl.FRAMEBUFFER) != gl.FRAMEBUFFER_COMPLETE {
			panic("PostEffect CheckFramebufferStatus err")
		}
	}
	targets := []uint32{gl.COLOR_ATTACHMENT0, gl.COLOR_ATTACHMENT1}
	gl.DrawBuffers(2, &targets[0]) // 需要显示指定渲染目标
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
}

func (e *PostEffect) Begin() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, e.fbo)
	gl.ClearColor(0, 0, 0, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func (e *PostEffect) End() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
	gl.ClearColor(0, 0, 0, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func (e *PostEffect) Draw() {
	e.shader.Use()
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, e.textures[0])
	gl.ActiveTexture(gl.TEXTURE1)
	gl.BindTexture(gl.TEXTURE_2D, e.textures[1])
	gl.BindVertexArray(frame3.ScreenVao)
	gl.DrawArrays(gl.TRIANGLES, 0, 6)
}

func NewPostEffect(width, height int32, name string) *PostEffect {
	res := &PostEffect{shader: frame3.NewShader(name), textures: make([]uint32, 2)}
	res.shader.Use()
	res.shader.Set1i("uImage0", 0)
	res.shader.Set1i("uImage1", 1)
	res.initData(width, height)
	return res
}
