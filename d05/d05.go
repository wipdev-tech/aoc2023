package main

import "fmt"

type rule struct {
	srcStart int
	srcEnd   int
	dstStart int
	dstEnd   int
}

type numRange struct {
	start int
	end   int
}

func main() {
	// seeds, rulesets := parseAlmanac("d05/d05test.txt")
	seeds, rulesets := parseAlmanac("d05/d05input.txt")

	ranges := []numRange{}
	for i := 0; i < len(seeds)-1; i += 2 {
		ranges = append(ranges, numRange{seeds[i], seeds[i] + seeds[i+1] - 1})
	}

	allRanges := []numRange{}
	for _, rng := range ranges {
		calcRanges := []numRange{rng}
		for _, rs := range rulesets {
			nRanges := len(calcRanges)
			for j := 0; j < nRanges; j++ {
				result := applyRuleset(calcRanges[0], rs)
				calcRanges = calcRanges[1:]
				calcRanges = append(calcRanges, result...)
			}

			fmt.Println(calcRanges)
		}
		allRanges = append(allRanges, calcRanges...)
	}
	fmt.Println(allRanges)

	var min int
	for i, r := range allRanges {
		if i == 0 || r.start < min {
			min = r.start
		}
	}

	fmt.Println("\nmin:", min)
}
