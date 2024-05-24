#version 410 core
layout (points) in;
layout (triangle_strip, max_vertices = 16) out;

in vec3[] vCol;
out vec3 gCol;

uniform mat4 uModel;
uniform mat4 uView;
uniform mat4 uProjection;

void main()
{
    vec4 pos = gl_in[0].gl_Position;
    mat4 sum = uProjection * uView * uModel;
    gCol = vCol[0];
    gl_Position = sum * vec4(pos + vec4(1.0, 0.0, 1.0, 0.0));
    EmitVertex();
    gCol = vCol[0];
    gl_Position = sum * vec4(pos + vec4(0.0, 0.0, 1.0, 0.0));
    EmitVertex();
    gCol = vCol[0];
    gl_Position = sum * vec4(pos + vec4(1.0, 1.0, 1.0, 0.0));
    EmitVertex();
    gCol = vCol[0];
    gl_Position = sum * vec4(pos + vec4(0.0, 1.0, 1.0, 0.0));
    EmitVertex();
    gCol = vCol[0];
    gl_Position = sum * vec4(pos + vec4(1.0, 1.0, 0.0, 0.0));
    EmitVertex();
    gCol = vCol[0];
    gl_Position = sum * vec4(pos + vec4(0.0, 1.0, 0.0, 0.0));
    EmitVertex();
    gCol = vCol[0];
    gl_Position = sum * vec4(pos + vec4(1.0, 0.0, 0.0, 0.0));
    EmitVertex();
    gCol = vCol[0];
    gl_Position = sum * pos;
    EmitVertex();
    EndPrimitive();
    gCol = vCol[0];
    gl_Position = sum * vec4(pos + vec4(1.0, 1.0, 1.0, 0.0));
    EmitVertex();
    gCol = vCol[0];
    gl_Position = sum * vec4(pos + vec4(1.0, 1.0, 0.0, 0.0));
    EmitVertex();
    gCol = vCol[0];
    gl_Position = sum * vec4(pos + vec4(1.0, 0.0, 1.0, 0.0));
    EmitVertex();
    gCol = vCol[0];
    gl_Position = sum * vec4(pos + vec4(1.0, 0.0, 0.0, 0.0));
    EmitVertex();
    gCol = vCol[0];
    gl_Position = sum * vec4(pos + vec4(0.0, 0.0, 1.0, 0.0));
    EmitVertex();
    gCol = vCol[0];
    gl_Position = sum * pos;
    EmitVertex();
    gCol = vCol[0];
    gl_Position = sum * vec4(pos + vec4(0.0, 1.0, 1.0, 0.0));
    EmitVertex();
    gCol = vCol[0];
    gl_Position = sum * vec4(pos + vec4(0.0, 1.0, 0.0, 0.0));
    EmitVertex();
    EndPrimitive();
}