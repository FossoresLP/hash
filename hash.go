// This package provides a function to hash arbitrary data using xxhash64.
// Any type that can be resolved to a primitive type (bool, int, uint, float, complex, string) can be hashed.
// Arrays, slices, structs, maps, interfaces, and pointers are supported.
// Types that cannot be resolved to a primitive type (chan, func, unsafe.Pointer) are not supported and will be silently ignored.
// The hash is calculated using reflection, so there is a performance cost and the result is not guaranteed to be stable across different Go versions.
// The hash is calculated in a way that avoids heap allocations as much as possible. Only reflection on maps causes heap allocations.
// A special case is implemented for map[string]any to avoid heap allocations.

package hash

import (
	"encoding/binary"
	"hash"
	"math"
	"reflect"
	"unsafe"

	"github.com/cespare/xxhash/v2"
	"github.com/zeebo/xxh3"
)

// Hash tries to hash the given data using xxhash64 using reflection.
// See the package documentation for more information on supported types and limitations.
func Hash(data any) []byte {
	h := xxhash.New()
	buf := make([]byte, 16)
	hashValue(reflect.ValueOf(data), h, buf)
	return h.Sum(buf[:0])
}

// Hash128 tries to hash the given data using xxhash128 (a variant of xxh3) using reflection.
// See the package documentation for more information on supported types and limitations.
func Hash128(data any) []byte {
	h := xxh3.New()
	buf := make([]byte, 16)
	hashValue(reflect.ValueOf(data), h, buf)
	buf = buf[:16]
	sum := h.Sum128()
	binary.BigEndian.PutUint64(buf[:8], sum.Hi)
	binary.BigEndian.PutUint64(buf[8:], sum.Lo)
	return buf
}

// HashWithHash tries to hash the given data using the provided hash.Hash using reflection.
// See the package documentation for more information on supported types and limitations.
// The hash.Hash must be reset before calling this function and it must not be used concurrently.
// The hash.Hash must be a hash function that produces a fixed-size output and writes must not return an error.
func HashWithHash(data any, h hash.Hash) []byte {
	buf := make([]byte, max(16, h.Size()))
	hashValue(reflect.ValueOf(data), h, buf)
	h.Sum(buf[:0])
	return buf
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
		switch m := v.Interface().(type) {
		case map[string]any:
			for k, v := range m {
				h.Write(unsafe.Slice(unsafe.StringData(k), len(k)))
				hashValue(reflect.ValueOf(v), h, buf)
			}
		default:
			for i := v.MapRange(); i.Next(); {
				hashValue(i.Key(), h, buf)
				hashValue(i.Value(), h, buf)
			}
		}
	case reflect.Interface, reflect.Pointer:
		if !v.IsNil() {
			hashValue(v.Elem(), h, buf)
		}
	}
}
