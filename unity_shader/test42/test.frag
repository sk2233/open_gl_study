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
    // 高光 不再渐变  采用截断的方式  中间 过渡防止 锯齿太生硬
    float value3 = smoothstep(0.49,0.51,dot(normalize(midDir),normalize(vNormal)));
    FragColor = texture(uTexture,vTex);
    FragColor.r*=value1;
    FragColor.g*=value2;
    FragColor.b*=value3;
}