package tester

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
	"time"
)


//TODO armar objeto para ver estadisticas
func Post(config *TesterConfig, cantidadDeRequest int) {
	var wg sync.WaitGroup
	wg.Add(cantidadDeRequest)

	start := time.Now()
	for i:=0; i < cantidadDeRequest; i++ {
		go sendPedido(config, i, &wg)
	}
	wg.Wait()
	duracion := time.Since(start)
	fmt.Printf("El proceso domoró %s", duracion)
}

func sendPedido(config *TesterConfig, numeroIteracion int,  wg *sync.WaitGroup) {
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
		return
	}

	fmt.Println("Request numero", numeroIteracion , "Status", response.StatusCode)


	wg.Done()

}

