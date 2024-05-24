#version 410 core
layout (location = 0) in vec3 pos;
layout (location = 1) in vec2 location;

uniform mat4 aModel;
uniform mat4 aView;
uniform mat4 aProjection;

out vec2 tempLoc;

void main()
{
    tempLoc = location;
    gl_Position = aProjection *aView * aModel * vec4(pos.xyz, 1.0);
}
