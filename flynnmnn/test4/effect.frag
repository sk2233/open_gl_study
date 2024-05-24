#version 410 core
out vec4 FragColor;

in vec2 vTex;
in float vDep;

uniform sampler2D uImage;
uniform sampler2D uDep;
uniform bool uOpen;

void main()
{
    FragColor = texture(uImage,vTex);
    if(uOpen){
        // 深度 越接近 越是交接处
        vec2 tex = gl_FragCoord.xy/vec2(1280.0,720.0);
        float dep = texture(uDep,tex).r;
        float diff = vDep-dep;
        FragColor.r = 1.0 - smoothstep(0.0,0.5,diff);
    }
}