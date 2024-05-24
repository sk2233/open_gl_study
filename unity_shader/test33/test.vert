#version 410 core
layout (location = 0) in vec3 iPos;
layout (location = 1) in vec2 iTex;

uniform mat4 uModel;
uniform mat4 uView;
uniform mat4 uProjection;
uniform float uTime;

out vec2 vTex;

void main()
{
    vTex = iTex;
    vec3 pos=iPos.xyz+vec3(0,cos(iPos.x*4+uTime),0);
    gl_Position = uProjection *uView * uModel * vec4(pos, 1.0);
}
