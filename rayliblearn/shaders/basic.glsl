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

uniform float mytime;
uniform vec2 resolution;
uniform vec2 rezOp;

vec3 colorA = vec3(0.149,0.141,0.912);
vec3 colorB = vec3(1.000,0.833,0.224);

void main() {

    // vec3 color = vec3(0.0);
    float pct = abs(sin(mytime));
    vec4 c1 = vec4(resolution, .9, pct);
    vec4 c2 = vec4(rezOp, .9, pct);
    vec4 c = mix(c1, c2, 1);
    
    // gl_FragColor = vec4(size,1, 0,0);
    gl_FragColor = c;
}