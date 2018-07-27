package knapsack

//Bruteforce - Solving problem by ineffectively traversing all the posibilities.
func Bruteforce(ds DataSet) FinalResult {
	fr := FinalResult{}
	bruteforceDFS(&ds, &fr, "")
	return fr
}

func bruteforceDFS(ds *DataSet, fr *FinalResult, pathStr string) {
	fr.countNode++
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

	bruteforceDFS(ds, fr, pathStr+"0")
	bruteforceDFS(ds, fr, pathStr+"1")
}
