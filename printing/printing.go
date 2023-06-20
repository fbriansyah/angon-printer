package printing

import "runtime"

type AgnPrinter interface {
	Print([]string)
}

func setHeaderFooter(tipeKertas string) (header []byte, footer []byte) {

	lineCount := 22

	switch tipeKertas {
	case "501":
		lineCount = 22
	case "502":
		lineCount = 29 // 1/3 KUARTO
	case "503":
		lineCount = 44 // 1/2 KUARTO
	case "504":
		lineCount = 32 // 8,1 INCH X 4 INCH
	default:
		lineCount = 22
	}
	// header := []byte{}

	header = append(header, byte(27), byte(103))
	header = append(header, byte(27), byte(108), byte(1))
	header = append(header, byte(27), byte(81), byte(1))
	header = append(header, byte(27), byte(78), byte(1))
	header = append(header, byte(27), byte('!'), byte(4))
	header = append(header, byte(27), byte('0'))
	header = append(header, byte(27), byte('C'), byte(lineCount))

	footer = []byte{12}

	return
}

func NewPrinter() AgnPrinter {
	if runtime.GOOS == "windows" {
		return &WinPrinting{}
	} else {
		return &LinuxPrinting{}
	}
}
