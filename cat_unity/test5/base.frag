#version 410 core

out vec4 FragColor;

in vec2 vTex;
uniform sampler2D uImage;

void main()
{
   vec4 col = texture(uImage,vTex);
   vec3 lum = vec3(0.299, 0.587, 0.114);
   FragColor = vec4( vec3(dot( col.rgb, lum)), col.a);
}