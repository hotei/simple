// goroutine.go (c) David Rook 2012 - released under Simplified BSD 2-clause License

package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

const BigRunTime = 7

func cpu_hog(n int) float64 {
	var stress = false
	if stress {
		const BIG_NUMBER = 1E7
		maxj := rand.Int31n(100)
		var k float64
		fmt.Printf("cpu_hog called with n = %d, maxj = %v\n", n, maxj)
		for i := 1; i < BIG_NUMBER; i++ { // waste lots of time
			for j := int32(1); j < maxj; j++ {
				k = math.Sqrt(float64(i)*float64(j)) * math.Sqrt(float64(j))
			}
		}
		return k
	} else {
		const BIG_TIME = BigRunTime * 2
		// sleep random amount [0..BIG_TIME]
		time.Sleep(time.Duration(rand.Int31n(int32(BIG_TIME))) * time.Second)
		return rand.Float64() * 100.0
	}
	return 0.0 // NOT-REACHED but go compiler requires it (OBE?)
}

func main() {
	fmt.Printf("At start there are %d goroutines running\n",
		runtime.NumGoroutine())
	NUM_CPUS := runtime.NumCPU()
	loops_to_do := NUM_CPUS * 5
	estTime := (BigRunTime * loops_to_do) / NUM_CPUS
	fmt.Printf("est runtime %d to %d seconds...\n", estTime, estTime*2)
	throttle := make(chan int, NUM_CPUS)
	fmt.Printf("starting %d goroutines in total\n", loops_to_do)
	var wg sync.WaitGroup
	for i := 1; i <= loops_to_do; i++ {
		throttle <- 1
		wg.Add(1)
		go func(z int) {
			fmt.Printf("goroutine[%d] started\n", z)
			rv := cpu_hog(z)
			_ = <-throttle
			wg.Done()
			fmt.Printf("cpu_hog(%d) = %.2f\n", z, rv)
		}(i)
	}
	fmt.Printf("%d new goroutines were started, %d are still running\n",
		loops_to_do, runtime.NumGoroutine())
	wg.Wait()
	fmt.Printf("At finish there are %d goroutines running\n",
		runtime.NumGoroutine())
}
