#version 410 core
layout (location = 0) in vec3 iPos;
layout (location = 1) in vec2 iTex;
layout (location = 2) in vec3 iNormal;

uniform mat4 uModel;
uniform mat4 uView;
uniform mat4 uProjection;

out vec2 vTex;
out vec3 vNormal;
out vec3 vPos;

void main()
{
    vTex = iTex;
    vNormal = vec3(uModel * vec4(iNormal, 1.0));// 保证法线的正确性
    vPos = vec3(uModel * vec4(iPos, 1.0));
    gl_Position = uProjection *uView * uModel * vec4(iPos.xyz, 1.0);
}
