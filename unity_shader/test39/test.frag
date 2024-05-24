#version 410 core

out vec4 FragColor;

in vec3 vNor;
in vec4 vPos;

void main()
{
//    FragColor = vec4(vNor/2+0.5,1.0);
    FragColor = vec4(vec3(vPos.z/vPos.w),1.0);
}