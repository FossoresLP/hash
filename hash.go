// This package provides a function to hash arbitrary data using xxhash64.
// Any type that can be resolved to a primitive type (bool, int, uint, float, complex, string) can be hashed.
// Arrays, slices, structs, maps, interfaces, and pointers are supported.
// Types that cannot be resolved to a primitive type (chan, func, unsafe.Pointer) are not supported and will be silently ignored.
// The hash is calculated using reflection, so there is a performance cost and the result is not guaranteed to be stable across different Go versions.

package hash

import (
	"encoding/binary"
	"hash"
	"math"
	"reflect"
	"unsafe"

	"github.com/cespare/xxhash/v2"
)

// Hash tries to hash the given data using xxhash64 using reflection.
// It supports the following types:
// - bool
// - int, int8, int16, int32, int64
// - uint, uint8, uint16, uint32, uint64, uintptr (by its value, not what it points to)
// - float32, float64
// - complex64, complex128
// - string
// - array, slice
// - struct
// - map
// - interface, pointer
// Data of other types such as chan, func, and unsafe.Pointer are not supported and will be silently ignored.
// Heap allocations are avoided as much as possible. Only maps cause heap allocations.
func Hash(data any) []byte {
	h := xxhash.New()
	buf := make([]byte, 16)
	hashValue(reflect.ValueOf(data), h, buf)
	return h.Sum(buf[:0])
}

func hashValue(v reflect.Value, h hash.Hash, buf []byte) {
	if !v.IsValid() {
		return
	}

	switch v.Kind() {
	case reflect.Bool:
		buf = buf[:1]
		if v.Bool() {
			buf[0] = 1
		} else {
			buf[0] = 0
		}
		h.Write(buf)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		buf = buf[:8]
		binary.BigEndian.PutUint64(buf, uint64(v.Int()))
		h.Write(buf)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		buf = buf[:8]
		binary.BigEndian.PutUint64(buf, v.Uint())
		h.Write(buf)
	case reflect.Float32, reflect.Float64:
		buf = buf[:8]
		binary.BigEndian.PutUint64(buf, math.Float64bits(v.Float()))
		h.Write(buf)
	case reflect.Complex64, reflect.Complex128:
		buf = buf[:16]
		binary.BigEndian.PutUint64(buf[:8], math.Float64bits(real(v.Complex())))
		binary.BigEndian.PutUint64(buf[8:], math.Float64bits(imag(v.Complex())))
		h.Write(buf)
	case reflect.String:
		str := v.String()
		h.Write(unsafe.Slice(unsafe.StringData(str), len(str)))
	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			hashValue(v.Index(i), h, buf)
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			hashValue(v.Field(i), h, buf)
		}
	case reflect.Map:
		for i := v.MapRange(); i.Next(); {
			hashValue(i.Key(), h, buf)
			hashValue(i.Value(), h, buf)
		}
	case reflect.Interface, reflect.Pointer:
		if !v.IsNil() {
			hashValue(v.Elem(), h, buf)
		}
	}
}
