#version 410 core
layout (location = 0) in vec3 iPos;
layout (location = 1) in vec2 iTex;
layout (location = 2) in vec3 iNor;

uniform mat4 uModel;
uniform mat4 uModelIT;
uniform mat4 uView;
uniform mat4 uProjection;

out vec2 vTex;
out vec3 vPos;
out vec3 vNor;

void main()
{
    vTex = iTex;
    gl_Position = uProjection *uView * uModel * vec4(iPos.xyz, 1.0);
    vPos = vec3(uModel * vec4(iPos, 1.0));
    vNor = vec3(uModelIT * vec4(iNor,1.0));
}
