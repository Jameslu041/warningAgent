package main

import (
	"flag"
	"warningAgent/src/config"
	"warningAgent/src/funcs"
)

var (
	cfg = flag.String("config", "", "warningAgent config file path.")
)

func main() {
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	c := make(chan int)
	funcs.BuildLopper(5)
	funcs.StartCollect()
	<-c

}
