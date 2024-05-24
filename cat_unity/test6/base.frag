#version 410 core

out vec4 FragColor;

in vec2 vTex;
uniform sampler2D uImage;
uniform float uTime;

void main()
{
   vec2 tex = vTex;
   float w = 0.1*cos(uTime+0.05*vTex.y);
   tex.x += w*0.5*cos(0.04*(-vTex.y+2.0));
   FragColor = texture(uImage,tex);
}