#version 410 core

out vec4 FragColor;

in vec2 vTex;
// 当前点的位置
in vec4 vPos;
// 上一个点的位置
in vec4 vLastPos;
in vec4 vWorldPos;

uniform sampler2D uTexture;
uniform sampler2D uNoise;

void main()
{
    vec4 color = texture(uTexture,vTex);
    float rate=clamp(vWorldPos.y/3,0.0,1.0)* texture(uNoise,vTex).r;
     rate=clamp(vWorldPos.y/3,0.0,1.0);
    // 高度雾
    FragColor = color*rate+vec4(0,1,0,1)*(1-rate);
}