#version 410 core

out vec4 FragColor;

in vec2 vTex;

uniform sampler2D uImage;// 直接使用 0即可
uniform vec4 uColor;

void main()
{
   FragColor=uColor * texture(uImage,vTex);
}