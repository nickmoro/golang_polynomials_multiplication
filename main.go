package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
)

func multiplyPols(pol1, pol2 []int) []int {
	n1, n2 := len(pol1), len(pol2)
	res := make([]int, n1+n2-1)
	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			res[i+j] += pol1[i] * pol2[j]
		}
	}
	return res
}

func forGoroutine(polyns [][]int, left int, right int, res *[]int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Will calculate from %d to %d:\n", left, right)
	for i := left; i <= right; i++ {
		fmt.Println(polyns[i])
	}
	*res = polyns[left]
	for cur := left + 1; cur <= right; cur++ {
		*res = multiplyPols(*res, polyns[cur])
	}
	fmt.Println("Returning res =", *res)
}

func solve(polyns [][]int, availableThreads int) []int {
	curPolyns := polyns
	for len(curPolyns) > 1 {
		fmt.Println("curPolyns =", curPolyns)
		polynsNForThr := (len(curPolyns) + availableThreads - 1) / availableThreads
		if polynsNForThr < 2 {
			polynsNForThr = 2
		}
		fmt.Println("polynsNForThr =", polynsNForThr)
		wg := new(sync.WaitGroup)
		usedThreads := (len(curPolyns) + polynsNForThr - 1) / polynsNForThr
		fmt.Println("usedThreads =", usedThreads)
		newPolyns := make([][]int, usedThreads)
		newPolynsIndex := 0
		for left := 0; left < len(curPolyns); left += polynsNForThr {
			right := left + polynsNForThr - 1
			if right >= len(curPolyns) {
				right = len(curPolyns) - 1
			}
			wg.Add(1)
			go forGoroutine(curPolyns, left, right, &newPolyns[newPolynsIndex], wg)
			newPolynsIndex++
		}
		fmt.Println("Waiting")
		wg.Wait()
		fmt.Println("Waited, get newPolyns =", newPolyns)
		curPolyns = newPolyns
	}
	return curPolyns[0]
}

// TODO
func PolynToString(polyn []int) string {
	var res string

	return res
}

func main() {
	// Input
	availableThreads := runtime.GOMAXPROCS(runtime.NumCPU())
	if availableThreads < 1 {
		log.Fatal("No available threads")
	}
	var polynsNumber int
	for polynsNumber < 1 {
		fmt.Println("Enter number of polyns:")
		fmt.Scan(&polynsNumber)
	}
	polyns := make([][]int, polynsNumber)
	// TODO: func StringToPolyn
	for i := 0; i < polynsNumber; i++ {
		fmt.Println("Enter degree of the polyn:")
		var degree int
		fmt.Scan(&degree)
		degree++
		polyns[i] = make([]int, degree)
		fmt.Println("Enter polyn's coefficients (via space):")
		for j := degree - 1; j >= 0; j-- {
			fmt.Scan(&polyns[i][j])
		}
	}
	fmt.Println("Polyns:\n", polyns)

	// Calculations
	ans := solve(polyns, availableThreads)

	// Output
	fmt.Println("Answer:")
	fmt.Println(ans)
	// TODO: func PolynToString
	fmt.Println(PolynToString(ans))
}
