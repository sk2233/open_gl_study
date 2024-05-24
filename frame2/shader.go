/*
@author: sk
@date: 2023/5/14
*/
package frame

import (
	"openGL/utils"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type Shader struct {
	program uint32
	cache   map[string]int32
}

func (s *Shader) SetMat4(name string, pos *float32) {
	location := s.GetLocation(name)
	gl.UniformMatrix4fv(location, 1, false, pos)
}

func (s *Shader) GetLocation(name string) int32 {
	if _, ok := s.cache[name]; !ok {
		s.cache[name] = gl.GetUniformLocation(s.program, gl.Str(name+"\x00"))
	}
	return s.cache[name]
}

func (s *Shader) Use() {
	gl.UseProgram(s.program)
}

func (s *Shader) Set1i(name string, v int32) {
	location := s.GetLocation(name)
	gl.Uniform1i(location, v)
}

func (s *Shader) Set1f(name string, v float32) {
	location := s.GetLocation(name)
	gl.Uniform1f(location, v)
}

func (s *Shader) Set3fv(name string, v *float32) {
	location := s.GetLocation(name)
	gl.Uniform3fv(location, 1, v)
}

func (s *Shader) SetBool(name string, v bool) {
	location := s.GetLocation(name)
	if v {
		gl.Uniform1i(location, gl.TRUE)
	} else {
		gl.Uniform1i(location, gl.FALSE)
	}
}

func (s *Shader) Set2fv(name string, v *float32) {
	location := s.GetLocation(name)
	gl.Uniform2fv(location, 1, v)
}

func NewShader(name string) *Shader {
	return &Shader{program: utils.LoadProgram(name), cache: make(map[string]int32)}
}
