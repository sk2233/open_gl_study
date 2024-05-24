#version 410 core
layout (location = 0) in vec3 iPos;
layout (location = 1) in vec2 iTex;
layout (location = 2) in vec3 iNor;

uniform mat4 uModel;
uniform mat4 uModelIT;
uniform mat4 uView;
uniform mat4 uProjection;

uniform mat4 uLight;  // 光场的 P * V

out vec4 vPosInLight;
out vec2 vTex;
out vec3 vNor;
out vec3 vPos;
out float vDep;

void main()
{
    gl_Position = uProjection * uView * uModel * vec4(iPos, 1.0);
    // 获取在广场下的位置
    vPosInLight = uLight * uModel * vec4(iPos,1.0);
    vTex=iTex;
    vNor = vec3(uModelIT*vec4(iNor,1.0));
    vPos = vec3(uModel*vec4(iPos,1.0));
    vDep = gl_Position.z/gl_Position.w;
}