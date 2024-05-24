#version 410 core
out vec4 FragColor;

in vec2 vTex;

uniform sampler2D uImage;

void main()
{
   FragColor = texture(uImage,gl_PointCoord);
}