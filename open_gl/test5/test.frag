#version 410 core
out vec4 FragColor;

in vec2 TexCoord;

uniform vec4 ourColor;
uniform sampler2D ourTexture; // 默认就是  TEXTURE0 也可以手动赋值为0
uniform sampler2D testTexture; // 是 1i类型  根据 id 去找对应的 TEXTURE

void main()
{
    FragColor =texture(ourTexture,TexCoord)*texture(testTexture,TexCoord)*ourColor;
}