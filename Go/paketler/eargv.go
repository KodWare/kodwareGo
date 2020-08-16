package main

import (
	"flag"
	"fmt"
)

func main() {
	isbool := flag.Bool("bool", false, "bool true-false veri alır .") // varsayılan: false |--bool *<herhangibirdeger>  True |
	str := flag.String("str", "merhabadünya", "str string veri alır") // varsayılan: merhabadünya | --str "<string veri>" "<string veri>"

	flag.Parse()

	fmt.Println("bool", *isbool, "\nstr  ", *str)

}
