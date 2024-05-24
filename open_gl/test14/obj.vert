#version 410 core
layout (location = 0) in vec3 iPos;
layout (location = 1) in vec2 iLoc;
layout (location = 2) in vec3 iNor;

uniform mat4 uModel;
uniform mat4 uView;
uniform mat4 uProjection;

out vec2 vLoc;
out vec3 vNor;

void main()
{
    vLoc = iLoc;
    vNor = iNor;
    gl_Position = uProjection *uView * uModel * vec4(iPos, 1.0);
}
