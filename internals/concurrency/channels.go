package concurrency

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"time"
)

// channels are a mechanism for goruotines to communicate with each other.
// channels are created with the make function.
// By default, sends and receives block until the other side is ready!!!

func Deadlock1() {
	ch := make(chan bool)
	x := <-ch  // waiting til channel sends a value
	ch <- true // sending never gets executed
	fmt.Println(x)
}

func Deadlock2() {
	ch := make(chan bool)
	ch <- true // send value and wait til receiver reads it
	x := <-ch  // receiving never gets executed
	fmt.Println(x)
}

func Sum(s ...int) int {

	middle := len(s) / 2
	s1 := s[:middle]
	s2 := s[middle:]

	if middle < 5 {
		sum := func(s []int, ch chan int) {
			result := 0
			for _, s := range s {
				result += s
			}
			ch <- result
		}
		ch := make(chan int)

		go sum(s1, ch)
		go sum(s2, ch)
		x, y := <-ch, <-ch
		return x + y
	} else {
		return Sum(s1...) + Sum(s2...)
	}

}

func ConcurrentParallelSequential() {
	url := "https://jsonplaceholder.typicode.com/todos"

	threads := func() int {
		maxProcs := runtime.GOMAXPROCS(0)
		numCPU := runtime.NumCPU()
		if maxProcs < numCPU {
			return maxProcs
		}
		return numCPU
	}()
	number := threads * 2
	fmt.Printf("Max Parallel: %v, Requests: %v\n", threads, number)
	start := time.Now()
	result := make([]time.Duration, number)
	func() {
		var wg sync.WaitGroup
		wg.Add(number)
		buffer := make(chan time.Duration, number)
		for i := 1; i < number+1; i++ {
			go func(i *int) {
				x := time.Now()
				http.Get(fmt.Sprintf("%v/%v", url, *i))
				buffer <- time.Since(x)
				wg.Done()
			}(&i)
		}
		go func() {
			wg.Wait()
			close(buffer)
		}()
		index := 0
		for time := range buffer {
			result[index] = time
			index++
		}

	}()
	fmt.Printf("Concurrent (%v)\nEach Request Function %v\n\n", time.Since(start), result)

	start = time.Now()
	func() {
		for i := 1; i < number+1; i++ {
			x := time.Now()
			http.Get(fmt.Sprintf("%v/%v", url, i))
			result[i-1] = time.Since(x)
		}
	}()
	fmt.Printf("Sequentiel (%v)\nEach Request Function %v\n\n", time.Since(start), result)

}
