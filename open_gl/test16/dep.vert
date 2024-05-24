#version 410 core
layout (location = 0) in vec3 iPos;
layout (location = 1) in vec3 iNor;
layout (location = 2) in vec2 iTex;

uniform mat4 uModel;
uniform mat4 uView;
uniform mat4 uPro;

//out vec2 vTex;

void main()
{
//    vTex = iTex;
    gl_Position = uPro*uView * uModel * vec4(iPos, 1.0);
}
