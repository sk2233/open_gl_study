#version 410 core
layout (location = 0) out vec4 FragColor;
//layout (location = 1) out vec4 DepColor;

in vec2 vTex;
in vec2 vPos;

uniform sampler2D uImage;
uniform sampler2D uMask;
uniform vec2 uPos;

void main()
{
    FragColor = texture(uImage,vTex);
    float rate = max(1.0-length(vPos-uPos),0.0);
    FragColor.a*=rate;
}