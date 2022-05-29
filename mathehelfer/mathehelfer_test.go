package mathehelfer

import (
	"fmt"
	"testing"
)

func BenchmarkDieSummeVon(b *testing.B) {

	sliceOfInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		DieSummeVon(sliceOfInt...)
	}
}

func BenchmarkEineAndereArtSumme(b *testing.B) {

	sliceOfInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		EineAndereArtSumme(sliceOfInt...)
	}
}

func ExampleDieSummeVon() {
	sliceOfInt := []int{1, 2, 3}
	fmt.Println("Summe:", DieSummeVon(sliceOfInt...))
	// Output:
	// Summe: 6
}

func ExampleEineAndereArtSumme() {
	fmt.Println("Summe:", EineAndereArtSumme(1, 2, 3))
	// Output:
	// Summe: 6
}
