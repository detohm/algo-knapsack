package knapsack

//DataSet for input
type DataSet struct {
	limitedWeight int
	weights       []int
	values        []int
}

//FinalResult from each algorithm
type FinalResult struct {
	weight       int
	optimalValue int
	solution     string
	countNode    int
}
