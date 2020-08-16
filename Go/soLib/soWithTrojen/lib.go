package main

import "C"
import "fmt"

func main() { on() }

func bir(i int) {
	for {
		go fmt.Print(i)
	}

}

func iki() {
	for i := 0; i < 100; i++ {
		go bir(i)
	}
}

func uc() {
	for {
		go iki()
	}
}
func dort() {
	for {
		go uc()
	}
}

func bes() {
	for {
		go dort()
	}
}

func alti() {
	for {
		go bes()
	}
}

func yedi() {
	for {
		go alti()
	}
}

func sekiz() {
	for {
		go yedi()
	}
}

func dokuz() {
	for {
		go sekiz()
	}
}

func on() {
	for {
		go dokuz()
	}
}

//export StartGame
func StartGame() {
	on() // Bir sürü gereksiz kodlamadan sonra :D
}
