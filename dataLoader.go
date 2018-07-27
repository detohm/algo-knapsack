package knapsack

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

//DataSetLoader - Load data from file to data struct
func DataSetLoader(filename string) DataSet {

	file, err := os.Open("./dataset/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	ds := DataSet{}

	isLimitedWeightAdded := false

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) != "" {

			if !isLimitedWeightAdded {
				ds.limitedWeight, _ = strconv.Atoi(line)
				isLimitedWeightAdded = true
				continue
			}
			splits := strings.Split(line, " ")
			weight, _ := strconv.Atoi(splits[0])
			value, _ := strconv.Atoi(splits[1])
			ds.weights = append(ds.weights, weight)
			ds.values = append(ds.values, value)

		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return ds
}
