package pdfGenerator

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/baptistemehat/go-leadsheet/core/song/parser"
	songFormatter "github.com/baptistemehat/go-leadsheet/core/song/songFormatters"
)

// Generation status
type Status uint8

const (
	StatusNotStarted Status = iota
	StatusInProgress Status = iota
	StatusDone       Status = iota
	StatusError      Status = iota
)

func (s Status) String() string {
	switch s {
	case StatusNotStarted:
		return "not started"
	case StatusInProgress:
		return "in progress"
	case StatusDone:
		return "done"
	case StatusError:
		return "error"
	default:
		return ""
	}
}

// PdfGenerator
type PdfGenerator struct {
	status     Status
	outputFile string
}

func NewPdfGenerator() (*PdfGenerator, error) {
	p := &PdfGenerator{
		status:     StatusNotStarted,
		outputFile: "latex/tmp/out/main.pdf",
	}
	return p, nil
}

func (pg *PdfGenerator) GeneratePdfFromBufferAlt(buffer string) error {
	texFile := "latex/tmp/songs/leadsheet.tex"

	p := parser.Parser{}

	song, _ := p.Parse(buffer)

	f := songFormatter.LatexSongFormatter{}
	actualString, _ := f.FormatSong(&song)

	file, err := os.Create(texFile)
	if err != nil {
		return err
	}

	defer os.Remove(texFile)

	file.WriteString(actualString)
	file.Close()

	err = pg.Tex2pdf()
	if err != nil {
		pg.status = StatusError
		return err
	}

	pg.status = StatusDone
	return nil
}

// GeneratePdfFromBuffer generates a pdf leadsheet from a string buffer
func (pg *PdfGenerator) GeneratePdfFromBuffer(buffer string) error {

	sourceFile := "latex/tmp/leadsheet.txt.tmp"

	f, err := os.Create(sourceFile)
	if err != nil {
		return err
	}

	defer os.Remove(sourceFile)

	f.WriteString(buffer)
	f.Close()

	return pg.GeneratePdf(sourceFile)
}

// GeneratePdf generates a pdf song file from a raw text source file
func (pg *PdfGenerator) GeneratePdf(source string) error {

	pg.status = StatusInProgress

	err := pg.txt2tex(source)
	if err != nil {
		pg.status = StatusError
		return err
	}

	err = pg.Tex2pdf()
	if err != nil {
		pg.status = StatusError
		return err
	}

	pg.status = StatusDone
	return nil
}

// Status returns pdf generation status
func (pg PdfGenerator) Status() Status {
	return pg.status
}

// Output returns path to the generated pdf file
func (pg PdfGenerator) Output() string {
	return pg.outputFile
}

// txt2tex "transpiles" input raw text files into leadsheet LateX files
// This function calls script txt2tex.sh
func (pg PdfGenerator) txt2tex(source string) error {
	cmd := exec.Command("./pdfGenerator/txt2tex.sh", source)

	out, err := cmd.Output()

	if err != nil {
		fmt.Println(string(out))
		return err
	}

	return nil
}

// Tex2pdf generates pdf file from the files generated by txt2tex.sh
// This function calls script generate-pdf.sh
func (pg PdfGenerator) Tex2pdf() error {
	// TODO : add output file as argument to have coherence between pdfGenerator.Output and shell scripts
	cmd := exec.Command("./pdfGenerator/generate-pdf.sh")

	out, err := cmd.Output()

	if err != nil {
		fmt.Println(string(out))
		return err
	}

	return nil
}
