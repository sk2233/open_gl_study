#version 410 core

out vec4 FragColor;

in vec2 tempLoc; // 贴图 位置
in vec3 tempNormal; // 法线方向
in vec3 tempPos; // 世界坐标下的位置

uniform vec3 aAmbient;// 环境光照
uniform vec3 aLightPos;// 光的位置
uniform vec3 aEyePos;
uniform sampler2D aTexture; // 默认使用0

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
    vec3 lightDir = normalize(aLightPos-tempPos);
    float tempDot = dot(normalize(tempNormal),lightDir);
    float diff = max(tempDot,0.0);// 防止负数出现
    // 计算高光
    vec3 reflectDir = normalize(lightDir);// 计算反射光
    vec3 eyeDir = normalize(tempPos-aEyePos);
//    float spec = pow(max(dot(reflectDir,normalize(tempPos-aEyePos)),0.0),aMaterial.shininess);// 计算反射光与视线的夹角
    float spec = pow(dot(normalize(reflectDir+eyeDir),normalize(tempNormal)),4.0);
    FragColor =texture(aTexture,tempLoc)*(vec4((aAmbient*aMaterial.ambient+diff*aMaterial.diffuse+
                spec*aMaterial.specular),1));
//    FragColor = vec4(gl_FragCoord.xyz, 1.0);
}