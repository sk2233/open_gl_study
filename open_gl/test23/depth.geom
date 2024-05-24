#version 410 core
layout (triangles) in;
layout (triangle_strip, max_vertices=18) out;

uniform mat4 uShadow[6]; // 6面  阴影的 P * V

out vec4 gPos; // 世界坐标系下的位置

void main()
{
    for(int face = 0; face < 6; face++)
    {
        gl_Layer = face; // 指定输出的面
        for(int i = 0; i < 3; i++) // 处理每个三角形顶点
        {
            gPos = gl_in[i].gl_Position;
            gl_Position = uShadow[face] * gPos;
            EmitVertex();// 储存一个点
        }
        EndPrimitive();//  储存一个三角形
    }
}