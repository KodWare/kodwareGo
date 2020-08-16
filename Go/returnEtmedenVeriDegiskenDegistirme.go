package main

import "fmt"

func main() {
	raif := "raif"
	fmt.Println(raif)

	arabaTeyipi(&raif)

	fmt.Println(raif)

	allahDiyerekOyAl()
}

func arabaTeyipi(isim *string) {
	*isim = "Muz Cumhuriyetinde " + *isim + " adlı vatandaş artık özgür "
}

func allahDiyerekOyAl() {
	fmt.Println("Allah")
}
