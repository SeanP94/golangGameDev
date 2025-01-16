#version 330

// Input vertex attributes
in vec3 vertexPosition;
in vec2 vertexTexCoord;
in vec3 vertexNormal;
in vec4 vertexColor;

// Input uniform values
uniform mat4 mvp;

// Output vertex attributes (to fragment shader)
out vec2 fragTexCoord;
out vec4 fragColor;

// NOTE: Add here your custom variables
varying vec2 pos;


attribute vec3 aPosition;
attribute vec2 aTextCoord;


void main()
{
    // Send vertex attributes to fragment shader
    // fragTexCoord = vertexTexCoord;
    // fragColor = vertexColor;

    // Calculate final vertex position
    // gl_Position = mvp*vec4(vertexPosition, 1.0);

    pos = aTextCoord;
    vec4 position = vec4(aPosition, 1.0);
    position.xy = position.xy * 2.1 - 1.;
    
    gl_Position = position;
}

// #version 330 core

// in vec3 vertexPosition;
// in vec3 vertexNormal;
// in vec4 vertexColor;

// out vec4 fragColor;
// out vec3 fragPosition;
// out vec3 fragNormal;

// uniform mat4 matModel;
// uniform mat4 matView;
// uniform mat4 matProjection;

// void main()
// {
//     fragPosition = vec3(matModel * vec4(vertexPosition, 1.0));
//     fragNormal = vertexNormal;  
//     fragColor = vertexColor;
    
//     gl_Position = matProjection * matView * vec4(fragPosition, 1.0);
// }