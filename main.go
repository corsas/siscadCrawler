package main

import (
	"github.com/joho/godotenv"
	"siscadCrawler/scrapper"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		return
	}

	scrapper.Run()
}
