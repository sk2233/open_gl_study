#version 410 core

out vec4 FragColor;

in vec2 vTex;
in vec3 vNor;

uniform sampler2D uTexture; // 默认使用0
uniform sampler2D uNoise;

uniform float uRate;

void main()
{
    if(texture(uNoise,vTex).r<uRate){
        discard;
    }
    FragColor = texture(uTexture,vTex);
}