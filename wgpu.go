package main

type CommandEncoderDescription string
type DeviceDescriptor int

// Backend
type Backend int

const (
	Vulkan = Backend(iota) // Vulkan
	Metal                  // Metal (Apple)
	D3D12                  // Direct3D-12 (Windows)
	OpenGL                 // OpenGL ES-3 (Linux, Android)
	WebGPU                 // WebGPU in the browser
	WebGL2                 // WebGL2 in the browser
)

// Power Preference
type PowerPreference int

const (
	LowPower = iota
	HighPerformance
)

// Blend Factor
type BlendFactor int

const (
	Zero              = BlendFactor(iota) // 0.0
	One                                   // 1.0
	Src                                   // S.component
	OneMinusSrc                           // 1.0 - S.component
	SrcAlpha                              // S.alpha
	OneMinusSrcAlpha                      // 1.0 - S.alpha
	Dst                                   // D.component
	OneMinusDst                           // 1.0 - D.component
	DstAlpha                              // D.alpha
	OneMinusDstAlpha                      // 1.0 - D.alpha
	SrcAlphaSaturated                     // min(S.alpha, 1.0 - D.alpha)
	Constant                              // Constant
	OneMinusConstant                      // 1.0 - Constant
)

// Blend Operation
type BlendOperation int

const (
	Add             = BlendOperation(iota) // Src + Dst
	Subtract                               // Src - Dst
	ReverseSubtract                        // Dst - Src
	Min                                    // min(Src, Dst)
	Max                                    // max(Src, Dst)
)
