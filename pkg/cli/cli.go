package cli

import (
	"fmt"
	"github.com/deidelson/go-api-tester/pkg/tester"
	"os"
	"os/exec"
	"strconv"
)

const cantidad_intendos_default = 1;

type TesterCli interface {
	InitCli()
	Run()
	SelectConfig(newConfigPath string)
}

type testerCli struct {
	testerConfig *tester.TesterConfig
	tester tester.RequestSender
}

func NewTesterCli(path string) TesterCli {
	cliInstance := &testerCli{
		tester: tester.NewRequestSender(),
	}
	cliInstance.SelectConfig(path)

	return cliInstance
}

func (cli *testerCli) InitCli() {
	clear()
	fmt.Println("API Rest tester")
	fmt.Println("----------------------")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("1- Correr prueba")
	fmt.Println("0- Salir")
	fmt.Println("")

	choice := scan("Opcion: ")

	switch choice {
	case "1":
		cli.Run()
	case "0":
		os.Exit(9)
	default:
		cli.InitCli()
	}
}

func (cli *testerCli) Run() {
	clear()
	cantidadIntentosString := scan("Indique la cantidad de intentos (default 1) ")
	cantidadIntentos, err := strconv.Atoi(cantidadIntentosString)
	if err != nil {
		cantidadIntentos = cantidad_intendos_default
	}
	clear()
	fmt.Println("Se va a hacer un", cli.testerConfig.Method, "al endpoint", cli.testerConfig.Url, "con la cantidadIntentos:", cantidadIntentos)
	scan("Presione cualquier tecla para continuar")
	clear()

	cli.tester.Send(cli.testerConfig, cantidadIntentos)
	cli.InitCli()
}

func (cli *testerCli) SelectConfig(newConfigPath string) {
	config, err := tester.LoadConfig(newConfigPath)
	if err != nil {
		panic("Error al cargar la configuración, revise el path y el archivo")
	}
	cli.testerConfig = config;
}

func clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func scan(text string) string {
	fmt.Print(text)
	var input string
	fmt.Scanln(&input)
	return input
}

