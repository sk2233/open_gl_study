/*
@author: sk
@date: 2023/5/3
*/
package utils

import (
	"bufio"
	"fmt"
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

func BindBuffer[T uint32 | float32](data []T, type0 uint32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(type0, vbo)
	gl.BufferData(type0, len(data)*4, gl.Ptr(data), gl.STATIC_DRAW)
	return vbo
}

func CreateVertexArray() uint32 {
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	return vao
}

func LoadShader(name string, type0 uint32) uint32 {
	res := gl.CreateShader(type0) // 创建单个shader
	bs := ReadAll(name)
	str, free := gl.Strs(string(bs)) // 绑定 shader
	gl.ShaderSource(res, 1, str, nil)
	gl.CompileShader(res) // 编译shader
	var status int32
	gl.GetShaderiv(res, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		panic("CompileShader fail")
	}
	free()
	return res
}

func LoadProgram(name string) uint32 {
	program := gl.CreateProgram()
	// 加载顺序 也是调用顺序
	if Has(name + ".vert") {
		vShader := LoadShader(name+".vert", gl.VERTEX_SHADER)
		gl.AttachShader(program, vShader)
		defer gl.DeleteShader(vShader)
		fmt.Println("has-----vert")
	} else {
		panic("顶点着色器是必须的")
	}
	if Has(name + ".geom") {
		gShader := LoadShader(name+".geom", gl.GEOMETRY_SHADER)
		gl.AttachShader(program, gShader)
		defer gl.DeleteShader(gShader)
		fmt.Println("has-----geom")
	}
	if Has(name + ".frag") {
		fShader := LoadShader(name+".frag", gl.FRAGMENT_SHADER)
		gl.AttachShader(program, fShader)
		defer gl.DeleteShader(fShader)
		fmt.Println("has-----frag")
	} else {
		panic("片元着色器是必须的")
	}
	gl.LinkProgram(program)
	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		panic("LinkProgram fail")
	}
	return program
}

func LoadTexture(name string) uint32 {
	var texture uint32
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)

	img := LoadImage(name)
	size := img.Bounds()
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(size.Dx()), int32(size.Dy()), 0, gl.RGBA,
		gl.UNSIGNED_BYTE, gl.Ptr(img.Pix))
	// 空间换性能 远处纹理的采样
	gl.GenerateMipmap(gl.TEXTURE_2D)

	// 缩放时采样规律
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	// 采样超出  0~1 时的行为
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	return texture
}

func Load3DTexture(names ...string) uint32 {
	var texture uint32
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_3D, texture)

	for i, name := range names {
		img := LoadImage(name)
		size := img.Bounds()
		gl.TexImage3D(gl.TEXTURE_3D, 0, gl.RGBA, int32(size.Dx()), int32(size.Dy()), int32(i), 0, gl.RGBA,
			gl.UNSIGNED_BYTE, gl.Ptr(img.Pix))
	}

	// 缩放时采样规律
	gl.TexParameteri(gl.TEXTURE_3D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_3D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	// 采样超出  0~1 时的行为
	gl.TexParameteri(gl.TEXTURE_3D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_3D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_3D, gl.TEXTURE_WRAP_R, gl.CLAMP_TO_EDGE)
	return texture
}

// LoadCubeTexture 贴图顺序必须是 right left top bottom front back
func LoadCubeTexture(names ...string) uint32 { // 立方体贴图 必须  6个
	var texture uint32
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_CUBE_MAP, texture)

	for i := 0; i < 6; i++ {
		img := LoadImage(names[i])
		size := img.Bounds()
		gl.TexImage2D(uint32(gl.TEXTURE_CUBE_MAP_POSITIVE_X+i), 0, gl.RGBA, int32(size.Dx()), int32(size.Dy()), 0, gl.RGBA,
			gl.UNSIGNED_BYTE, gl.Ptr(img.Pix))
	}

	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_WRAP_R, gl.REPEAT) // 多了一维度 的采样规则
	return texture
}

