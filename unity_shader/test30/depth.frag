#version 410 core

in vec2 vTex;

uniform int uOpen;
uniform sampler2D uTexture;

void main()
{
    if(uOpen==1&&texture(uTexture,vTex).a<0.5){
        discard;
    }
}