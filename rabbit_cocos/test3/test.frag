#version 410 core

layout (location = 0)  out vec4 FragColor;
layout (location = 1)  out vec4 DepColor;

in vec2 vTex;
in float vDepth;

uniform sampler2D uTexture; // 默认使用0

void main()
{
    FragColor = texture(uTexture,vTex);
    DepColor = vec4(vDepth);
}