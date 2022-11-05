package main

import (
	"flag"
	"fmt"
	"os/exec"
)

var (
	source *string = flag.String("source", "", "Path to source raw text file")
)

func generatePdf(src *string) error {

	cmd := exec.Command("./generate-pdf.sh", *src)

	out, err := cmd.Output()

	if err != nil {
		fmt.Println(string(out))
		return err
	}

	return nil
}

func main() {

	flag.Parse()

	// no source file provided
	if len(*source) < 1 {
		flag.CommandLine.Usage()
		fmt.Println("path to source file is empty")
	}

	if err := generatePdf(source); err != nil {
		fmt.Println(err)
	}

}
