#version 410 core
layout (location = 0) in vec3 iPos;
layout (location = 1) in vec3 iNor;
layout (location = 2) in vec2 iTex;

uniform mat4 uModel;
uniform mat4 uView;
uniform mat4 uProjection;

out vec2 vTex;
out vec3 vNor;

void main()
{
    vTex = iTex;
    vNor = iNor;
    gl_Position = uProjection *uView * uModel * vec4(iPos.xyz, 1.0);
}
