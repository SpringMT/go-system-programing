package main

import (
	"encoding/csv"
	"log"
	"os"
)

var records = [][]string{
	{"first_name", "last_name", "username"},
	{"Rob", "Pike", "rob"},
	{"Ken", "Thompson", "ken"},
	{"Robert", "Griesemer", "gri"},
}

func main() {

	file, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}
	writer := csv.NewWriter(file)
	for _, record := range records {
		if err = writer.Write(record); err != nil {
			panic(err)
		}
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		log.Fatal(err)
	}
	file.Close()
}
