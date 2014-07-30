package main

import (
	"flag"
	"fmt"
	"metego"
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
		nominatim := NewNominatim(city)
		omet := NewOMET(nominatim)
		fmt.Printf("--- %s ---\n", nominatim.Display_name)
		fmt.Printf("Status OMte:%s\n", omet.Status)
		fmt.Printf("Pression atmosphérique:%d\n", omet.Pressure[0])
		fmt.Printf("---\n")
	}
}
