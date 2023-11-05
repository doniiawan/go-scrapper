package main

import (
	"flag"
	"log"
)

func main() {
	log.Println("Starting cron..")

	var feature string
	flag.StringVar(&feature, "f", "", "feature that wants to be run")
	flag.Parse()

	if feature == "" {
		return
	}

	switch feature {
	case "scrapper-tokopedia":
		runApp()
	default:
		return
	}
}
