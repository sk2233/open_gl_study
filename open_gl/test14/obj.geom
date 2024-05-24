#version 410 core
layout (triangles) in;
layout (line_strip, max_vertices = 6) out;

in vec2 vLoc[];
in vec3 vNor[];

out vec2 gLoc;

void main() {
    for(int i=0;i<3;i++){
        gLoc = vLoc[i];
        gl_Position = gl_in[i].gl_Position;
        EmitVertex();
        gl_Position = gl_in[i].gl_Position+vec4(vNor[i],0);
        EmitVertex();
    }
    EndPrimitive();
}
