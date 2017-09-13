package chan_func

import (
	"sync"
	"fmt"
)

func Run() {
	wg := sync.WaitGroup{}
	ch := make(chan func() error, 10)

	wg.Add(1)
	go func() {
		defer wg.Done()

		for f := range ch {
			err := f()
			if err != nil {
				panic(err)
			}
		}
	}()

	for i := 0; i < 3; i++ {
		num := i
		ch <- func() error {
			for val := num * 10; val < num * 10 + 3; val++ {
				fmt.Println(val)
			}
			return nil
		}
	}

	close(ch)
	wg.Wait()

	fmt.Println("Done")
}
