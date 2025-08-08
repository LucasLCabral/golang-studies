package main

import (
	"fmt"

	"github.com/LucasLCabral/golang-studies/7-Packaging/1/math"
)

func main() {
	m := math.NewMath(1, 2)
	fmt.Println(m.Add())
}
