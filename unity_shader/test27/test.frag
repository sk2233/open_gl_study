#version 410 core

out vec4 FragColor;

in vec2 vTex;
in vec3 vPos;
in mat4 vModel;

uniform sampler2D uTexture; // 默认使用0
uniform sampler2D uNormal;
uniform vec3 uLightPos;
uniform vec3 uEyePos;


void main()
{
    float value1 = 0.3;// 全局环境光
    vec3 lightDir = uLightPos-vPos;
    // 注意法线  也必须变化到世界坐标系
    vec3 normal = (vModel *texture(uNormal,vTex)).xyz;
    float value2 = max(dot(normalize(normal),normalize(lightDir)),0.0);
    vec3 eyeDir = uEyePos-vPos;
    vec3 midDir = normalize(lightDir)+normalize(eyeDir);
    float value3 = pow(max(dot(normalize(midDir),normalize(normal)),0.0),32.0);
    FragColor=texture(uTexture,vTex)*(value1+value2+value3);
}