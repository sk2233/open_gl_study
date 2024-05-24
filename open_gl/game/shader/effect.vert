#version 410 core
layout (location = 0) in vec2 iPos;
layout (location = 1) in vec2 iTex;

out vec2 vTex;

uniform float uTime;

void main()
{
    vTex=iTex+vec2(sin(uTime/2),cos(uTime/2));
    gl_Position = vec4(iPos,0,1);
}
