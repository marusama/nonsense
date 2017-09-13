package chan_chan

import (
	"sync"
	"fmt"
	"time"
	"github.com/davecgh/go-spew/spew"
	"github.com/soider/d"
)

type sss []sss
type mmm map[int]mmm
type ccc chan ccc
type fff func () fff




type fff2 func (int) fff2

type fff3 func (fff3) fff3

type c struct {
  some *c
}

/*
channel in channel
 */
func Run() {

	wg := sync.WaitGroup{}
	ch := make(chan chan int, 10)

	wg.Add(1)
	go func() {
		defer wg.Done()

		wg_int := sync.WaitGroup{}
		for c := range ch {
			wg_int.Add(1)
			go func(ci chan int) {
				defer wg_int.Done()

				for val := range ci {
					fmt.Println(val)
					time.Sleep(time.Second)
				}
			}(c)
		}
		wg_int.Wait()
	}()

	for i := 0; i < 3; i++ {
		c := make(chan int)
		go func(num int) {
			for val := num * 10; val < num * 10 + 10; val++ {
				c <- val
			}
			close(c)
		}(i)
		ch <- c
	}

	close(ch)
	wg.Wait()

	fmt.Println("Done")

}

func Run2() {
	s := make(sss, 5)
	s[0] = sss{}
	s[0] = append(s[0], sss{})
	spew.Dump(s)
	d.D(s)

	m := make(mmm)
	spew.Dump(m)
	d.D(m)

	m2 := mmm{}
	spew.Dump(m2)
	d.D(m2)

	ch := make(ccc, 10)
	spew.Dump(ch)
	d.D(ch)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for c := range ch {
			spew.Dump(c)
			d.D(c)
		}
		wg.Done()
	}()

	ch <- make(ccc)
	close(ch)
	wg.Wait()

	var f fff
	f = func () fff {
		fmt.Println("here")
		return f
	}
	f()()()()

	var f2 fff2
	f2 = func(a int) fff2 {
		fmt.Println("f2", a)
		return f2
	}
	f2(1)(2)(3)

	//var f3 fff3
	//f3 = func(ff3 fff3) fff3 {
	//	fmt.Println("f3")
	//	return ff3(f3)
	//}
	//f3(f3)
}