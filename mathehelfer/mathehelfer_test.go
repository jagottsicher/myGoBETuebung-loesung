package mathehelfer

import "testing"

func BenchmarkDieSummeVon(b *testing.B) {

	sliceOfInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		DieSummeVon(sliceOfInt...)
	}

}
