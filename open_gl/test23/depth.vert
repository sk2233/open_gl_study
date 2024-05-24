#version 410 core
layout (location = 0) in vec3 position;

uniform mat4 uModel;

void main()
{
    // P * V 在几何着色器中完成
    gl_Position = uModel * vec4(position, 1.0);
}