#version 410 core

out vec4 FragColor;

in vec2 vTex;
in vec3 vPos;

uniform sampler2D uImage;
uniform vec3 uLightPos;
uniform vec3 uEyePos;

void main()
{
    vec3 color = texture(uImage,vTex).rgb;
    color+=normalize(uEyePos);
    color = color/(color+vec3(1));
    FragColor = vec4(color,1);
}