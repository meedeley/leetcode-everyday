package main

func TwoSum(numbs []int, target int) []int {
	for i := 0; i < len(numbs); i++ {
		for j := i+1; j < len(numbs); j++ {			
			if numbs[i] + numbs[j] == target {
				value := []int{i, j}
				return value;
			}
		}	
	}

	return []int{}

}