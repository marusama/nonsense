package semaphore

import (
	"context"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestSemaphore_Acquire(t *testing.T) {

	println(runtime.GOMAXPROCS(0))

	s := New(1)

	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			s.Acquire(context.Background())
			time.Sleep(10 * time.Minute)
			wg.Done()
		}()
	}

	wg.Wait()

	println("Done")
}

func BenchmarkSemaphore_Acquire(b *testing.B) {
	sem := New(b.N)

	for i := 0; i < b.N; i++ {
		_ = sem.Acquire(nil)
	}

	if sem.Occupied() != sem.GetLimit() {
		b.Error("full filled semaphore is expected")
	}
}

func BenchmarkSemaphore_Acquire_Release(b *testing.B) {
	sem := New(b.N)

	for i := 0; i < b.N; i++ {
		_ = sem.Acquire(nil)
		sem.Release()
	}

	if sem.Occupied() != 0 {
		b.Error("empty semaphore is expected")
	}
}
