#version 410 core

out vec4 FragColor;

in vec2 vTex;
uniform sampler2D uImage;
uniform float uTime;

void main()
{
   float l = length(vTex*2-vec2(1.0));
   float value =(float(sin((l+uTime)*-10.0))+1.0)/20.0*(1-l);
   vec2 tex = vTex+vec2(value);
   FragColor=texture(uImage,tex);
}