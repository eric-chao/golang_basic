package practice

func walk(n int) int {

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	return walk(n-1) + walk(n-2)
}
