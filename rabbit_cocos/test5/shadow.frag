#version 410 core
layout (location = 0) out vec4 FragColor;
layout (location = 1) out vec4 DepColor;

in vec4 vPosInLight;
in vec2 vTex;
in vec3 vNor;
in vec3 vPos;
in float vDep;

uniform sampler2D uTexture;
uniform sampler2D uShadow;
uniform vec3 uEye;
uniform vec4 uCol = vec4(0.0,1.0,1.0,1.0);
uniform bool uOpen;

void main()
{
    FragColor = texture(uTexture,vTex);
    // 齐次化
    vec3 pos = vPosInLight.xyz/vPosInLight.w;
    pos = pos*0.5+0.5;// 从 -1~1  ->   0~1
    // 当前深度 大于光照 贴图深度  有阴影  加点阈值 深度贴图精度不够
    if( pos.z > texture(uShadow,pos.xy).r+0.001){
        FragColor.rgba *= 0.5;
    }
    if (uOpen){
        float temp = 1.0-abs(dot(normalize(vNor),normalize(uEye-vPos)));
        float rate = pow(temp,2.2);
        FragColor = mix(FragColor,uCol,rate);
        DepColor = vec4(vDep);
    }
}