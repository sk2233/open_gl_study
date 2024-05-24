#version 410 core
layout (location = 0) in vec2 iPos;
layout (location = 1) in vec2 iTex;

out vec2 vTex;

uniform mat4 uModel;
uniform mat4 uProj;

void main()
{
    vTex=iTex;
    gl_Position =uProj* uModel * vec4(iPos,0,1);
}
