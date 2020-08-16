package main //
import "C"   // C import edilmiş olmak zorunda!
import (
	"fmt"
	"log"
)

func main() {} // main boş olacak

func merhaba() { // Bu arkadaş paylaşımlı fonksiyon değil . Go 'da paylaşımlı fonksiyonlar büyük harf ile başlar . Bu kullanım yanlış
	fmt.Println("Merhaba Dünya")
}

//"export <funcAdi>" yapısını kullanmak zorundayız !

//export Merhaba
func Merhaba() { // Fonksiyonumuz paylaşımlı oldu (büyük harf ile başladığı için)
	fmt.Println("Merhaba Dünya !")
}

//export MerhabaUn
func MerhabaUn() string {
	return "Merhaba Dünya" // Doğrudan string'i kullanamayız
}

//export PrintUn
func PrintUn(text string) { // Bu da örnek için , so kütüphanelerinde doğrudan String'i kullanamayız
	log.Println(text)
}

// Derleme :

// go build -builmode=c-shared sözDizimi.go
// sözDizimi ve sözDizimi.h adında dosyalar oluşacak
// Oluşan so dosyasının sadece aynı işletim sistemine ve mimarisine sahip cihazlarda kullanbileceğini unutmayın
