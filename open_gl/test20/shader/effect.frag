#version 410 core

out vec4 FragColor;

in vec2 vTex;

uniform sampler2D uImage0;
uniform sampler2D uImage1;
uniform float offset = 1.0/(54.0*5);
uniform float weight[3] = float[] (0.227027,  0.1216216,  0.016216);

void main()
{
   vec3 res = texture(uImage1,vTex).rgb*weight[0];
   for(int i=1;i<3;i++){
      res+=texture(uImage1,vTex+vec2(offset*i,0)).rgb*weight[i];
      res+=texture(uImage1,vTex-vec2(offset*i,0)).rgb*weight[i];
      res+=texture(uImage1,vTex+vec2(0,offset*i)).rgb*weight[i];
      res+=texture(uImage1,vTex-vec2(0,offset*i)).rgb*weight[i];
   }
   res+=texture(uImage0,vTex).rgb;
   res/=res+1;
   res=pow(res,vec3(1/2.2));
   FragColor= vec4(res,1);
}