#version 410 core
layout (location = 0) in vec4 iPos;
layout (location = 1) in vec4 iCol;
layout (location = 2) in vec2 iOff;

out vec4 vCol;

void main()
{
   vCol = iCol;
    gl_Position = iPos+vec4(iOff,0.0,0.0);
}