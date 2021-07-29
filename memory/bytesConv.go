package memory

import (
	"encoding/binary"
	"math"
)

func Float32ToByte(f float32) []byte {
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, math.Float32bits(f))
	return data
}

func ByteToFloat32(data []byte) float32 {
	return math.Float32frombits(binary.LittleEndian.Uint32(data))
}

func Floats32ToByte(f []float32) []byte {
	data := make([]byte, 4*len(f))
	for i, f := range f {
		binary.LittleEndian.PutUint32(data[i*4:], math.Float32bits(f))
	}
	return data
}

func ByteToFloats32(data []byte) []float32 {
	f := make([]float32, len(data)/4)
	for g := 0; g < len(data); g += 4 {
		f[g/4] = math.Float32frombits(binary.LittleEndian.Uint32(data[g:]))
	}

	return f
}

func IntToByte(i int) []byte {
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, uint32(i))
	return data
}

func ByteToInt(data []byte) int {
	return int(binary.LittleEndian.Uint32(data))
}

func ByteToIntPtr(data []byte) uintptr {
	return uintptr(binary.LittleEndian.Uint32(data))
}

func IntsToByte(i []int) []byte {
	data := make([]byte, 4*len(i))
	for index, integer := range i {
		binary.LittleEndian.PutUint32(data[index*4:], uint32(integer))
	}

	return data
}

func ByteToInts(data []byte) []int {
	i := make([]int, len(data)/4)
	for g := 0; g < len(data); g += 4 {
		i[g/4] = int(binary.LittleEndian.Uint32(data[g:]))
	}

	return i
}
