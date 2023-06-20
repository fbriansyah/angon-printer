package printing

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

type LinuxPrinting struct{}

func (wp *LinuxPrinting) Print(texts []string) {
	printText := ""

	header, footer := setHeaderFooter("501")

	for _, text := range texts {
		printText += fmt.Sprintf("%s %s %s", string(header), text, string(footer))
	}

	argstr := []string{"-c", fmt.Sprintf("echo -e '%s' | lpr", printText)}
	_, err := exec.Command("/bin/bash", argstr...).Output()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
