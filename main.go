package main

import (
	"github.com/fbriansyah/go-printer/printing"
)

func main() {
	printer := printing.NewPrinter()

	printer.Print([]string{"Haoooooo11\nnama:\tFebrian", "Haoooooo22", "Hai Hai"})
}
