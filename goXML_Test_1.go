package main

import (
	"flag"

	"github.com/grokify/go-wordpressxml"
)

// Function to create the executable for HTML export
func main() {
	inFileName := flag.String("infile", "", "The name of the XML file to convert")
	outFileName := flag.String("outfile", "", "The name of the HTML file to export to")
	eltitle := flag.String("title", "", "A title for the export file (optional)")

	flag.Parse()

	wp := wordpressxml.NewWordPressXML()

	if *inFileName == "" {
		panic("an input XML file must be specified")
	}

	if *outFileName == "" {
		panic("an output HTML file must be specified")
	}

	err := wp.ReadFile(*inFileName)
	if err != nil {
		panic(err)
	}

	wp.ItemsToHTML(*outFileName, *eltitle)
}
