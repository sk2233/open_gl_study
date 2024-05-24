#version 410 core

out vec4 FragColor;

in vec3 tempPos;

uniform samplerCube aTexture; // 默认使用0

void main()
{
    FragColor =texture(aTexture,tempPos);
}