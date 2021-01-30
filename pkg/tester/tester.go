package tester

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func Post(config *TesterConfig, cantidadDeRequest int) {
	var wg sync.WaitGroup
	wg.Add(cantidadDeRequest)

	for i:=0; i < cantidadDeRequest; i++ {
		go sendPedido(config, i, &wg)
	}
	wg.Wait()
}

func sendPedido(config *TesterConfig, numeroIteracion int,  wg *sync.WaitGroup) {
	fmt.Println("Ejecutando request numero: ", numeroIteracion)
	requestBody, err := ioutil.ReadFile(config.JsonBodyPath)

	if err != nil {
		panic("Error critico")
	}
	request, err := http.NewRequest(config.Method, config.Url,  bytes.NewBuffer(requestBody))
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
		panic(err.Error())
	}

	defer response.Body.Close()

	fmt.Println("Request numero", numeroIteracion , "Status", response.StatusCode)


	wg.Done()

}

