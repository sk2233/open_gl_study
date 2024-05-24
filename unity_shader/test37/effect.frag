#version 410 core

out vec4 FragColor;

in vec2 vTex;

uniform sampler2D uImage0;
uniform sampler2D uImage1;
//uniform sampler2D uImage2;
//uniform sampler2D uImage3;
//uniform sampler2D uImage4;
uniform int uIndex;

void main()
{
    if(uIndex==0){
        FragColor= texture(uImage0,vTex);
    }else if(uIndex==1){
        FragColor= texture(uImage1,vTex);
    }
//    else if(uIndex==2){
//        FragColor= texture(uImage2,vTex);
//    }else if(uIndex==3){
//        FragColor= texture(uImage3,vTex);
//    }else if(uIndex==4){
//        FragColor= texture(uImage4,vTex);
//    }
//    FragColor= texture(uImage0,vTex);
}