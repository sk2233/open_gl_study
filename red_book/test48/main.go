/*
@author: sk
@date: 2023/6/23
*/
package main

import (
	frame "openGL/frame2"
	"openGL/utils"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

func main() {
	frame.Run(NewMainGame(), 1280, 720)
}

type MainGame struct {
	eye    *frame.Camera
	objs   []*PointObj
	shader *frame.Shader
}

func NewMainGame() *MainGame {
	return &MainGame{eye: frame.NewCamera(), objs: make([]*PointObj, 0)}
}

func (m *MainGame) Init(window *glfw.Window) {
	m.eye.SetPos(mgl32.Vec3{10, 10, 10})

	vxl := utils.ParseVxl("test48/shad.vxl")
	ra2Objs := vxl.ParseRa2Objs()
	for _, ra2Obj := range ra2Objs {
		obj := NewPointObj(ra2Obj.Data, int32(len(ra2Obj.Data)/9))
		obj.VertexAttr(0, 3, 9*4, 0)
		obj.VertexAttr(1, 3, 9*4, 3*4)
		//obj.VertexAttr(2, 3, 9*4, 6*4)
		obj.SetModel(ra2Obj.Transform)
		obj.Scale(ra2Obj.Scale, ra2Obj.Scale, ra2Obj.Scale)
		m.objs = append(m.objs, obj)
	}

	m.shader = frame.NewShader("test48/obj")
	m.shader.Use()
	p := mgl32.Perspective(45, 1280.0/720.0, 0.1, 1000)
	m.shader.SetMat4("uProjection", &p[0])
}

func (m *MainGame) Update(window *glfw.Window) {
	frame.ApplyInput(m.eye, window)
}

func (m *MainGame) Draw() {
	m.shader.Use()
	v := m.eye.GetView()
	m.shader.SetMat4("uView", &v[0])
	for _, obj := range m.objs {
		obj.Draw(m.shader)
	}
}
