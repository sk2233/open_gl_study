#version 410 core
layout (location = 0) in vec4 iPos;
layout (location = 1) in vec4 iCol;

out vec4 vCol;

uniform mat4 uP;
uniform mat4 uV;
uniform mat4 uM;

void main()
{
   vCol = iCol;
    gl_Position = uP*uV*uM*iPos;
}