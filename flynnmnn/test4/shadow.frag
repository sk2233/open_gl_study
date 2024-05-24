#version 410 core
layout (location = 0) out vec4 DepColor;

in float vDep;

void main()
{
    DepColor = vec4(vDep);
}