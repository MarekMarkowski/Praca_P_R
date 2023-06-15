package main

import (
	"fmt"
	"sync"
)

func sumArrayPart(arr []int, start, end int, wg *sync.WaitGroup, result chan<- int) {
	sum := 0
	for i := start; i < end; i++ {
		sum += arr[i]
	}
	result <- sum
	wg.Done()
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numWorkers := 4 // Liczba gorutyn
	result := make(chan int)

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Podziel tablicę na fragmenty i przetwarzaj równolegle
	for i := 0; i < numWorkers; i++ {
		start := i * (len(arr) / numWorkers)
		end := (i + 1) * (len(arr) / numWorkers)
		if i == numWorkers-1 {
			end = len(arr)
		}
		go sumArrayPart(arr, start, end, &wg, result)
	}

	// Zbierz wyniki z poszczególnych gorutyn
	go func() {
		wg.Wait()
		close(result)
	}()

	totalSum := 0
	for partialSum := range result {
		totalSum += partialSum
	}

	fmt.Println("Suma elementów tablicy:", totalSum)
}
