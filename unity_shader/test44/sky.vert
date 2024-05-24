#version 410 core
layout (location = 0) in vec3 iPos;

out vec3 vTex;

uniform mat4 uProjection;
uniform mat4 uView;

void main()
{
    vTex = iPos;
    // 天空盒 本来 就位于 世界原点
    vec4 pos = uProjection * uView * vec4(iPos, 1.0);
    gl_Position = pos.xyzz;// 保证 最后渲染
    // gl_Position 最终会转换到齐次空间
}