package main

import (
	"flag"
	"fmt"
	"github.com/oleggator/tnt-generator/generator"
	"github.com/oleggator/tnt-generator/parser"
	"log"
	"os"
)

func main() {
	inputFile := flag.String("i", "schema.avsc", "avro schema")
	outputDir := flag.String("o", "generated", "output directory")
	flag.Parse()

	// open sample avro scheme
	file, err := os.Open(*inputFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	cStructs, err := parser.Parse(file)
	if err != nil {
		log.Fatalln(err)
	}

	// generate models header
	modelsHPath := fmt.Sprintf("%s/models.h", *outputDir)
	if err := generator.GenerateModelsH(modelsHPath, cStructs); err != nil {
		log.Fatalln(err)
	}

	// generate models implementation
	modelsCPath := fmt.Sprintf("%s/models.c", *outputDir)
	if err := generator.GenerateModelsC(modelsCPath, cStructs); err != nil {
		log.Fatalln(err)
	}

	if err := generator.Format(modelsHPath, modelsCPath); err != nil {
		log.Fatalln(err)
	}
}
