package main

import (
	"fmt"
	"time"
)

// -------- 1 ----
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

//-----------2----------

// func worker(id int) {
// 	fmt.Printf("Worker %d starting\n", id)
// 	time.Sleep(1 * time.Second)
// 	fmt.Printf("Worker %d done\n", id)
// }

func main() {
	// --- 1 ---

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	time.Sleep(5 * time.Second)
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	// close(jobs)
	for j := 100; j <= 100+numJobs; j++ {
		jobs <- j
	}
	for a := 1; a <= 2*numJobs; a++ {
		<-results
	}

	// ---- 2 ----
	// var wg sync.WaitGroup

	// for i := 1; i <= 5; i++ {
	// 	if i == 1 {
	// 		wg.Add(1)

	// 		go func() {
	// 			worker(i)
	// 			defer wg.Done()
	// 		}()
	// 	}
	// }
	// wg.Wait()
}
