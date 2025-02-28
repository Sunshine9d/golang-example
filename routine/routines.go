package routine

func calculateSquare(x int) int {
	return x * x
}

func calculate() {
	for i := 0; i < 10000; i++ {
		calculateSquare(i)

	}
}
