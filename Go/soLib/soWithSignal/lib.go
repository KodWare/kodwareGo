package main

import "C"
import (
	"log"
	"os"
	"os/signal"
)

func main() {}

func run(kanal chan os.Signal) {
	for {
		log.Printf("\t \033[31m'%v'\033[0m Savuşturuldu !\n", <-kanal)
	} // infloop
}

//export HandleSignals
func HandleSignals() {
	kanal := make(chan os.Signal)
	signal.Notify(kanal)
	go run(kanal)
}
