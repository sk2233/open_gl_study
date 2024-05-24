#version 410 core
out vec4 FragColor;

in vec3 gCol;

void main()
{
   FragColor = vec4(gCol.rgb,1.0);
}