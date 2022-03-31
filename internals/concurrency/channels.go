package concurrency

import "fmt"

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
