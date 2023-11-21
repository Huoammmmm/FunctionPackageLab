package mathops

func Add (a,b int) int {
	return a + b
}

func Subtract (a, b int) int  {
	return a - b
}

func Multiply (a, b int) int {
	return a * b
}

func Divide (a, b int) (int, int) {
	return a/b, a % b
}

func Power(base, exponent int) int {
    result := 1
    for i := 0; i < exponent; i++ {
        result *= base
    }
    return result
}
