#version 410 core

out vec4 FragColor;

in vec2 vTex;

void main()
{
   FragColor=vec4(gl_FragCoord.x/1280,gl_FragCoord.y/720,0,1);
   FragColor = vec4(vTex,0,1);
}