#version 410 core

out vec4 FragColor;

in vec2 vTex;
in vec3 vPos;

uniform sampler2D uImage;
uniform sampler2D uNormal;
uniform sampler2D uHeight;
uniform vec3 uLightPos;
uniform vec3 uEyePos;

void main()
{
    vec3 lightDir = normalize(uLightPos-vPos);
    vec3 eyeDir = normalize(uEyePos-vPos);
    // 获取高度进行纹理 法线贴图偏移
    float height = texture(uHeight,vTex).r;
    vec2 tex=vTex+eyeDir.xy / eyeDir.z * height;

    vec3 midDir = normalize(lightDir+eyeDir);
    vec3 normal = texture(uNormal,tex).xyz*2-1;
    float value = pow(dot(midDir,normal),32);

    FragColor = texture(uImage,tex)+vec4(0,value,0,0);
}