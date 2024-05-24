#version 410 core

out vec4 FragColor;

in vec2 vTex;

uniform sampler2D uImage;
uniform float uOffset;

void main()
{
    vec2 tex = vec2(vTex.x/9+uOffset,vTex.y);
    FragColor = texture(uImage,tex);
}