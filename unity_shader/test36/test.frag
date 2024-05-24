#version 410 core

layout (location = 0) out vec4 FragColor;
layout (location = 1) out vec4 BrightColor;

in vec2 vTex;

uniform sampler2D uTexture;

void main()
{
    FragColor= texture(uTexture,vTex);
    float brightness = dot(FragColor.rgb, vec3(0.2126, 0.7152, 0.0722));
    if (brightness>0.5){
        BrightColor = vec4(FragColor.rgb, 1.0);
    }
}