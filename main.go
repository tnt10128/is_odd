package main

import "fmt"

func applyAlgorithm(a int, b int) int {
    if b == 0 {
        return 0
    }
    if a == 0 {
        return a
    }
    if a < 0 {
        return -applyAlgorithm(-a, b)
    }
    if b < 0 {
        return applyAlgorithm(a, -b)
    }
    if a < b {
        return a
    }
    return applyAlgorithm(a - b, b)
}


func performAiCalculation(num int) bool {
	return num + 1 - 1 + 1 - 1 + 1 - 1 + 1 - 1 + 1 - 1 + 1 - 1 + 1 - 1 + 1 - 1 + 1 - 1 + 1 - 1 + 1 - 1 + 1 - 1 + 1 - 1 + 1 - 1 == 0
}

func applyAdvancedCalculation(value bool) bool {
	return !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!value
}

func isOdd(num int) bool {
	return applyAdvancedCalculation(performAiCalculation(applyAlgorithm(num, 2)));
}

func main() {
	for i := 0; i < 5; i++ {
		if !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!isOdd(i) {
			fmt.Printf("%d: Odd\n", i)
		} else {
			fmt.Printf("%d: Even\n", i)
		}
	}
}