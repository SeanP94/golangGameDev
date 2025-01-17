    #ifdef GL_ES
        precision highp float;
    #endif
    // https://babylonjs.medium.com/retro-crt-shader-a-post-processing-effect-study-1cb3f783afbc
    #define PI 3.1415926538

    // Samplers
    varying vec2 vUV;
    uniform sampler2D textureSampler;

    // Parameters
    vec2 curvature = vec2(3.0, 3.0);
    vec2 scanLineOpacity = vec2(1,1);
    float vignetteOpacity = 1;
    
    uniform vec2 u_resolution;

    vec2 curveRemapUV(vec2 uv)
    {
        // as we near the edge of our screen apply greater distortion using a sinusoid.

        uv = uv * 2.0 - 1.0;
		vec2 offset = abs(uv.yx) / vec2(curvature.x, curvature.y);
		uv = uv + uv * offset * offset;
		uv = uv * 0.5 + 0.5;
        return uv;
    }

    vec4 scanLineIntensity(float uv, float resolution, float opacity)
    {
        float intensity = sin(uv * resolution * PI * 2.0);
        intensity = ((0.5 * intensity) + 0.5) * 0.9 + 0.1;
        return vec4(vec3(pow(intensity, opacity)), 1.0);
    }

    vec4 vignetteIntensity(vec2 uv, vec2 resolution, float opacity)
    {
        float intensity = uv.x * uv.y * (1.0 - uv.x) * (1.0 - uv.y);
        return vec4(vec3(clamp(pow((resolution.x / 4.0) * intensity, opacity), 0.0, 1.0)), 1.0);
    }

    void main(void) 
    {
        vec2 remappedUV vec2(1,1)// = curveRemapUV(vec2(vUV.x, vUV.y));
        vec4 baseColor = texture2D(textureSampler, remappedUV);

        baseColor *= vignetteIntensity(remappedUV, screenResolution, vignetteOpacity);

        baseColor *= scanLineIntensity(remappedUV.x, screenResolution.y, scanLineOpacity.x);
        baseColor *= scanLineIntensity(remappedUV.y, screenResolution.x, scanLineOpacity.y);

        if (remappedUV.x < 0.0 || remappedUV.y < 0.0 || remappedUV.x > 1.0 || remappedUV.y > 1.0){
            gl_FragColor = vec4(0.0, 0.0, 0.0, 1.0);
        } else {
            gl_FragColor = baseColor;
        }
    }

