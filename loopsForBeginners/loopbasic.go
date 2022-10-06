package loopsForBeginners

func BasicLoop () []int{
	a := make([]int, 0, 10)

	// set values for a slice
	for i := 0; i < 10 ; i++{
		a[i] = i
	}

	return a

}