#version 410 core
layout (location = 0) in vec3 iPos;
layout (location = 1) in vec2 iTex;
layout (location = 2) in vec3 iNor;

out vec3 vNor;
out vec3 vPos;
out vec2 vTex;

uniform mat4 uModel;
uniform mat4 uView;
uniform mat4 uProjection;

void main()
{
    vNor = vec3(uModel * vec4(iNor, 1.0));
    vPos = vec3(uModel * vec4(iPos, 1.0));
    vTex = iTex;
    gl_Position = uProjection * uView * uModel * vec4(iPos, 1.0);
}