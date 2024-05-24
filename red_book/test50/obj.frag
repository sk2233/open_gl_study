#version 410 core
out vec4 FragColor;

in vec2 vTex;

uniform sampler3D uImage;

void main()
{
   FragColor = texture(uImage,vec3(vTex,vTex.x));
}