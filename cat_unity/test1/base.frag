#version 410 core

out vec4 FragColor;

in vec2 vTex;
uniform sampler2D uImage;
uniform float uTime;
uniform vec4 uCol1 = vec4(1,0,0,1);
uniform vec4 uCol2 = vec4(0,1,0,1);

void main()
{
   vec2 tex = vec2(vTex.x+float(sin(vTex.y+uTime*5)*0.1),vTex.y);
   FragColor=texture(uImage,tex);
   vec4 col = uCol1*vTex.y+uCol2*(1-vTex.y);
   FragColor = FragColor *col;
}