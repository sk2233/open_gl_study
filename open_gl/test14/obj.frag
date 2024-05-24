#version 410 core

out vec4 FragColor;

in vec2 gLoc;

uniform sampler2D uTexture; // 默认使用0

void main()
{
   FragColor=texture(uTexture,gLoc);
}