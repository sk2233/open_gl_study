#version 410 core
layout (location = 0) in vec3 iPos;
layout (location = 1) in vec2 iTex;

uniform mat4 uView;
uniform mat4 uPro;

out vec2 vTex;
out vec3 vPos;

void main()
{
    vTex = iTex;
    vPos = iPos;
    gl_Position = uPro*uView * vec4(iPos, 1.0);
}
