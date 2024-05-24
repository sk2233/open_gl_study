#version 410 core
out vec4 FragColor;

in vec2 TexCoords;

uniform sampler2D screenTexture;

void main()
{
    vec3 col = texture(screenTexture, TexCoords).rgb;
//    FragColor = vec4(vec3(1)-col, 1.0);// 反向
    float avg = (col.r+col.g+col.b)/3;// 灰度
    FragColor = vec4(avg,avg,avg, 1.0);
}