#version 330
#ifdef GL_ES
precision mediump float;
#endif

uniform float u_time;

uniform vec2 u_resolution;

vec3 colorA = vec3(0.149,0.141,0.912);
vec3 colorB = vec3(1.000,0.833,0.224);

void main() {
    vec3 color = vec3(0.0);
    
    float pct = abs(sin(u_time));
    // float pct = abs(sin(1));
    
    // Mix uses pct (a value from 0-1) to
    // mix the two colors
    // color = mix(colorA, colorB, pct);
    vec4 c1 = vec4(.5, .1, .9, pct);
    vec4 c2 = vec4(.5, .1, .9, pct);
    vec4 c = mix(c1, c2, u_resolution.x);

    // gl_FragColor = vec4(color,1.0);
}