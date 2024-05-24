/*
@author: sk
@date: 2023/5/27
*/
package frame3

import (
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type Shader struct {
	program uint32
	cache   map[string]int32
}

func (s *Shader) SetMat4(name string, pos *float32) {
	location := s.getLocation(name)
	gl.UniformMatrix4fv(location, 1, false, pos)
}

func (s *Shader) getLocation(name string) int32 {
	if _, ok := s.cache[name]; !ok {
		s.cache[name] = gl.GetUniformLocation(s.program, gl.Str(name+"\x00"))
	}
	return s.cache[name]
}

func (s *Shader) Use() {
	gl.UseProgram(s.program)
}

func (s *Shader) Set1i(name string, v int32) {
	location := s.getLocation(name)
	gl.Uniform1i(location, v)
}

func (s *Shader) Set1f(name string, v float32) {
	location := s.getLocation(name)
	gl.Uniform1f(location, v)
}

func (s *Shader) Set3fv(name string, v *float32) {
	location := s.getLocation(name)
	gl.Uniform3fv(location, 1, v)
}

func (s *Shader) Set4fv(name string, v *float32) {
	location := s.getLocation(name)
	gl.Uniform4fv(location, 1, v)
}

func (s *Shader) initProgram(name string) {
	s.program = gl.CreateProgram()
	// 加载顺序 也是调用顺序
	if Has(name + ".vert") {
		vShader := s.loadShader(name+".vert", gl.VERTEX_SHADER)
		gl.AttachShader(s.program, vShader)
		defer gl.DeleteShader(vShader)
		fmt.Println(name + " has-----vert")
	}
	if Has(name + ".geom") {
		gShader := s.loadShader(name+".geom", gl.GEOMETRY_SHADER)
		gl.AttachShader(s.program, gShader)
		defer gl.DeleteShader(gShader)
		fmt.Println(name + " has-----geom")
	}
	if Has(name + ".frag") {
		fShader := s.loadShader(name+".frag", gl.FRAGMENT_SHADER)
		gl.AttachShader(s.program, fShader)
		defer gl.DeleteShader(fShader)
		fmt.Println(name + " has-----frag")
	}
	gl.LinkProgram(s.program)
	var status int32
	gl.GetProgramiv(s.program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		panic("LinkProgram fail name " + name)
	}
}

func (s *Shader) loadShader(path string, type0 uint32) uint32 {
	res := gl.CreateShader(type0) // 创建单个shader
	bs := ReadAll(path)
	str, free := gl.Strs(string(bs)) // 绑定 Shader
	gl.ShaderSource(res, 1, str, nil)
	gl.CompileShader(res) // 编译shader
	var status int32
	gl.GetShaderiv(res, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		panic("CompileShader fail path " + path)
	}
	free()
	return res
}

func NewShader(name string) *Shader {
	res := &Shader{cache: make(map[string]int32)}
	res.initProgram(name)
	return res
}
