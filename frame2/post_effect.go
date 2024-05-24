/*
@author: sk
@date: 2023/5/27
*/
package frame

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type PostEffect struct {
	fbo      uint32
	rbo      uint32
	textures []uint32
}

func (e *PostEffect) initData(width, height int32) {
	// Framebuffer
	gl.GenFramebuffers(1, &e.fbo)
	gl.BindFramebuffer(gl.FRAMEBUFFER, e.fbo)
	// texture
	gl.GenTextures(int32(len(e.textures)), &e.textures[0])
	attachments := make([]uint32, 0)
	for i, texture := range e.textures {
		gl.BindTexture(gl.TEXTURE_2D, texture)
		// 默认支持 HDR
		gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGB16F, width, height, 0, gl.RGB,
			gl.FLOAT, nil)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
		// 颜色 附件
		gl.FramebufferTexture2D(gl.FRAMEBUFFER, uint32(gl.COLOR_ATTACHMENT0+i), gl.TEXTURE_2D, texture, 0)
		if gl.CheckFramebufferStatus(gl.FRAMEBUFFER) != gl.FRAMEBUFFER_COMPLETE {
			panic("PostEffect CheckFramebufferStatus err")
		}
		attachments = append(attachments, uint32(gl.COLOR_ATTACHMENT0+i))
	}
	gl.GenRenderbuffers(1, &e.rbo)
	gl.BindRenderbuffer(gl.RENDERBUFFER, e.rbo)
	gl.RenderbufferStorage(gl.RENDERBUFFER, gl.DEPTH24_STENCIL8, width, height)
	// 深度缓冲 模版附件  必须是一个完整的 附件否则  对应渲染部分 会失效
	gl.FramebufferRenderbuffer(gl.FRAMEBUFFER, gl.DEPTH_STENCIL_ATTACHMENT, gl.RENDERBUFFER, e.rbo)
	if gl.CheckFramebufferStatus(gl.FRAMEBUFFER) != gl.FRAMEBUFFER_COMPLETE {
		panic("PostEffect CheckFramebufferStatus err")
	}
	gl.DrawBuffers(int32(len(attachments)), &attachments[0])
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
}

// Begin 要绘制到 该 表面时调用 使用的shader可以指定 多个输出 到对应的纹理
func (e *PostEffect) Begin() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, e.fbo)
	gl.ClearColor(0, 0, 0, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

// End 绘制完成 要绘制到窗体表面时调用
func (e *PostEffect) End() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
	gl.ClearColor(0, 0, 0, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

// GetTexture 获取第指定个 输出 纹理 用于 自定义 渲染
func (e *PostEffect) GetTexture(index int) uint32 {
	return e.textures[index]
}

// Draw 直接绘制填充到屏幕  shader in 0 vec2 iPos  in 1 vec2 iTex
// 可以使用uniform sampler2D 设置为 0 1 2 ...  接受对应的贴图
func (e *PostEffect) Draw() {
	for i, texture := range e.textures {
		gl.ActiveTexture(uint32(gl.TEXTURE0 + i))
		gl.BindTexture(gl.TEXTURE_2D, texture)
	}
	gl.BindVertexArray(ScreenVao)
	gl.DrawArrays(gl.TRIANGLES, 0, 6)
}

func NewPostEffect(width, height int32, count int) *PostEffect {
	res := &PostEffect{textures: make([]uint32, count)}
	res.initData(width, height)
	return res
}
