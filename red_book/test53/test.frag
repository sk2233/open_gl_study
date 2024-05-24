#version 410 core

out vec4 FragColor;

in vec2 gTex;

uniform sampler2D uTexture;

void main()
{
    FragColor = texture(uTexture,gTex);
//    vec4 temp = imageLoad(uTexture,gTex);
//    imageStore()
//    imageSize()
//    imageAtomicXxx()
}