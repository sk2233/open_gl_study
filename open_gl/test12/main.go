///*
//@author: sk
//@date: 2023/5/14
//*/
package main

import (
	"log"
	frame "openGL/frame2"
	"openGL/utils"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	width  = 800
	height = 600
)

var (
	camera = frame.NewCamera()
)

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	window, err := glfw.CreateWindow(width, height, "Game", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}
	gl.Viewport(0, 0, int32(width), int32(height))
	gl.Enable(gl.DEPTH_TEST)
	skyShader := frame.NewShader("test12/sky")
	// cube VAO
	var cubeVAO, cubeVBO uint32
	gl.GenVertexArrays(1, &cubeVAO)
	gl.GenBuffers(1, &cubeVBO)
	gl.BindVertexArray(cubeVAO)
	gl.BindBuffer(gl.ARRAY_BUFFER, cubeVBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(cubeVs)*4, gl.Ptr(cubeVs), gl.STATIC_DRAW)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 5*4, 0)
	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointerWithOffset(1, 2, gl.FLOAT, false, 5*4, 3*4)
	// skybox VAO
	var skyboxVAO, skyboxVBO uint32
	gl.GenVertexArrays(1, &skyboxVAO)
	gl.GenBuffers(1, &skyboxVBO)
	gl.BindVertexArray(skyboxVAO)
	gl.BindBuffer(gl.ARRAY_BUFFER, skyboxVBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(skyVs)*4, gl.Ptr(skyVs), gl.STATIC_DRAW)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 3*4, 0)
	sky := utils.LoadCubeTexture("test11/sky/back.jpg", "test11/sky/bottom.jpg", "test11/sky/front.jpg",
		"test11/sky/left.jpg", "test11/sky/right.jpg", "test11/sky/top.jpg")
	skyShader.Use()
	skyShader.Set1i("skybox", 0)
	for !window.ShouldClose() {
		gl.DepthFunc(gl.LEQUAL) // change depth function so depth test passes when values are equal to depth buffer's content
		skyShader.Use()
		view := camera.GetView()
		skyShader.SetMat4("view", &view[0])
		projection := utils.GetDefaultPerspective()
		skyShader.SetMat4("projection", &projection[0])
		// skybox cube
		gl.BindVertexArray(skyboxVAO)
		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_CUBE_MAP, sky)
		gl.DrawArrays(gl.TRIANGLES, 0, 36)
		gl.BindVertexArray(0)
		gl.DepthFunc(gl.LESS)

		window.SwapBuffers()
		glfw.PollEvents()
	}
}

// loads a cubemap texture from 6 individual texture faces
// order:
// +X (right)
// -X (left)
// +Y (top)
// -Y (bottom)
// +Z (front)
// -Z (back)
//// -------------------------------------------------------
//unsigned int loadCubemap(vector<std::string> faces)
//{
//unsigned int textureID;
//glGenTextures(1, &textureID);
//glBindTexture(GL_TEXTURE_CUBE_MAP, textureID);
//
//int width, height, nrChannels;
//for (unsigned int i = 0; i < faces.size(); i++)
//{
//unsigned char *data = stbi_load(faces[i].c_str(), &width, &height, &nrChannels, 0);
//if (data)
//{
//glTexImage2D(GL_TEXTURE_CUBE_MAP_POSITIVE_X + i, 0, GL_RGB, width, height, 0, GL_RGB, GL_UNSIGNED_BYTE, data);
//stbi_image_free(data);
//}
//else
//{
//std::cout << "Cubemap texture failed to load at path: " << faces[i] << std::endl;
//stbi_image_free(data);
//}
//}
//glTexParameteri(GL_TEXTURE_CUBE_MAP, GL_TEXTURE_MIN_FILTER, GL_LINEAR);
//glTexParameteri(GL_TEXTURE_CUBE_MAP, GL_TEXTURE_MAG_FILTER, GL_LINEAR);
//glTexParameteri(GL_TEXTURE_CUBE_MAP, GL_TEXTURE_WRAP_S, GL_CLAMP_TO_EDGE);
//glTexParameteri(GL_TEXTURE_CUBE_MAP, GL_TEXTURE_WRAP_T, GL_CLAMP_TO_EDGE);
//glTexParameteri(GL_TEXTURE_CUBE_MAP, GL_TEXTURE_WRAP_R, GL_CLAMP_TO_EDGE);
//
//return textureID;
//}
