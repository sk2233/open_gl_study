#version 410 core
layout (location = 0) in vec3 aPos;
layout (location = 1) in vec2 aTexCoord;
//out vec4 tempColor;

out vec2 TexCoord;
uniform mat4 transform;

void main()
{
    gl_Position = transform * vec4(aPos.x, aPos.y, aPos.z, 1.0);
    TexCoord = aTexCoord;
//    tempColor = vec4(aPos.xyz,1.0);
}
