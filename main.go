package main

import (
	"fmt"
	"log"
)

func main() {
	var v string
	fmt.Print("Введите данные: ")
	_, err := fmt.Scan(&v)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Вы ввели следующие данные: %v\n", v)
}
