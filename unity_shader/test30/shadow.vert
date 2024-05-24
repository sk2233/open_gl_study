#version 410 core
layout (location = 0) in vec3 iPos;

uniform mat4 uModel;
uniform mat4 uView;
uniform mat4 uProjection;

uniform mat4 uLight;  // 光场的 P * V

out vec4 vPosInLight;

void main()
{
    gl_Position = uProjection * uView * uModel * vec4(iPos, 1.0);
    // 获取在广场下的位置
    vPosInLight = uLight * uModel * vec4(iPos,1.0);
}