import "fmt"

func main() {
	done := make(chan interface{})

	res := intGenerator(done, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	outta := sq(res)

	for {
		select {
		case i, ok := <-outta:
			if !ok {
				outta = nil
			} else {
				fmt.Println(i)
			}
		}

		if outta == nil {
			break
		}
	}
}

func intGenerator(done <-chan interface{}, i ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, val := range i {
			out <- val
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	resChan := make(chan int)

	go func() {
		for val := range in {
			resChan <- val
		}
		close(resChan)
	}()
	return resChan
}
