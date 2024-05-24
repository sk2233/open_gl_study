#version 410 core
layout (location = 0) in vec3 iPos;
layout (location = 1) in vec3 iNor;

uniform mat4 uModel;
uniform mat4 uView;
uniform mat4 uProjection;

out vec3 vNor;
out vec4 vPos;

void main()
{
    vNor = vec3(uModel*vec4(iNor,1.0));
    gl_Position = uProjection *uView * uModel * vec4(iPos, 1.0);
    vPos = gl_Position;
}
