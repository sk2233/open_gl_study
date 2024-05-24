#version 410 core
layout (location = 0) in vec3 iPos;

uniform mat4 uModel;
uniform mat4 uView;
uniform mat4 uProjection;

out vec3 vPos;

void main()
{
    vPos = vec3(uModel * vec4(iPos,1.0));
    gl_Position = uProjection * uView * uModel * vec4(iPos, 1.0);
}