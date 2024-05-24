#version 410 core

out vec4 FragColor;

in vec2 vTex;
in vec3 vNormal;
in vec3 vPos;

uniform sampler2D uTexture; // 默认使用0
uniform vec3 uLightPos;
uniform vec3 uLightDir;// 模拟聚光灯需要
uniform vec3 uEyePos;
uniform float uThreshold;

void main()
{
    float value1 = 0.3;// 全局环境光
    vec3 lightDir = uLightPos-vPos;
    float value2 = max(dot(normalize(vNormal),normalize(lightDir)),0.0);
    vec3 eyeDir = uEyePos-vPos;
    vec3 midDir = normalize(lightDir)+normalize(eyeDir);
    float value3 = max(dot(normalize(midDir),normalize(vNormal)),0.0);
    FragColor.r=value1;
    FragColor.g=value2;
    FragColor.b=value3;
    if(dot(normalize(uLightDir),normalize(vPos-uLightPos))>uThreshold){
        FragColor.r=1;
    }
}