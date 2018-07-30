package knapsack

import (
	"fmt"
	"strings"
	"time"
)

//Solve - Solve with dataset and specified algorithm
func Solve(ds DataSet, algorithm string) {

	fmt.Printf("Starting %s ...\n", algorithm)
	start := time.Now()

	finalResult := FinalResult{}

	switch strings.ToLower(algorithm) {
	case "bruteforce": //or depth first search
		finalResult = Bruteforce(ds)
		break
	case "backtracking":
		finalResult = Backtracking(ds)
		break
	case "branchandbound":
		finalResult = BranchAndBound(ds)
		break
	case "btbb":
		finalResult = BTBB(ds)
		break
	}

	end := time.Since(start)
	fmt.Printf("%s : Solved in %v \n", algorithm, end)
	fmt.Printf("value: %d, answer: %s, traversed %d nodes \n\n",
		finalResult.optimalValue,
		finalResult.solution,
		finalResult.countNode)
	//fmt.Println(finalResult)

}
