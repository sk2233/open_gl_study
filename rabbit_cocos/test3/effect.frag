#version 410 core

out vec4 FragColor;

in vec2 vTex;

uniform sampler2D uTexture; // 默认使用0
uniform sampler2D uDep;
uniform float uNear = 0.6;
uniform float uFar = 0.8;
uniform vec4 uCol = vec4(1.0,0.0,0.0,1.0);

void main()
{
    float rate = texture(uDep,vTex).r;
    rate = smoothstep(uNear,rate,uFar);
    vec4 col =  texture(uTexture,vTex);
    FragColor = rate*col+(1-rate)*uCol;
}