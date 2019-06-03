package main

import (
	"flag"

	"github.com/bodhi369/echoatom/pkg/api"
	"github.com/bodhi369/echoatom/pkg/utl/config"
)

func main() {

	cfgPath := flag.String("p", "D:/GoProject/go/src/github.com/bodhi369/echoatom/cmd/api/conf.local.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	checkErr(err)

	checkErr(api.Start(cfg))
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
