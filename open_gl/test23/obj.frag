#version 410 core
out vec4 FragColor;

void main()
{
    FragColor = vec4(gl_FragCoord.z/100);
}