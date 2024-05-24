#version 410 core
layout (location = 0) in vec3 iPos;
layout (location = 1) in vec2 iTex;

uniform mat4 uModel;
uniform mat4 uView;
uniform mat4 uProjection;

uniform mat4 uLastView;

out vec2 vTex;
out vec4 vPos;
out vec4 vLastPos;
out vec4 vWorldPos;

void main()
{
    vTex=iTex;
    gl_Position = uProjection *uView * uModel * vec4(iPos, 1.0);
    vPos = gl_Position;
    vLastPos = uProjection *uLastView * uModel * vec4(iPos, 1.0);
    vWorldPos = uModel * vec4(iPos, 1.0);
}
