package main

import (
	"github.com/oleggator/tnt-generator/generator"
	"github.com/oleggator/tnt-generator/parser"
	"log"
	"os"
)

func main() {
	// open sample avro scheme
	file, err := os.Open("./schema.avsc")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	cStructs, err := parser.Parse(file)
	if err != nil {
		log.Fatalln(err)
	}

	// generate models header
	if err := generator.GenerateModelsH("generated/models.h", cStructs); err != nil {
		log.Fatalln(err)
	}

	// generate models implementation
	if err := generator.GenerateModelsC("generated/models.c", cStructs); err != nil {
		log.Fatalln(err)
	}

	if err := generator.Format("generated/models.h", "generated/models.c"); err != nil {
		log.Fatalln(err)
	}
}
