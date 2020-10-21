package function

type F func(int int) int

func higher(v int, f F) int {
	return f(v)
}