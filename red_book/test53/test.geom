#version 410 core
layout (triangles) in;
layout (triangle_strip, max_vertices = 3) out;

out vec2 gTex;
in vec2[] vTex;

void main()
{
    for(int i=0;i<3;i++){
        gl_Position = gl_in[i].gl_Position;
        gTex = vTex[i];
        EmitVertex();
    }
    EndPrimitive();
}