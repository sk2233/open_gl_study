#version 410 core
layout (location = 0) in vec3 iPos;

uniform mat4 uModel;
uniform mat4 uView;
uniform mat4 uProjection;

void main()
{
    // 默认只写入了红色 通道
    gl_Position = uProjection * uView * uModel * vec4(iPos, 1.0);
}