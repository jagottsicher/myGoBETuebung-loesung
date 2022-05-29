// mathehelfee ist eine Package, dass nützliche Funktionen für Mathe zur Verfügung stellt
package mathehelfer

// DieSummeVon liefert die Summe über beliebige Anzahl von Ganzzahlen
// Dazu wird in einer range immer weiter aufaddiert
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

// EineAndereArtSumme liefert die Summe über beliebige Anzahl von Ganzzahlen
// Dazu wird in einer range immer eine Summe aus Rückgabewert aus einer Funktion
// mit nur zwei Summanden aufaddiert
func EineAndereArtSumme(eingabeWerte ...int) int {

	var summe int
	summe = eingabeWerte[0]

	for _, v := range eingabeWerte[1:] {


		summe = zweiGanzzahlenAddieren(summe, v)
	}

	return summe
}