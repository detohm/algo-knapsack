package knapsack

//Backtracking - eliminate the paths which exceed the limited weight
func Backtracking(ds DataSet) FinalResult {
	fr := FinalResult{}
	backtrackingDFS(&ds, &fr, "")
	return fr
}

func backtrackingDFS(ds *DataSet, fr *FinalResult, pathStr string) {
	fr.countNode++

	/*---Backtracking---*/
	tempSumWeight := 0
	for i := 0; i < len(pathStr); i++ {
		if pathStr[i] == '1' {
			tempSumWeight += ds.weights[i]
		}
	}

	//eliminate traversing the useless nodes
	if tempSumWeight > ds.limitedWeight {
		return
	}
	/*---End Backtracking---*/

	if len(pathStr) == len(ds.weights) {

		sumWeight := 0
		sumValue := 0
		for i := 0; i < len(pathStr); i++ {
			if pathStr[i] == '1' {
				sumWeight += ds.weights[i]
				sumValue += ds.values[i]
			}
		}

		//update optimal solution if found
		if sumWeight <= ds.limitedWeight && sumValue > fr.optimalValue {
			fr.weight = sumWeight
			fr.optimalValue = sumValue
			fr.solution = pathStr
		}
		return

	}

	backtrackingDFS(ds, fr, pathStr+"0")
	backtrackingDFS(ds, fr, pathStr+"1")
}
