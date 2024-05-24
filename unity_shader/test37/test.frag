#version 410 core

layout (location = 0) out vec4 FragColor0;
layout (location = 1) out vec4 FragColor1;
//layout (location = 2) out vec4 FragColor2;
//layout (location = 3) out vec4 FragColor3;
//layout (location = 4) out vec4 FragColor4;

in vec2 vTex;

uniform sampler2D uTexture; // 默认使用0
uniform int uIndex;

void main()
{
    if(uIndex==0){
        FragColor0 = texture(uTexture,vTex);
    }else if(uIndex==1){
        FragColor1 = texture(uTexture,vTex);
    }
//    else if(uIndex==2){
//        FragColor2 = texture(uTexture,vTex);
//    }else if(uIndex==3){
//        FragColor3 = texture(uTexture,vTex);
//    }else if(uIndex==4){
//        FragColor4 = texture(uTexture,vTex);
//    }
//    FragColor0 = texture(uTexture,vTex);
//    FragColor1 = texture(uTexture,vTex);
}