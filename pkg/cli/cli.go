package cli

import (
	"fmt"
	"github.com/deidelson/go-api-tester/pkg/tester"
	"github.com/deidelson/go-api-tester/pkg/util"
	"os"
)



type TesterCli interface {
	InitCli()
	Run()
}

type testerCli struct {
	tester tester.RequestSender
}

func NewTesterCli() TesterCli {
	cliInstance := &testerCli{
		tester: tester.NewRequestSender(),
	}
	return cliInstance
}

func (cli *testerCli) InitCli() {
	util.ClearConsole()
	fmt.Println("API Rest tester")
	fmt.Println("----------------------")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("1- Test de estres")
	fmt.Println("2- Test de estres con intervalos")
	fmt.Println("0- Salir")
	fmt.Println("")

	choice := util.Scan("Opcion: ")

	switch choice {
	case "1":
		util.ClearConsole()
		cli.tester.StressTest()
		cli.InitCli()
	case "2":
		util.ClearConsole()
		cli.tester.IntervalStressTest()
		cli.InitCli()
	case "0":
		os.Exit(9)
	default:
		cli.InitCli()
	}
}

func (cli *testerCli) Run() {
	util.ClearConsole()
	cli.tester.StressTest()
	cli.InitCli()
}



