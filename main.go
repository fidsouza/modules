package main

import (
	"fmt"
	"modules/cars"
	"modules/mathematic"
)

func main() {
	summarize := mathematic.Sum(20, 30)
	fmt.Println("Resultados: ", summarize)
	vehchile := cars.Car{Brand: "Fiat"}

	fmt.Println(vehchile.Run())

}
