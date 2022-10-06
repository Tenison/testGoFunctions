package loopsForBeginners

func BasicLoop () []int{
	a := make([]int, 0, 10)

	// set values for a slice
	for i := 0; i < 10 ; i++{
		a = append(a, i)
	}

	return a
}

// create collection with values greater than five
func WhileLoops () []int{
	a := []int{1,2,3,4,5,6,7,9,0}
	b := make([]int,0,10)

	for _, value := range(a){
		if value > 5 {
			b = append(b, value)
		}
	}

	return b
}