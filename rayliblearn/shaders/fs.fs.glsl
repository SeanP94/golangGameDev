#version 330

// Input vertex attributes (from vertex shader)
in vec3 fragPosition;
in vec2 fragTexCoord;
//in vec4 fragColor;
in vec3 fragNormal;

// Input uniform values
uniform sampler2D texture0;
uniform vec4 colDiffuse;

// Output fragment color
// out vec4 finalColor;

// NOTE: Add here your custom variables
uniform vec4 ambient;
uniform vec3 viewPos;


void main() {
    vec4 c1 = vec4(0.5, .1, .9, .1);
    vec4 c2 = vec4(0.1, .8, .7, .1);
    vec4 c  = mix(c1,c2, fragPosition.x);

    gl_FragColor = c;
}