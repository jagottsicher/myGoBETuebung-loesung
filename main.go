package main

import (
	"fmt"
	"mathehelfer"
)

func main() {

	eingabeWerteSlice := []int{1, 2, 3, 4, 5, 6}

	fmt.Println("Summe:", mathehelfer.DieSummeVon(eingabeWerteSlice...))
	// fmt.Println("Summe:", mathehelfer.EineAndereArtSumme(eingabeWerteSlice...))
}
