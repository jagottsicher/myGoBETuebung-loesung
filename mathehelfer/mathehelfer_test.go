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

func TestDieSummeVon(t *testing.T) {
	sliceOfInt := []int{1, 2, 3}
	// Alternative
	// wrongOutput := fmt.Sprintf("Summe: %v", DieSummeVon(sliceOfInt...))
	// if wrongOutput != "Summe: 6" {
	wrongOutput := fmt.Sprintln("Summe:", DieSummeVon(sliceOfInt...))
	if wrongOutput != "Summe: 6\n" {
		t.Error("Expected:", "Summe: 6", "got:", wrongOutput)
	}
}

func TestEineAndereArtSumme(t *testing.T) {
	sliceOfInt := []int{1, 2, 3}
	// Alternative
	// wrongOutput := fmt.Sprintf("Summe: %v", EineAndereArtSumme(sliceOfInt...))
	// if wrongOutput != "Summe: 6" {
	wrongOutput := fmt.Sprintln("Summe:", EineAndereArtSumme(sliceOfInt...))
	if wrongOutput != "Summe: 6\n" {
		t.Error("Expected:", "Summe: 6", "got:", wrongOutput)
	}
}

func TestDieSummeVonTableTest(t *testing.T) {

	type testData struct {
		inputInts []int
		ergebnis  int
	}

	tests := []testData{
		testData{[]int{1, 2, 3, 4, 5, 6}, 21},
		testData{[]int{1, 2, 4, 8, 16}, 31},
		testData{[]int{1, -2, 3, -4, 5}, 3},
		testData{[]int{-5, 1, 1, 1, 1, 1}, 0},
		testData{[]int{6, 5, 4, 3, 2, 1}, 21},
	}

	for _, v := range tests {
		v1 := DieSummeVon(v.inputInts...)
		if v1 != v.ergebnis {
			t.Error("Expected", v.ergebnis, "got", v1)
		}
	}
}

func TestEineAndereArtSummeTableTest(t *testing.T) {

	type testData struct {
		inputInts []int
		ergebnis  int
	}

	tests := []testData{
		testData{[]int{1, 2, 3, 4, 5, 6}, 21},
		testData{[]int{1, 2, 4, 8, 16}, 31},
		testData{[]int{1, -2, 3, -4, 5}, 3},
		testData{[]int{-5, 1, 1, 1, 1, 1}, 0},
		testData{[]int{6, 5, 4, 3, 2, 1}, 21},
	}

	for _, v := range tests {
		v1 := EineAndereArtSumme(v.inputInts...)
		if v1 != v.ergebnis {
			t.Error("Expected", v.ergebnis, "got", v1)
		}
	}
}
