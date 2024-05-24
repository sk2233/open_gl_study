#version 410 core
layout (location = 0) in vec3 iPos;
layout (location = 1) in vec2 iTex;
layout (location = 2) in vec3 iNor;

uniform mat4 uModel;
uniform mat4 uView;
uniform mat4 uProjection;

out vec2 vTex;
out vec3 vNor;
out float vDep;

void main()
{
    vTex=iTex;
    vNor =iNor;
    vec4 pos = uProjection * uView * uModel * vec4(iPos, 1.0);
    vDep = pos.z/pos.w;
    gl_Position =pos;
}
