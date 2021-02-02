package tester

import (
	"bytes"
	"fmt"
	"github.com/deidelson/go-api-tester/pkg/util"
	"net/http"
	"sync"
	"time"
)

const (
	configPath                       = "./tester.json"
	defaultConcurrency               = 10
	defaultInterations               = 10
	defualtTimeInMSBetweenIterations = 10
)

type RequestSender interface {
	StressTest()
	IntervalStressTest()
}

type requestSender struct {
	stadistics *testerStadistic
	config     *TesterConfig
}

func NewRequestSender() RequestSender {
	config, err := CreateTesterConfigFromPath(configPath)
	if err != nil {
		panic("Error al cargar la configuración, revise el path y el archivo")
	}
	return &requestSender{
		stadistics: NewTesterStadistic(),
		config:     config,
	}
}

func (sender *requestSender) StressTest() {
	concurrency := util.ScanAsIntWithDefault("Seleccione la concurrencia (default 10): ", defaultConcurrency)

	sender.runTestWithConcurrency(concurrency, 1, 0)

	continuar := util.Scan("Para correr la prueba de nuevo presione 1 (cualquier otra tecla envia al inicio): ")
	if continuar == "1" {
		sender.StressTest()
	}
}

func (sender *requestSender) IntervalStressTest() {
	concurrency := util.ScanAsIntWithDefault("Seleccione la concurrencia (default 10): ", defaultConcurrency)
	iterations := util.ScanAsIntWithDefault("Seleccione la cantidad de iteraciones (default 10): ", defaultInterations)
	timeBetweenIterations := util.ScanAsIntWithDefault("Seleccione la cantidad de segundos entre iteraciones (default 10 segundos): ", defualtTimeInMSBetweenIterations)

	sender.runTestWithConcurrency(concurrency, iterations, timeBetweenIterations)

	continuar := util.Scan("Para correr la prueba de nuevo presione 1 (cualquier otra tecla envia al inicio): ")
	if continuar == "1" {
		sender.IntervalStressTest()
	}
}

func (sender *requestSender) runTestWithConcurrency(concurrency int, iterations int, timeBetweenIterations int) {
	sender.stadistics.resetResults()
	sender.stadistics.startCounting()
	var wg sync.WaitGroup
	wg.Add(concurrency*iterations)

	for j := 0; j < iterations; j++ {
		for i := 0; i < concurrency; i++ {
			go sender.sendRequest(sender.config, i, &wg)
		}
		time.Sleep(time.Duration(timeBetweenIterations) * time.Second)
	}

	wg.Wait()
	sender.stadistics.stopCounting()
	sender.stadistics.printStatistics()
}

func (sender *requestSender) sendRequest(config *TesterConfig, numeroIteracion int, wg *sync.WaitGroup) {
	fmt.Println("Ejecutando request numero: ", numeroIteracion)
	request, err := http.NewRequest(config.Method, config.Url, bytes.NewBuffer(config.getBodyAsByteArray()))
	if err != nil {
		fmt.Println("Error al crear la request ", err.Error())
		panic("Error al crear la request")
	}
	request.Header.Set(config.JwtHeader, config.JwtHeaderValue)
	request.Header.Set("Content-type", "application/json")

	client := http.Client{}

	response, err := client.Do(request)

	if err != nil {
		fmt.Println("Error en la request ", err.Error())
		sender.stadistics.addResult("ERROR " + err.Error())
		wg.Done()
		return
	}
	defer response.Body.Close()

	fmt.Println("Request numero", numeroIteracion, "Status", response.Status)
	sender.stadistics.addResult(response.Status)

	wg.Done()
}
