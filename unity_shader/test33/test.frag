#version 410 core

out vec4 FragColor;

in vec2 vTex;

uniform sampler2D uTexture;
uniform vec2 uOffset;

void main()
{
    FragColor=texture(uTexture,vTex/8+uOffset);
}