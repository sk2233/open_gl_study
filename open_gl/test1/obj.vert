#version 410 core
layout (location = 0) in vec2 iPos;
layout (location = 1) in vec3 iCol;

out vec3 vCol;

void main()
{
    vCol=iCol;
    gl_Position =vec4(iPos-vec2(gl_InstanceID/100.0-0.5),0,1);
}
