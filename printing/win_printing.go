package printing

import (
	"fmt"
	"log"

	"github.com/alexbrainman/printer"
)

type WinPrinting struct{}

func (wp *WinPrinting) Print(texts []string) {
	name, err := printer.Default()
	if err != nil {
		log.Fatalln(err)
	}

	p, err := printer.Open(name)
	if err != nil {
		log.Fatalln(err)
	}
	defer p.Close()

	err = p.StartDocument("test page", "RAW")
	if err != nil {
		log.Fatalf("StartDocument failed: %v", err)
	}
	defer p.EndDocument()

	err = p.StartPage()
	if err != nil {
		log.Fatalf("StartPage failed: %v", err)
	}

	header, footer := setHeaderFooter("501")
	// fmt.Println(header, footer)

	for _, text := range texts {
		fmt.Fprintf(p, "%s %s %s", string(header), text, string(footer))
	}

	err = p.EndPage()
	if err != nil {
		log.Fatalf("EndPage failed: %v", err)
	}
}
