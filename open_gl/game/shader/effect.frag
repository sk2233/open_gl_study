#version 410 core

out vec4 FragColor;

in vec2 vTex;

uniform sampler2D uImage;// 直接使用 0即可

void main()
{
   // 颜色反转
   FragColor= vec4(1-texture(uImage,vTex).xyz,1);
}