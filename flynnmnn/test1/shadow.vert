#version 410 core
layout (location = 0) in vec3 iPos;
layout (location = 1) in vec2 iTex;

uniform mat4 uModel;
uniform mat4 uView;
uniform mat4 uProjection;

out vec2 vTex;
out vec2 vPos;

void main()
{
    vTex = iTex;
    gl_Position = uProjection * uView * uModel * vec4(iPos, 1.0);
    vec4 pos = uModel * vec4(iPos, 1.0);
    vPos = pos.xz;
}