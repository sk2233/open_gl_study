#version 410 core
out vec4 FragColor;

in vec3 vPos;

uniform vec3 uLightPos;
uniform samplerCube uShadow;
uniform float uFar;

void main()
{
    FragColor = vec4(1.0);
    float depth = texture(uShadow,vPos-uLightPos).r;
    float len = distance(vPos,uLightPos);
    if (len > depth * uFar+0.005){
        FragColor*=0.5;
    }
}