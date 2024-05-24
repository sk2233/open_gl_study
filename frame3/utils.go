/*
@author: sk
@date: 2023/5/27
*/
package frame3

import (
	"image"
	"image/draw"
	_ "image/png"
	"os"

	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	BasePath = "/Users/bytedance/Documents/go/openGL/"
)

func ReadAll(name string) []byte {
	bs, err := os.ReadFile(BasePath + name)
	HandleErr(err)
	return bs
}

func Has(name string) bool {
	stat, err := os.Stat(BasePath + name)
	return err == nil && stat != nil
}

func OpenFile(name string) *os.File {
	file, err := os.Open(BasePath + name)
	HandleErr(err)
	return file
}

func HandleErr(err error) {
	if err != nil {
		panic(err)
	}
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

func GetAxle(window *glfw.Window, small, big glfw.Key) float32 {
	if window.GetKey(small) == glfw.Press {
		return -1
	}
	if window.GetKey(big) == glfw.Press {
		return 1
	}
	return 0
}

func CollisionPos(rect IRect, pos complex64) bool {
	return VecGt(pos, rect.GetPos()) && VecLt(pos, rect.GetPos()+rect.GetSize())
}

func CollisionRect(rect1, rect2 IRect) bool {
	if !VecGt(rect1.GetPos()+rect1.GetSize(), rect2.GetPos()) {
		return false
	}
	if !VecGt(rect2.GetPos()+rect2.GetSize(), rect1.GetPos()) {
		return false
	}
	return true
}

func VecLt(pos1, pos2 complex64) bool {
	return real(pos1) < real(pos2) && imag(pos1) < imag(pos2)
}

func VecGt(pos1, pos2 complex64) bool {
	return real(pos1) > real(pos2) && imag(pos1) > imag(pos2)
}

func VecMul(vec1, vec2 complex64) complex64 {
	return complex(real(vec1)*real(vec2), imag(vec1)*imag(vec2))
}

func Vec2Float(vec complex64) (float32, float32) {
	return real(vec), imag(vec)
}

func CreateDefaultShader() *Shader {
	return NewShader("frame3/shader/base")
}
