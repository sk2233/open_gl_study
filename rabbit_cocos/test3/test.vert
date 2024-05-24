#version 410 core
layout (location = 0) in vec3 iPos;
layout (location = 1) in vec2 iTex;

uniform mat4 uModel;
uniform mat4 uView;
uniform mat4 uProjection;

out vec2 vTex;
out float vDepth;

void main()
{
    vTex = iTex;
    gl_Position = uProjection *uView * uModel * vec4(iPos.xyz, 1.0);
    vDepth = gl_Position.z/gl_Position.w;
}
