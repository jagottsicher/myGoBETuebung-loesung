// ein Test, ob mathehelfer zu finden ist
package mathehelfer

func DieSummeVon(eingabeWerte ...int) int {

	summe := 0

	for _, v := range eingabeWerte {
		summe += v
	}

	return summe


}


func zweiGanzzahlenAddieren(a, b int) int {
	return a + b
}

func EineAndereArtSumme(eingabeWerte ...int) int {

	var summe int
	summe = eingabeWerte[0]

	for _, v := range eingabeWerte[1:] {
		summe = zweiGanzzahlenAddieren(summe, v)
	}

	return summe
}
