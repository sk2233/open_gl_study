#version 410 core

out vec4 FragColor;

in vec2 vTex;
in vec3 vNormal;
in vec3 vPos;

uniform sampler2D uTexture; // 默认使用0
uniform vec3 uLightPos;
uniform vec3 uEyePos;

void main()
{
    float value1 = 0.3;// 全局环境光
    vec3 lightDir = uLightPos-vPos;
    float value2 = max(dot(normalize(vNormal),normalize(lightDir)),0.0);
    vec3 eyeDir = uEyePos-vPos;
    vec3 midDir = normalize(lightDir)+normalize(eyeDir);
    float value3 = max(dot(normalize(midDir),normalize(vNormal)),0.0);
//    FragColor.r=texture(uTexture,vec2(value1)).r;
//    FragColor.g=texture(uTexture,vec2(value2)).g;
//    FragColor.b=texture(uTexture,vec2(value3)).b;
    FragColor.r=value1;
    FragColor.g=value2;
    FragColor.b=value3;
}