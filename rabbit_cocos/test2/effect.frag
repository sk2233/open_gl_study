#version 410 core

out vec4 FragColor;

in vec2 vTex;

uniform sampler2D uTexture; // 默认使用0

float lum(vec4 col){
    return 0.2125 * col.r+ 0.7154 * col.g+ 0.0721 * col.b;
}

void main()
{
    vec2 offset = 1.0/vec2(1280.0,720.0);
    vec2 offsets[9] = vec2[](-offset,vec2(0,-offset.y),vec2(offset.x,-offset.y),vec2(-offset.x,0),vec2(0.0),vec2(offset.x,0),vec2(-offset.x,offset.y),vec2(0,offset.y),offset);
    float gxs[9] = float[](-1,-2,-1,0,0,0,1,2,1);
    float gys[9] = float[](-1,0,1,-2,0,2,-1,0,1);
    float gx = 0.0;
    float gy = 0.0;
    for(int i=0;i<9;i++){
        vec4 col = texture(uTexture,vTex+offsets[i]);
        float temp = lum(col);
        gx+=gxs[i]*temp;
        gy+=gys[i]*temp;
    }
    float temp = abs(gx)+abs(gy);
    FragColor = vec4(temp,temp,temp,1.0);
}