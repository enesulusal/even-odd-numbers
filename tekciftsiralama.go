package main

import "fmt"

func main() {

var sayilar [10]int = [10]int {31, 5, 19, 94, 17, 3, 8, 72, 86, 10}

for i:=0 ; i < len(sayilar) ; i++{

if sayilar[i] % 2 == 0 {
    fmt.Println("ciftSayi = ", sayilar[i] )
}else{
    fmt.Println("tekSayi = ", sayilar[i] )

}
}

var siralama int

siralama = arrangement(siralama)
func arrangement(toplanmaYeri []int) (toplanmaYerıIkı []int) {
    var buyukSayi int
    var kucukSayi int
    var index int


    for j := 0 ; j < len(toplanmaYeri) ; j++{


        kucukSayi = toplanmaYeri[j]

        for i := 0; i < len(toplanmaYeri); i++{

            if toplanmaYeri[i] < kucukSayi{

                kucukSayi = toplanmaYeri[i]

                index = i 

            }

        }

        toplanmaYerıIkı = append(toplanmaYerıIkı, kucukSayi)

        buyukSayi = funchiki (toplanmaYeri)

        toplanmaYeri[index] = buyukSayi
    }
         return toplanmaYerıIkı
}

func funchiki (toplanmaYeri []int) (buyukSayi []int) {
    buyukSayi = 0
    for i := 0; i < len(toplanmaYeri); i++ {
        if toplanmaYeri[i] > buyukSayi {

            buyukSayi = toplanmaYeri[i]

        }

    }

    return buyukSayi
}