func LoadImage(name string) *image.RGBA {
	reader, err := os.Open(BasePath + name)
	HandleErr(err)
	img, _, err := image.Decode(reader)
	HandleErr(err)
	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Pt(0, 0), draw.Src)
	return rgba
}

// LoadObj 顶点  3    纹理  2   法向量  3
func LoadObj(name string) []float32 {
	// 顶点  3    纹理  2   法向量  3
	scanner := bufio.NewScanner(OpenFile(name))
	fs := make([]float32, 0)
	vs := make([][]float32, 0)
	vts := make([][]float32, 0)
	vns := make([][]float32, 0)
	for scanner.Scan() {
		temp := scanner.Text()
		if len(temp) < 3 {
			continue
		}
		switch temp[:2] {
		case "f ":
			items := parseItems(temp)
			for _, item := range items {
				indexs := strings.Split(item, "/")
				fs = append(fs, vs[ToInt(indexs[0])-1]...)
				fs = append(fs, vts[ToInt(indexs[1])-1]...)
				fs = append(fs, vns[ToInt(indexs[2])-1]...)
			}
		case "v ":
			items := parseItems(temp)
			vs = append(vs, []float32{ToFloat(items[0]), ToFloat(items[1]), ToFloat(items[2])})
		case "vt":
			items := parseItems(temp)
			vts = append(vts, []float32{ToFloat(items[0]), ToFloat(items[1])})
		case "vn":
			items := parseItems(temp)
			vns = append(vns, []float32{ToFloat(items[0]), ToFloat(items[1]), ToFloat(items[2])})
		}
	}
	return fs
}

func parseItems(temp string) []string {
	items := strings.Split(temp[2:], " ")
	res := make([]string, 0)
	for _, item := range items {
		temp = strings.TrimSpace(item)
		if len(temp) > 0 {
			res = append(res, temp)
		}
	}
	return res
}

func NewTextureBuff(width, height int32) (uint32, uint32) {
	var frameBuff, textureBuff uint32
	gl.GenFramebuffers(1, &frameBuff)
	gl.BindFramebuffer(gl.FRAMEBUFFER, frameBuff)
	// 颜色附件
	gl.GenTextures(1, &textureBuff)
	gl.BindTexture(gl.TEXTURE_2D, textureBuff)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, width, height, 0, gl.RGBA,
		gl.UNSIGNED_BYTE, nil)
	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, textureBuff, 0)
	// 深度信息 与模版信息 附件
	var rbo uint32
	gl.GenRenderbuffers(1, &rbo)
	gl.BindBuffer(gl.RENDERBUFFER, rbo)
	gl.RenderbufferStorage(gl.RENDERBUFFER, gl.DEPTH24_STENCIL8, width, height)
	gl.FramebufferRenderbuffer(gl.FRAMEBUFFER, gl.DEPTH_STENCIL_ATTACHMENT, gl.RENDERBUFFER, rbo)
	// 检查是否准备完毕
	if gl.CheckFramebufferStatus(gl.FRAMEBUFFER) != gl.FRAMEBUFFER_COMPLETE {
		panic("FRAMEBUFFER ERR")
	}
	return frameBuff, textureBuff
}

func NewDepTextureBuff(width, height int32) (uint32, uint32) {
	var frameBuff, textureBuff uint32
	// 深度贴图
	gl.GenTextures(1, &textureBuff)
	gl.BindTexture(gl.TEXTURE_2D, textureBuff)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.DEPTH_COMPONENT, width, height, 0, gl.DEPTH_COMPONENT, gl.FLOAT, nil)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	// 设置深度贴图
	gl.GenFramebuffers(1, &frameBuff)
	gl.BindFramebuffer(gl.FRAMEBUFFER, frameBuff)
	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.DEPTH_ATTACHMENT, gl.TEXTURE_2D, textureBuff, 0)
	// 取消绘制
	gl.DrawBuffer(gl.NONE)
	gl.ReadBuffer(gl.NONE)
	return frameBuff, textureBuff
}
