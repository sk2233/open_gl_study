#version 410 core
out vec4 FragColor;

in vec3 vTex;

uniform samplerCube uSkybox;

void main()
{
    FragColor = texture(uSkybox,vTex);
}