package tester

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
)

type RequestSender interface {
	Send(config *TesterConfig, concurrency int)
}

type requestSender struct {
	stadistics *testerStadistic
}

func NewRequestSender() RequestSender {
	return &requestSender{
		stadistics: NewTesterStadistic(),
	}
}

func (sender *requestSender) Send(config *TesterConfig, concurrency int) {
	var wg sync.WaitGroup
	wg.Add(concurrency)

	sender.stadistics.resetResults()
	sender.stadistics.startCounting()

	for i:=0; i < concurrency; i++ {
		go sender.sendRequest(config, i, &wg)
	}
	wg.Wait()

	sender.stadistics.stopCounting()
	sender.stadistics.printStatidistics()

	continuar := scan("Para correr la prueba de nuevo presione 1 (cualquier otra tecla envia al inicio): ")
	if continuar == "1" {
		sender.Send(config, concurrency)
	}

}

func (sender *requestSender) sendRequest(config *TesterConfig, numeroIteracion int,  wg *sync.WaitGroup) {
	fmt.Println("Ejecutando request numero: ", numeroIteracion)
	request, err := http.NewRequest(config.Method, config.Url,  bytes.NewBuffer(config.getBodyAsByteArray()))
	if err != nil {
		fmt.Println("Error al crear la request ", err.Error())
		panic("Error al crear la request")
	}
	request.Header.Set(config.JwtHeader, config.JwtHeaderValue)
	request.Header.Set("Content-type", "application/json")

	client := http.Client{

	}

	response, err := client.Do(request)

	if err != nil {
		fmt.Println("Error en la request ", err.Error())
		sender.stadistics.addResult("ERROR "+err.Error())
		wg.Done()
		return
	}
	defer response.Body.Close()

	fmt.Println("Request numero", numeroIteracion , "Status", response.Status)
	sender.stadistics.addResult(response.Status)


	wg.Done()

}

func scan(text string) string {
	fmt.Print(text)
	var input string
	fmt.Scanln(&input)
	return input
}

