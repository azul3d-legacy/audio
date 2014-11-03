# Azul3D - audio #
This package provides a generic audio package for Go. It is very similar to Go's image package, except for audio.

# Version 2 #
* Documentation
 * [azul3d.org/audio.v1](http://azul3d.org/audio.v2)
 * `import "azul3d.org/audio.v2"`
* Changes
 * Single-sample types were replaced by their direct Go type (see [#6](https://github.com/azul3d/audio/issues/6):
  * `PCM8  -> uint8`
  * `PCM16 -> int16`
  * `PCM32 -> int32`
  * `ALaw  -> uint8`
  * `MuLaw -> uint8`
 * Slices of samples types were renamed (see [#6](https://github.com/azul3d/audio/issues/6):
  * `PCM8         -> type Uint8 []uint8`
  * `PCM16        -> type Int16 []int16`
  * `PCM32        -> type Int32 []int32`
  * `ALaw         -> type ALaw []uint8`
  * `MuLawSamples -> type MuLaw []uint8`
 * Conversion functions were renamed (see [#6](https://github.com/azul3d/audio/issues/6):
  * `ALawToPCM16 -> ALawToInt16`
  * `F64ToPCM16   -> Float64ToInt16`
  * `F64ToPCM32   -> Float64ToInt32`
  * `F64ToPCM8    -> Float64ToUint8`
  * `PCM16ToALaw  -> Int16ToALaw`
  * `PCM16ToF64   -> Int16ToFloat64`
  * `PCM16ToMuLaw -> Int16ToMuLaw`
  * `PCM32ToF64   -> Int32ToFloat64`
  * `MuLawToPCM16 -> MuLawToInt16`
  * `PCM8ToF64    -> Uint8ToFloat64`

# Version 1 #
* Documentation
 * [azul3d.org/audio.v1](http://azul3d.org/audio.v1)
 * `import "azul3d.org/audio.v1"`

