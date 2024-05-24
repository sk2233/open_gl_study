#version 410 core
layout (location = 0) out vec4 FragColor;
//layout (location = 1) out vec4 DepColor;

in vec2 vTex;
in vec2 vPos;

uniform sampler2D uImage;
uniform sampler2D uMask;
uniform vec2 uPos;

void main()
{
    FragColor = texture(uImage,vTex);
    vec2 pos = vPos-uPos;
    if(pos.x>0&&pos.y>0&&pos.x<2&&pos.y<2){
        FragColor = texture(uMask,pos/2.0);
    }
}