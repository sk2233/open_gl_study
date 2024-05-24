#version 410 core
layout (location = 0) out vec4 FragColor;
layout (location = 1) out vec4 NorDep;

in vec2 vTex;
in vec3 vNor;
in float vDep;

uniform sampler2D uTexture;

void main()
{
    FragColor = texture(uTexture,vTex);
    NorDep = vec4(vDep,vNor.xyz);
}