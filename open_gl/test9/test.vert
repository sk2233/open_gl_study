#version 410 core
layout (location = 0) in vec3 pos;

out vec3 tempPos;

uniform mat4 aModel;
uniform mat4 aView;
uniform mat4 aProjection;

void main()
{
    tempPos = pos;
    gl_Position = aProjection *aView * vec4(pos, 1.0);
}

