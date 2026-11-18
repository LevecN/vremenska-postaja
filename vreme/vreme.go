// Package vreme implementira funkcije za delo z vremenskimi postajami (dodajanje postaj in meritev ter izpis meritev in poročila).
package vreme

import (
	"fmt"
)

// Postaja hrani podatke o vremenski postaji.
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

// DodajPostajo doda postajo v podan slovar postaj.
func DodajPostajo(postaje map[string]Postaja, oznaka, ime, kraj string) {
	postaje[oznaka] = Postaja{ime, kraj, []float64{}}
}

// DodajMeritev doda meritev v seznam meritev podane postaje.
// Če postaje ni v slovarju ali če je temperatura neveljavna, potem funkcija vrne napako.
func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura float64, minTemp, maxTemp float64) error {
	if _, ok := postaje[oznaka]; !ok {
		return fmt.Errorf("Postaja z oznako '%s' ne obstaja", oznaka)
	}

	if temperatura < minTemp || temperatura > maxTemp {
		return fmt.Errorf("Neveljavna meritev: %f", temperatura)
	}

	p := postaje[oznaka]
	p.meritve = append(p.meritve, temperatura)
	postaje[oznaka] = p

	return nil
}

// statistika je interna funkcija za izračun najnižje, najvišje in povprečne temperature.
func statistika(postaje map[string]Postaja, oznaka string, stMeritev int) (najnizja, najvisja, povprecna float64, ok bool) {
	if _, ok := postaje[oznaka]; !ok || len(postaje[oznaka].meritve) < stMeritev {
		return 0.0, 0.0, 0.0, false
	}

	min_v := postaje[oznaka].meritve[0]
	max_v, sum_v := min_v, 0.0

	for _, v := range postaje[oznaka].meritve {
		if v < min_v {
			min_v = v
		} else if v > max_v {
			max_v = v
		}

		sum_v += v
	}

	return min_v, max_v, sum_v / float64(len(postaje[oznaka].meritve)), true
}

// IzpisMeritev izpiše meritve vseh postaj.
func IzpisMeritev(postaje map[string]Postaja) {
	fmt.Println("MERITVE:")

	for k, v := range postaje {
		fmt.Printf("%s - %s (%s): %v\n", k, v.ime, v.kraj, v.meritve)
	}
}

// IzpisPorocila izpiše statistiko in opis vseh postaj.
func IzpisPorocila(postaje map[string]Postaja, stMeritev int) {
	for k, v := range postaje {
		min_v, max_v, avg_v, ok := statistika(postaje, k, stMeritev)

		if ok {
			opis := ""

			switch {
			case avg_v >= 25:
				opis = "Vroče"
			case avg_v >= 15:
				opis = "Toplo"
			case avg_v >= 5:
				opis = "Hladno"
			default:
				opis = "Mrzlo"
			}

			fmt.Printf(
				"%s: min %.1f °C, max %.1f °C, povprečje %.1f °C -> %s\n",
				v.ime, min_v, max_v, avg_v, opis,
			)
		} else {
			fmt.Printf("%s: premalo meritev\n", v.ime)
		}
	}
}
