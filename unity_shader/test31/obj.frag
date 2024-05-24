#version 410 core
out vec4 FragColor;

in vec3 vNor;
in vec3 vPos;
in vec2 vTex;

uniform vec3 uEyePos;
uniform samplerCube uSkybox;
uniform sampler2D uTex;

void main()
{
    float ratio = 1.0 / 1.5;
    vec3 I = vPos - uEyePos;
    // 反射
    vec3 R1 =  reflect(I,vNor);
    vec4 V1 =  texture(uSkybox,R1);
    // 折射
    vec3 R2 = refract(I, vNor, ratio);
    vec4 V2 = texture(uSkybox,R2);
    // 菲尼尔现象 反射折射都存在  按比例分割
    float rate = pow(dot(normalize(uEyePos-vPos),normalize(vNor)),5.0);
    FragColor = (V2*rate+V1*(1-rate))*texture(uTex,vTex);
}