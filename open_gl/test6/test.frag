#version 410 core

out vec4 FragColor;

in vec2 tempLoc; // 贴图 位置
in vec3 tempNormal; // 法线方向
in vec3 tempPos; // 世界坐标下的位置

uniform vec3 aAmbient;// 环境光照
uniform vec3 aLightPos;// 光的位置
uniform vec3 aLightDir;// 光的方向
uniform vec3 aEyePos;
uniform sampler2D aTexture; // 默认使用0
uniform sampler2D specTexture;

struct Material{
    vec3 ambient;
    vec3 diffuse;
    vec3 specular;
    float shininess;
};

uniform Material aMaterial;

void main()
{
    // 计算漫反射
    // 计算法线与 入射光线的夹角
    vec3 lightDir = normalize(aLightPos-tempPos);
//    lightDir = normalize(aLightDir);// 测试  平行光
//    if (dot(-lightDir,normalize(aLightDir))<0.0001){// 位于范围外 只给个环境光
//        FragColor =texture(aTexture,tempLoc)*vec4((aAmbient*aMaterial.ambient),1);
//    }else{// 聚光灯效果
//
//    }
    float tempDot = dot(normalize(tempNormal),lightDir);
    float diff = max(tempDot,0.0)/length(aLightPos-tempPos);// 防止负数出现
    // 计算高光
    vec3 reflectDir = reflect(-lightDir, tempNormal);// 计算反射光
    float spec = pow(max(dot(reflectDir,normalize(tempPos-aEyePos)),0.0),aMaterial.shininess);// 计算反射光与视线的夹角
    FragColor =texture(aTexture,tempLoc)*(vec4((aAmbient*aMaterial.ambient+diff*aMaterial.diffuse),1)+
    vec4(spec*aMaterial.specular,1)*texture(specTexture,tempLoc));
}