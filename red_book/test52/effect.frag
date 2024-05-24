#version 410 core

out vec4 FragColor;

in vec2 vTex;

uniform sampler2D uTexture;
uniform float uTime;

void main()
{
    vec2 offset = vec2(1.0*cos(uTime)*6.0/1280.0,0.0);
    float clrR = texture(uTexture,vTex+offset).r;
    float clrG = texture(uTexture,vTex).g;
    float clrB = texture(uTexture,vTex-offset).b;
    FragColor = vec4(clrR,clrG,clrB,1.0);
}