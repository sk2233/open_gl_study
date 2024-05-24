#version 410 core

out vec4 FragColor;

in vec2 tempLoc; // 贴图 位置

uniform sampler2D aTexture; // 默认使用0

void main()
{
  vec4 temp =texture(aTexture,tempLoc);
//  if(temp.a<0.1){
//    discard;
//  }
  FragColor = temp;
}