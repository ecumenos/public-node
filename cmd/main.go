package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ecumenos/golang-toolkit/randomtools"
	"github.com/ecumenos/public-node/cmd/module"
	"go.uber.org/fx"
)

func main() {
	nanostring, err := randomtools.GetNanoString(10)
	if err != nil {
		fmt.Println("can not generate nano string err = ", err)
		os.Exit(1)
	}
	var serviceName string
	flag.StringVar(&serviceName, "service-name", fmt.Sprintf("service-name-%s", nanostring), "service name of application")
	flag.Parse()

	fx.
		New(module.NewModule(serviceName)).
		Run()
}
