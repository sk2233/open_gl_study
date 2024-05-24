#version 410 core

out vec4 FragColor;

in vec2 vTex;

uniform sampler2D uTexture; // 默认使用0
uniform sampler2D uNoise;

uniform float uThreshold;

void main()
{
    vec4 noise = texture(uNoise,vTex);
    if(noise.r<fract(uThreshold/5.0)){
        discard;
    }
    FragColor = texture(uTexture,vTex);
}