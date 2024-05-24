#version 410 core

out vec4 FragColor;

in vec2 vTex;

uniform sampler2D uImage;// 直接使用 0即可

void main()
{
    // 颜色反转
//    FragColor= vec4(1-texture(uImage,vTex).xyz,1);
    float offsetX = vTex.x/gl_FragCoord.x; // 获取 x轴偏移
    float offsetY = vTex.y/gl_FragCoord.y; // 获取 y轴偏移
    vec4 temp = texture(uImage,vTex)*0.227027;
    float weight[2] = float[](0.1216216,0.016216);
    for(int i=1;i<3;i++){
        temp +=texture(uImage,vTex+vec2(float(i)*offsetX,0))*0.2442;
        temp +=texture(uImage,vTex+vec2(float(-i)*offsetX,0))*weight[i-1];
        temp +=texture(uImage,vTex+vec2(0,float(i)*offsetY))*weight[i-1];
        temp +=texture(uImage,vTex+vec2(0,float(-i)*offsetY))*weight[i-1];
    }
    FragColor = temp;
}