#version 410 core

out vec4 FragColor;

in vec2 vTex;

uniform sampler2D uTexture0;
uniform sampler2D uTexture1;
uniform float uWidth = 2;

// 先通过法线判断 是否为面片拼接处  再通过深度 判断 是否为边缘
bool isEdge(vec4 p1,vec4 p2){
    return float(abs(p1.r-p2.r))>0.1;
}

void main()
{
    FragColor = texture(uTexture0,vTex);
    float offsetX = vTex.x/gl_FragCoord.x;
    float offsetY = vTex.y/gl_FragCoord.y;
    // 多采样
    vec4 v1 = texture(uTexture1,vTex+vec2(-offsetX,-offsetY)*uWidth);
    vec4 v2 = texture(uTexture1,vTex+vec2(offsetX,-offsetY)*uWidth);
    vec4 v3 = texture(uTexture1,vTex+vec2(-offsetX,offsetY)*uWidth);
    vec4 v4 = texture(uTexture1,vTex+vec2(offsetX,offsetY)*uWidth);
    // 根据 法线与 深度 是否突变 进行描边
    if(isEdge(v1,v4) || isEdge(v2,v3)){
        FragColor = vec4(1.0,0.0,0.0,1.0);
    }
}
