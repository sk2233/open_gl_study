#version 410 core
out vec4 FragColor;

in vec4 vPosInLight;

uniform sampler2D uShadow;

void main()
{
    FragColor = vec4(1.0);
    // 齐次化
    vec3 pos = vPosInLight.xyz/vPosInLight.w;
    pos = pos*0.5+0.5;
    // 当前深度 大于光照 贴图深度  有阴影  加点阈值 深度贴图精度不够
    if( pos.z > texture(uShadow,pos.xy).r+0.005){
        FragColor.r *= 0.5;
    }
}