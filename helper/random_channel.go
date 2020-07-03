package helper

func Random(n int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 0; i < n; i++ {
			select {
			case c <- 0:
			case c <- 1:
			case c <- 2:
			case c <- 3:
			case c <- 4:
			case c <- 5:
			case c <- 6:
			case c <- 7:
			case c <- 8:
			case c <- 9:
			}
		}
	}()
	return c
}
