package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	leapYear()
	fmt.Println("\n--------------------------------")
	arrayToObject()
	fmt.Println("\n--------------------------------")
	itemNumbers()
	fmt.Println("\n--------------------------------")
	playJson()
}

func leapYear() {
	var tahun int
	fmt.Println("Tahun Kabisat")
	fmt.Print("Masukkan tahun : ")
	fmt.Scanln(&tahun)
	if tahun%4 == 0 && tahun%100 != 0 || tahun%400 == 0 {
		fmt.Println("Tahun", tahun, "adalah tahun kabisat")
	} else {
		fmt.Println("Tahun", tahun, "bukan tahun kabisat")
	}
}

func arrayToObject() {
	type identitas struct {
		namadepan    string
		namabelakang string
	}

	orang := identitas{
		namadepan:    "Adhi",
		namabelakang: "Ardiansyah",
	}

	fmt.Println()
	fmt.Println("Array ke Object")
	fmt.Println(orang)
}

func itemNumbers() {
	bil := []int{2, 3, 5, 7, 11, 12, 32, 45, 65, 90, 70, 80, 101}

	fmt.Println()
	fmt.Println("Daftar bilangan =", bil)

	fmt.Print("Bilangan Genap = ")
	for i := 0; i < len(bil); i++ {
		if bil[i]%2 == 0 {
			fmt.Print(bil[i], " ")
		}
	}

	fmt.Println()
	fmt.Print("Bilangan Ganjil = ")
	for i := 0; i < len(bil); i++ {
		if bil[i]%2 != 0 {
			fmt.Print(bil[i], " ")
		}
	}

	fmt.Println()
	fmt.Print("Bilangan Kelipatan 5 = ")
	for i := 0; i < len(bil); i++ {
		if bil[i]%5 == 0 {
			fmt.Print(bil[i], " ")
		}
	}

	fmt.Println()
	fmt.Print("Bilangan Prima = ")
	for i := 0; i < len(bil); i++ {
		var isPrime bool = true

		for j := 2; j < i; j++ {
			if bil[i]%j == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			fmt.Print(bil[i], " ")
		}
	}

	fmt.Println()
	fmt.Print("Bilangan Prima kurang dari 100 = ")
	for i := 0; i < len(bil); i++ {
		var isPrime bool = true

		for j := 2; j < i; j++ {
			if bil[i]%j == 0 {
				isPrime = false
				break
			}
		}

		if isPrime && bil[i] < 100 {
			fmt.Print(bil[i], " ")
		}
	}
	fmt.Println()
}

func playJson() {
	fmt.Println()
	fmt.Println("Bermain dengan JSON")

	type User struct {
		NamaLengkap string `json:"Nama"`
		Umur        int
	}

	var jsonString = `{"Nama": "Adhi Ardiansyah", "Umur": 20}`
	var jsonData = []byte(jsonString)

	var data User

	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Nama Lengkap :", data.NamaLengkap)
	fmt.Println("Umur  :", data.Umur)
}
