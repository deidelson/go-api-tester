package cli

import (
	"fmt"
	"github.com/deidelson/go-api-tester/pkg/tester"
	"github.com/deidelson/go-api-tester/pkg/util/osutil"
	"os"
)



type TesterCli interface {
	InitCli()
}

type testerCli struct {
	tester tester.RequestSender
}

func NewTesterCli() TesterCli {
	//TODO Create factory and inject
	cliInstance := &testerCli{
		tester: tester.NewRequestSender(),
	}
	return cliInstance
}

func (cli *testerCli) InitCli() {
	osutil.ClearConsole()
	fmt.Println("API Rest tester v0.1.1")
	fmt.Println("----------------------")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("1- Simple stress test")
	fmt.Println("2- Interval stress test")
	fmt.Println("0- Exit")
	fmt.Println("")

	choice := osutil.Scan("Option: ")

	switch choice {
	case "1":
		osutil.ClearConsole()
		cli.tester.StressTest()
		cli.InitCli()
	case "2":
		osutil.ClearConsole()
		cli.tester.IntervalStressTest()
		cli.InitCli()
	case "0":
		os.Exit(9)
	default:
		cli.InitCli()
	}
}



