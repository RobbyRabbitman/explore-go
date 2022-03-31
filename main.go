package main

import (
	"explore-go/internals/concurrency"
)

func main() {
	//types.PrintZeroValues()
	//concurrency.Greet(4)
	//concurrency.Deadlock1()
	//concurrency.Deadlock2()
	//fmt.Println(concurrency.Sum(3, 41, 2, 32, 12, 4, 123, 5, 5123, 4, 23))
	concurrency.ConcurrentParallelSequential()
}
