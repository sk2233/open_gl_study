#version 410 core
out vec4 FragColor;

in vec2 vTex;
in vec3 vPos;
in vec3 vNor;

uniform vec3 uEyePos;
uniform samplerCube uSkybox;
// 还需要 法线贴图  没有素材
uniform sampler2D uNor;
uniform sampler2D uTex;

void main()
{
    float ratio = 1.0 / 1.5;
    vec3 I = vPos - uEyePos;
    // 法线方向不是世界坐标下的 也没有装换 有问题
    vec3 nor =vNor;
    // 反射
    vec3 R1 =  reflect(I,nor);
    vec4 V1 =  texture(uSkybox,R1);
    // 折射
    vec3 R2 = refract(I, nor, ratio);
    vec4 V2 = texture(uSkybox,R2);
    // 菲尼尔现象 反射折射都存在  按比例分割
    float rate = pow(dot(normalize(uEyePos-vPos),normalize(nor)),5.0);
    FragColor = (V2*rate+V1*(1-rate))*texture(uTex,vTex);
}