package tester

import (
	"bytes"
	"fmt"
	"github.com/deidelson/go-api-tester/pkg/util/osutil"
	"github.com/deidelson/go-api-tester/pkg/util/web"
	"log"
	"sync"
	"time"
)

const (
	configPath                       = "./tester.json"
	defaultConcurrency               = 10
	defaultIterations                = 10
	defualtTimeInMSBetweenIterations = 10
)

type RequestSender interface {
	StressTest()
	IntervalStressTest()
}

type requestSender struct {
	stadistics *testerStadistic
	config     *TesterConfig
	httpService web.HttpService
}

func NewRequestSender() RequestSender {
	config, err := CreateTesterConfigFromPath(configPath)
	if err != nil {
		panic("Error loading configuration, check file path and format")
	}
	return &requestSender{
		stadistics:  newTesterStadistic(),
		config:      config,
		httpService: web.NewHttpService(),
	}
}

func (sender *requestSender) StressTest() {
	concurrency := osutil.ScanAsIntWithDefault("Concurrency (default 10): ", defaultConcurrency)

	sender.runTestWithConcurrency(concurrency, 1, 0)

	continuar := osutil.Scan("Press 1 to run again (other key send to main menu): ")
	if continuar == "1" {
		sender.StressTest()
	}
}

func (sender *requestSender) IntervalStressTest() {
	concurrency := osutil.ScanAsIntWithDefault("Concurrency (default 10): ", defaultConcurrency)
	iterations := osutil.ScanAsIntWithDefault("Iterations (default 10): ", defaultIterations)
	timeBetweenIterations := osutil.ScanAsIntWithDefault("Seconds between each iteration (default 10 seconds): ", defualtTimeInMSBetweenIterations)

	sender.runTestWithConcurrency(concurrency, iterations, timeBetweenIterations)

	runAgain := osutil.Scan("Press 1 to run again (other key send to main menu): ")
	if runAgain == "1" {
		sender.IntervalStressTest()
	}
}

func (sender *requestSender) runTestWithConcurrency(concurrency int, iterations int, timeBetweenIterations int) {
	sender.stadistics.resetResults()
	sender.stadistics.startCounting()
	var wg sync.WaitGroup
	wg.Add(concurrency*iterations)
	resultsChannel := make(chan string)

	for j := 0; j < iterations; j++ {
		log.Println("Iteration: ", j)
		for i := 0; i < concurrency; i++ {
			go sender.sendRequest(sender.config, &wg, resultsChannel)
		}
		time.Sleep(time.Duration(timeBetweenIterations) * time.Second)
	}

	go func() {
		log.Println("Waiting for response processing...")
		wg.Wait()
		close(resultsChannel)
	}()

	for result := range resultsChannel {
		sender.stadistics.addResult(result)
	}
	sender.stadistics.stopCounting()
	sender.stadistics.printStatistics()
}

func (sender *requestSender) sendRequest(config *TesterConfig, wg *sync.WaitGroup, results chan string) {
	defer wg.Done()
	response, err := sender.httpService.Send(config.Method, config.Url, bytes.NewBuffer(config.getBodyAsByteArray()), config.Headers)
	if response != nil && response.Body != nil {
		err := response.Body.Close()
		if err != nil {
			fmt.Println("Error closing body", err.Error())
		}
	}

	if err != nil {
		results <- "ERROR " + err.Error()
		return
	}

	results <- response.Status
}
