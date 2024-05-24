#version 410 core

out vec4 FragColor;

in vec2 vTex;

uniform sampler2D uTexture;

void main()
{
    FragColor=texture(uTexture,vTex);
}