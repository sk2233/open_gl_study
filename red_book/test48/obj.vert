#version 410 core
layout (location = 0) in vec3 iPos;
layout (location = 1) in vec3 iCol;

out vec3 vCol;

void main()
{
    vCol = iCol;
    gl_Position = vec4(iPos.xyz,1.0);
}