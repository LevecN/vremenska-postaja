package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/LevecN/vremenska-postaja/vreme"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "vremenska-postaja",
		Usage: "Preprost program za delo z vremenskimi postajami",
		Flags: []cli.Flag{
			&cli.FloatFlag{
				Name:  "minTemp",
				Usage: "Najnižja dovoljena temperatura meritve",
				Value: -50,
			},
			&cli.FloatFlag{
				Name:  "maxTemp",
				Usage: "Najvišja dovoljena temperatura meritve",
				Value: 50,
			},
			&cli.IntFlag{
				Name:  "stMeritev",
				Usage: "Najmanjše število meritev, potrebnih za izračun statistike",
				Value: 3,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			minTemp := cmd.Float("minTemp")
			maxTemp := cmd.Float("maxTemp")
			stMeritev := cmd.Int("stMeritev")
			return vremenskaPostaja(minTemp, maxTemp, stMeritev)
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func vremenskaPostaja(minTemp, maxTemp float64, stMeritev int) error {
	var postaje map[string]vreme.Postaja = make(map[string]vreme.Postaja)

	vreme.DodajPostajo(postaje, "KP01", "Koper Luka", "Koper")
	vreme.DodajPostajo(postaje, "KR01", "Kredarica", "Triglav")
	vreme.DodajPostajo(postaje, "LJ01", "Ljubljana Bežigrad", "Ljubljana")
	vreme.DodajPostajo(postaje, "MB01", "Maribor Tabor", "Maribor")

	for _, v := range []float64{24.5, 27, 26.1, 28.4} {
		if err := vreme.DodajMeritev(postaje, "KP01", v, minTemp, maxTemp); err != nil {
			fmt.Println(err)
		}
	}

	for _, v := range []float64{2, -1.5, 3.5} {
		if err := vreme.DodajMeritev(postaje, "KR01", v, minTemp, maxTemp); err != nil {
			fmt.Println(err)
		}
	}

	for _, v := range []float64{18.2, 21.5, 19.8, 16.5} {
		if err := vreme.DodajMeritev(postaje, "LJ01", v, minTemp, maxTemp); err != nil {
			fmt.Println(err)
		}
	}

	for _, v := range []float64{12.4} {
		if err := vreme.DodajMeritev(postaje, "MB01", v, minTemp, maxTemp); err != nil {
			fmt.Println(err)
		}
	}

	vreme.IzpisMeritev(postaje)
	fmt.Println()

	if err := vreme.DodajMeritev(postaje, "KR02", -10.2, minTemp, maxTemp); err != nil {
		fmt.Println(err)
	}
	fmt.Println()

	if err := vreme.DodajMeritev(postaje, "KR01", -50.50, minTemp, maxTemp); err != nil {
		fmt.Println(err)
	}
	fmt.Println()

	if err := vreme.DodajMeritev(postaje, "KR01", 15, minTemp, maxTemp); err != nil {
		fmt.Println(err)
	}

	vreme.IzpisMeritev(postaje)
	fmt.Println()

	vreme.IzpisPorocila(postaje, stMeritev)

	return nil
}
