/*
@author: sk
@date: 2023/5/29
*/
package main

import (
	"fmt"
	frame "openGL/frame2"
	"openGL/utils"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

func main() {
	frame.Run(NewMainGame(), 1280, 720)
}

type MainGame struct {
	cube, plane          *frame.Obj
	depth, shadow        *frame.Shader
	eyePro               mgl32.Mat4
	screenBuf, screenTex uint32
	camera               *frame.Camera
	lightPos             mgl32.Vec3
}

func NewMainGame() *MainGame {
	return &MainGame{}
}

func (m *MainGame) Init(window *glfw.Window) {
	gl.GenTextures(1, &m.screenTex)
	gl.BindTexture(gl.TEXTURE_CUBE_MAP, m.screenTex)
	for i := 0; i < 6; i++ { // 生成 6方贴图
		gl.TexImage2D(uint32(gl.TEXTURE_CUBE_MAP_POSITIVE_X+i), 0, gl.DEPTH_COMPONENT, 1024, 1024, 0, gl.DEPTH_COMPONENT, gl.FLOAT, nil)
	}
	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_MAG_FILTER, gl.NEAREST)

	gl.GenFramebuffers(1, &m.screenBuf)
	gl.BindFramebuffer(gl.FRAMEBUFFER, m.screenBuf)
	gl.FramebufferTexture(gl.FRAMEBUFFER, gl.DEPTH_ATTACHMENT, m.screenTex, 0)
	gl.DrawBuffer(gl.NONE)
	gl.ReadBuffer(gl.NONE)
	if gl.CheckFramebufferStatus(gl.FRAMEBUFFER) != gl.FRAMEBUFFER_COMPLETE {
		panic("FRAMEBUFFER fail")
	}

	m.cube = frame.NewObj(cubeVs, 36)
	m.cube.VertexAttr(0, 3, 3*4, 0)
	m.cube.BindTexture(gl.TEXTURE0, m.screenTex)
	m.plane = frame.NewObj(planeVs, 6)
	m.plane.VertexAttr(0, 3, 3*4, 0)
	m.plane.BindTexture(gl.TEXTURE0, m.screenTex)

	m.eyePro = utils.GetDefaultPerspective()
	m.camera = frame.NewCamera()
	m.lightPos = mgl32.Vec3{0, 1, 0}

	m.depth = frame.NewShader("test23/depth")
	m.depth.Use()
	dirs := []mgl32.Vec3{{1, 0, 0}, {-1, 0, 0}, {0, 1, 0}, {0, -1, 0}, {0, 0, 1}, {0, 0, -1}}
	ups := []mgl32.Vec3{{0, -1, 0}, {0, -1, 0}, {0, 0, 1}, {0, 0, -1}, {0, -1, 0}, {0, -1, 0}}
	for i := 0; i < 6; i++ {
		pv := m.eyePro.Mul4(mgl32.LookAtV(m.lightPos, m.lightPos.Add(dirs[i]), ups[i]))
		m.depth.SetMat4(fmt.Sprintf("uShadow[%d]", i), &pv[i])
	}
	m.depth.Set3fv("uLightPos", &m.lightPos[0])
	m.depth.Set1f("uFar", 100)

	m.shadow = frame.NewShader("test23/shadow")
	m.shadow.Use()
	m.shadow.Set3fv("uLightPos", &m.lightPos[0])
	m.shadow.Set1f("uFar", 100)
}

func (m *MainGame) Update(window *glfw.Window) {
	frame.ApplyInput(m.camera, window)
}

func (m *MainGame) Draw() {
	gl.Viewport(0, 0, 1024, 1024)
	gl.BindFramebuffer(gl.FRAMEBUFFER, m.screenBuf)
	gl.Clear(gl.DEPTH_BUFFER_BIT)
	m.depth.Use()
	m.cube.Reset()
	m.cube.Translate(-2, 1, 1.5)
	m.cube.Draw(m.depth)
	m.cube.Reset()
	m.cube.Translate(2, 1, -1.5)
	m.cube.Draw(m.depth)
	m.plane.Draw(m.depth)

	gl.Viewport(0, 0, 1280*2, 720*2)
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	m.shadow.Use()
	view := m.camera.GetView()
	m.shadow.SetMat4("uView", &view[0])
	m.shadow.SetMat4("uProjection", &m.eyePro[0])
	m.cube.Reset()
	m.cube.Translate(-2, 1, 1.5)
	m.cube.Draw(m.depth)
	m.cube.Reset()
	m.cube.Translate(2, 1, -1.5)
	m.cube.Draw(m.depth)
	m.plane.Draw(m.depth)
}
