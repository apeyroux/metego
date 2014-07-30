package main

import (
	"flag"
	"fmt"
	metego "github.com/j4/metego"
	"log"
)

var city string // arg ligne commande

func init() {
	flag.StringVar(&city, "city", "", "ville de recherche")
}

func main() {
	flag.Parse()

	if city == "" {
		flag.Usage()

	} else {
		nominatim, err := metego.NewNominatim(city)

		if err != nil {
			log.Fatal("Bad ... %s", err)
		}

		omet, err := metego.NewOMet(nominatim)

		if err != nil {
			log.Fatal("Bad ... %s", err)
		}

		fmt.Printf("--- %s ---\n", nominatim.Display_name)
		fmt.Printf("Status OMte:%s\n", omet.Status)
		fmt.Printf("Pression atmosphérique:%d\n", omet.Pressure[0])
		fmt.Printf("---\n")
	}
}
