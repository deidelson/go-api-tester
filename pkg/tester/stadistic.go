package tester

import (
	"fmt"
	"time"
)

type testerStadistic struct {

	startTime time.Time
	testDuration time.Duration
	results map[string]int

}

func newTesterStadistic() *testerStadistic {
	return &testerStadistic{
		results: make(map[string]int),
	}
}

func (stadistic *testerStadistic) resetResults() {
	stadistic.results = make(map[string]int)
}

func (stadistic *testerStadistic) startCounting() {
	stadistic.startTime = time.Now()
}

func (stadistic *testerStadistic) stopCounting() {
	stadistic.testDuration = time.Since(stadistic.startTime)
}

func (stadistic *testerStadistic) addResult(result string) {
	count, present := stadistic.results[result]
	if present {
		stadistic.results[result] = count+1
	} else {
		stadistic.results[result] = 1
	}
}

func (stadistic *testerStadistic) printStatistics() {
	fmt.Println("")
	fmt.Println("Test Statistics")
	fmt.Println("----------------")
	fmt.Println("")
	fmt.Println("Duration: ", stadistic.testDuration)
	fmt.Println("")
	fmt.Println("Results: ")
	fmt.Println("")
	for result, duration := range stadistic.results {
		fmt.Println("Count of", result + ":", duration)
	}
	fmt.Println("")
}



