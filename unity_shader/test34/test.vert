#version 410 core
layout (location = 0) in vec3 iPos;
layout (location = 1) in vec2 iTex;

uniform mat4 uModel;
uniform mat4 uView;
uniform mat4 uProjection;
uniform vec3 uEye;// 必须使用模型坐标系下的 位置 在外面转换好

out vec2 vTex;

void main()
{
    vec3 center = vec3(0,0,0);
    vec3 zDir = normalize(uEye-center);
    vec3 xDir = normalize(cross(vec3(0,1,0),zDir));
    vec3 yDir = normalize(cross(zDir,xDir));
    vec3 pos = iPos.xyz;
    pos = xDir*pos.x+yDir*pos.y+zDir*pos.z;
    vTex = iTex;
    gl_Position = uProjection *uView * uModel * vec4(pos, 1.0);
}
