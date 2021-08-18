package main

import (
	"fmt"
	"math/big"
	"sync"
	"time"
)

func runner1(wg *sync.WaitGroup) {
	defer wg.Done() // This decreases counter by 1
	fmt.Print("\nI am first runner")

}

func runner2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print("\nI am second runner")
}

func execute() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	// We are increasing the counter by 2
	// because we have 2 goroutines
	go runner1(wg)
	go runner2(wg)

	// This Blocks the execution
	// until its counter become 0
	wg.Wait()
}

func factorial(x *big.Int) *big.Int {
	n := big.NewInt(1)
	if x.Cmp(big.NewInt(0)) == 0 {
		return n
	}
	return n.Mul(x, factorial(n.Sub(x, n)))
}

func factorialParallel(x *big.Int, wg *sync.WaitGroup) {

	factorial(x)
	defer wg.Done()

}

func say3(s string, wg *sync.WaitGroup) {

	for i := 0; i < 1000; i++ {
		wg.Add(150)
		for j := 0; j < 150; j++ {
			go factorialParallel(big.NewInt(10), wg)
		}
		wg.Wait()
	}

}

func say(s string) {
	var wg1 sync.WaitGroup
	wg1.Add(1)
	say3(s, &wg1)
	wg1.Wait()
	say1("da bre")
}

func say3NotParallel(s string) {
	for i := 0; i < 1000*150; i++ {
		factorial(big.NewInt(10))
	}
}

func say1(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func main3() {
	//var wg sync.WaitGroup
	//wg.Add(1)
	start := time.Now()
	//say("world")
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println("TIME: ", elapsed/1)
	//wg.Wait()
	fmt.Println("TIME 0: ", elapsed/1)
	var wg sync.WaitGroup
	start1 := time.Now()

	say3("world", &wg)
	wg.Wait()
	end1 := time.Now()
	elapsed1 := end1.Sub(start1)

	start2 := time.Now()
	say3NotParallel("world")
	end2 := time.Now()
	elapsed2 := end2.Sub(start2)

	fmt.Println(int(150/100)*100 + 150%100)
	fmt.Println("TIME 1: ", elapsed1/1)
	fmt.Println("TIME 2: ", elapsed2/1)
	say1("hello")
}
func abc(i int, slice []string, wg *sync.WaitGroup) {
	defer wg.Done()
	val := slice[i]
	fmt.Printf("i: %v, val: %v\n", i, val)
}

func printABC(slice []string) {
	for _, val := range slice {
		fmt.Printf("val: %v\n", val)
	}
}

func main2() {
	slice := []string{"a", "b", "c", "d", "e", "q", "a", "b", "c", "d", "e", "q"}
	for i := 'A'; i <= 'Z'; i++ {
		slice = append(slice, string(i))
	}
	//MORA U PARALELI TA PETLJA A NE U ODNOSU NA OSTALO
	//	sliceLength := len(slice)
	//var wg sync.WaitGroup
	go printABC(slice)
	/*sliceLength := len(slice)
	/var wg sync.WaitGroup
	wg.Add(sliceLength)
	fmt.Println("Running for loop…")
	for i := 0; i < sliceLength; i++ {
		go abc(i, slice, &wg)

	}

	wg.Wait()*/
	fmt.Println("Finished for loop")
}
