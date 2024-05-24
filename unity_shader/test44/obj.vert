#version 410 core
layout (location = 0) in vec3 iPos;
layout (location = 1) in vec3 iNor;
layout (location = 2) in vec2 iTex;

out vec2 vTex;
out vec3 vPos;
out vec3 vNor;

uniform mat4 uModel;
uniform mat4 uView;
uniform mat4 uProjection;

void main()
{
    vTex = iTex;
    vNor = iNor;
    vPos = vec3(uModel * vec4(iPos, 1.0));
    gl_Position = uProjection * uView * uModel * vec4(iPos, 1.0);
}