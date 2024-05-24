#version 410 core
in vec4 gPos;

uniform vec3 uLightPos;
uniform float uFar;

void main()
{
    // 获取与光源的距离
    float lightDistance = length(gPos.xyz - uLightPos);
    // 变到 0 ～ 1  并写入深度值
    gl_FragDepth = lightDistance / uFar;
}