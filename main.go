package main

import (
	"fmt"
	"log"
)

func main() {
	// n := 0
	var v string
	fmt.Print("Введите целое число: ")
	_, err := fmt.Scan(&v)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Вы ввели число: %v\n", v)
}
