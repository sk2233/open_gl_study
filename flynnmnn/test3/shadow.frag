#version 410 core
layout (location = 0) out vec4 FragColor;
//layout (location = 1) out vec4 DepColor;

in vec2 vTex;
in vec2 vPos;

uniform sampler2D uImage;
uniform sampler2D uMask;
uniform vec2 uPos;
uniform float uLen;

void main()
{
    FragColor = texture(uImage,vTex);
    float rate =smoothstep(uLen,uLen+1, length(vPos-uPos));
    FragColor.r*=rate;
}