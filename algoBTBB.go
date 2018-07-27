package knapsack

//BTBB - combine branch&bound and backtracking together
func BTBB(ds DataSet) FinalResult {
	fr := FinalResult{}
	bTBBDFS(&ds, &fr, "")
	return fr
}

func bTBBDFS(ds *DataSet, fr *FinalResult, pathStr string) {
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

	/*---Branch&Bound---*/
	sumCurrentValue := 0
	for i := 0; i < len(pathStr); i++ {
		if pathStr[i] == '1' {
			sumCurrentValue += ds.values[i]
		}
	}
	sumRemainingValue := 0
	for i := len(ds.values) - 1; i > len(pathStr)-1; i-- {
		sumRemainingValue += ds.values[i]
	}
	if fr.optimalValue > sumCurrentValue+sumRemainingValue {
		return
	}
	/*---End Branch&Bound---*/

	bTBBDFS(ds, fr, pathStr+"0")
	bTBBDFS(ds, fr, pathStr+"1")
}
