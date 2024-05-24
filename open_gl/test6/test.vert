#version 410 core
layout (location = 0) in vec3 pos;
layout (location = 1) in vec3 normal;
layout (location = 2) in vec2 location;

uniform mat4 aModel;
uniform mat4 aView;
uniform mat4 aProjection;

out vec2 tempLoc;
out vec3 tempNormal;
out vec3 tempPos;

void main()
{
    tempLoc = location;
    tempNormal = normal;
    tempPos = vec3(aModel *vec4(pos.xyz, 1.0));// 移除一维
    gl_Position = aProjection *aView * aModel * vec4(pos.xyz, 1.0);
}
