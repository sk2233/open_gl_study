#version 410 core

out vec4 FragColor;

in vec2 vTex;
uniform sampler2D uImage;
uniform sampler2D uNoise;
uniform float uTime;

void main()
{
   float threshold = float(fract(uTime*0.1));
   float value = texture(uNoise,vTex).r;
   value = smoothstep(threshold-0.1,value,threshold+0.1);
   FragColor=texture(uImage,vTex)*value;
}