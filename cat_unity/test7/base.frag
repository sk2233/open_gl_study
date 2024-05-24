#version 410 core

out vec4 FragColor;

in vec2 vTex;
uniform sampler2D uImage;
uniform float uAlpha = 1.0;

void main()
{
   vec4 col = texture(uImage,vTex);
   if (col.a > 0.0){
      FragColor = vec4(1.0,1.0,1.0,uAlpha);
   }
}