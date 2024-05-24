#version 410 core

out vec4 FragColor;

in vec2 vTex;
in vec3 vPos;
in vec3 vNor;

uniform sampler2D uTexture; // 默认使用0
uniform vec3 uEye;
uniform vec4 uCol = vec4(0.0,1.0,0.0,1.0);

void main()
{
    float temp = 1.0-abs(dot(normalize(vNor),normalize(uEye-vPos)));
    float rate = pow(temp,2.2);
    vec4 col = texture(uTexture,vTex);
    col = vec4(1.0,1.0,1.0,0.0);
    FragColor = mix(col,uCol,rate);
}