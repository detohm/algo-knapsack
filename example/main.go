package main

import (
	"fmt"
	"time"

	"github.com/detohm/algo-knapsack"
)

func main() {

	dataSet := knapsack.DataSetLoader("s-3.txt")

	fmt.Println("Starting Depth-First Search (Bruteforce)...")
	start := time.Now()
	finalResult := knapsack.Bruteforce(dataSet)
	end := time.Since(start)
	fmt.Printf("Depth-First Search : Solved in %v \n", end)
	fmt.Println(finalResult)
	fmt.Println()

	fmt.Println("Starting Backtracking...")
	start = time.Now()
	finalResult = knapsack.Backtracking(dataSet)
	end = time.Since(start)
	fmt.Printf("Backtracking : Solved in %v \n", end)
	fmt.Println(finalResult)
	fmt.Println()

	fmt.Println("Starting Branch&Bound...")
	start = time.Now()
	finalResult = knapsack.BranchAndBound(dataSet)
	end = time.Since(start)
	fmt.Printf("Branch&Bound : Solved in %v \n", end)
	fmt.Println(finalResult)
	fmt.Println()

	fmt.Println("Starting BTBB...")
	start = time.Now()
	finalResult = knapsack.BTBB(dataSet)
	end = time.Since(start)
	fmt.Printf("BTBB : Solved in %v \n", end)
	fmt.Println(finalResult)
	fmt.Println()

}
