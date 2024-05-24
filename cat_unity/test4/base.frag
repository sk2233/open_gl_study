#version 410 core

out vec4 FragColor;

in vec2 vTex;
uniform sampler2D uImage;
uniform float uTime;

void main()
{
   float count = float(floor(uTime));
   vec2 tex = vec2(floor(vTex*vec2(count)))/vec2(count);
   FragColor=texture(uImage,tex);
}