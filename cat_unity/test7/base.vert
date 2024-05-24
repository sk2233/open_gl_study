#version 410 core
layout (location = 0) in vec3 iPos;
layout (location = 1) in vec2 iTex;

out vec2 vTex;

uniform mat4 uProjection;
uniform mat4 uModel;

void main()
{
    vTex = iTex;
    gl_Position =uProjection * uModel * vec4(iPos,1);
}
