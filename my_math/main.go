package my_math

func Add(a, b int) int {
	return a + b
}

type Number interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64
}

func AddGenerics[T Number](a, b T) T {
	return a + b
}

func Mul[T Number](a, b T) T {
	return a * b
}
