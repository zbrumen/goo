package generics

type Number8 interface {
	~uint8 | ~int8
}

type Number16 interface {
	~uint16 | ~int16
}

type Number32 interface {
	~uint32 | ~int32
}

type Number64 interface {
	~uint | ~int | ~uint64 | ~int64 | ~uintptr
}

type Integer interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~int | ~int8 | ~int16 | ~int32 | ~int64
}

type UnsignedNumber interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type SignedNumber interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Decimal interface {
	~float32 | ~float64
}

type Number interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

type Complex interface {
	~complex64 | ~complex128
}

type Ordered interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64 | ~string
}

type Builtin interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64 | ~string | ~complex64 | ~complex128
}
