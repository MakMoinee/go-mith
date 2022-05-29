package utils

type Number interface {
	int | float32 | float64 | int32 | int16 | int8
}

func SumNumber[T Number](t []T) T {
	var total T

	for _, data := range t {
		total += data
	}

	return total
}
