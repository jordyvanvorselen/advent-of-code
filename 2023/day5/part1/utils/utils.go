package utils

func FindMinInSlice(slice []int) int {
	if len(slice) == 0 {
		return 0
	}

	// Initialize min with the first element of the slice
	min := slice[0]

	// Iterate through the rest of the elements
	for _, num := range slice {
		if num < min {
			min = num
		}
	}

	return min
}
