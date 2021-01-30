package cli

import (
	"fmt"
	"github.com/deidelson/go-api-tester/pkg/tester"
	"os"
	"os/exec"
)

type TesterCli interface {
	InitCli()
	Run()
	SelectConfig(newConfigPath string)
}

type testerCli struct {
	testerConfig *tester.TesterConfig
}

func NewTesterCli(path string) TesterCli {
	cliInstance := &testerCli{
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
	fmt.Println("1- RUN")
	fmt.Println("2- Cambiar configuración")
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
	tester.Post(cli.testerConfig, 10)
	scan("Presione cualquier tecla para continuar")
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
