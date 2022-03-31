package concurrency

import (
	"fmt"
	"time"
)

func Greet(number int) {
	sleepMs := 100
	print := func(number, sleep int, a ...any) {
		for i := 0; i < number; i++ {
			time.Sleep(time.Millisecond * time.Duration(sleep))
			fmt.Println(a...)
		}
	}

	// number, sleepMs, "Mundo" are evaluated in this goroutine, but print is executed in the spawning goroutine
	go print(number, sleepMs, "Mundo")
	print(number, sleepMs, "Hola")

	// This would be blocking: so n times "Hola", then n times "Mundo"
	// print(number, sleepMs, "Hola")
	// go print(number, sleepMs, "Mundo")

	// go print(number, sleepMs, "Mundo")
	// go print(number, sleepMs, "Hola")
	// If Greet should exit, when goroutines are done
	// yes, its hacky :D
	// time.Sleep(time.Millisecond * time.Duration(sleepMs*number*2))
}
