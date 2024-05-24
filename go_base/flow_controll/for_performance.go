package main

import (
	"fmt"
	"time"
)

type S [24]int64

var sX = make([]S, 1000)
var sY = make([]S, 1000)
var sZ = make([]S, 1000)
var sumX, sumY, sumZ int64

func loop() {
	sumX = 0
	for j := 0; j < len(sX); j++ {
		sumX += sX[j][0]
	}
}

func rangeOneIterVar() {
	sumY = 0
	for j := range sY {
		sumY += sY[j][0]
	}
}

func rangeTwoIterVar() {
	sumZ = 0
	for _, v := range sZ {
		sumZ += v[0]
	}
}

func measureTime(f func(), iterations int) time.Duration {
	start := time.Now()
	for i := 0; i < iterations; i++ {
		f()
	}
	return time.Since(start)
}

func main() {
	iterations := 100000
	timeLoop := measureTime(loop, iterations)
	timeRangeOneIterVar := measureTime(rangeOneIterVar, iterations)
	timeRangeTwoIterVar := measureTime(rangeTwoIterVar, iterations)

	fmt.Printf("Loop: %v\n", timeLoop)                                         //140.1314ms
	fmt.Printf("Range with one iterator variable: %v\n", timeRangeOneIterVar)  //149.5651ms
	fmt.Printf("Range with two iterator variables: %v\n", timeRangeTwoIterVar) //418.814ms
}
