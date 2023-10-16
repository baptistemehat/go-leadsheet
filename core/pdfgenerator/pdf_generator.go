package pdfgenerator

import (
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/baptistemehat/go-leadsheet/core/common/logger"
	"github.com/baptistemehat/go-leadsheet/core/config"
	"github.com/baptistemehat/go-leadsheet/core/datamodel/song"
	"github.com/baptistemehat/go-leadsheet/core/dataprocessing/parsing"
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

// Builder stores parser and formatter to use for text processing
type Builder struct {
	Parser        parsing.LeadsheetParser
	SongFormatter song.SongFormatter
}

// PdfGenerator
type PdfGenerator struct {
	status     Status
	outputFile string
	builder    Builder
	config     config.Configuration
}

// NewPdfGenerator create a new PdfGenerator
func NewPdfGenerator(builder Builder, config config.Configuration) (*PdfGenerator, error) {
	p := &PdfGenerator{
		status:     StatusNotStarted,
		outputFile: "latex/tmp/out/main.pdf",
		builder:    builder,
		config:     config,
	}
	return p, nil
}

// Status returns pdf generation status
func (pg PdfGenerator) Status() Status {
	return pg.status
}

// Output returns path to the generated pdf file
func (pg PdfGenerator) Output() string {
	return pg.outputFile
}

// This function calls script generate-pdf.sh
func (pg PdfGenerator) tex2pdf() error {
	// TODO : add output file as argument to have coherence between pdfGenerator.Output and shell scripts
	cmd := exec.Command(pg.config.Script)

	_, err := cmd.Output()

	if err != nil {
		return err
	}

	return nil
}

// WriteStringToFile writes string buffer to a file
func WriteStringToFile(buffer, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString(buffer)

	return nil
}

// GeneratePdfFromBuffer generates pdf from a string buffer
func (pg *PdfGenerator) GeneratePdfFromBuffer(buffer string) error {

	pg.status = StatusInProgress

	// TODO : create folder if it doesn't exist
	// TODO : these folders should be configurable

	// Write input to file
	filename := strconv.FormatInt(time.Now().UnixMilli(), 10) + ".txt"
	if err := WriteStringToFile(buffer, pg.config.Storage+"/"+filename); err != nil {
		pg.status = StatusError
		logger.Logger.Err(err).Msg("Failed to save input to file: ")
		return err
	} else {
		logger.Logger.Debug().Msgf("Song saved to file %s", filename)
	}

	// Parse input
	song, err := pg.builder.Parser.Parse(buffer)
	if err != nil {
		pg.status = StatusError
		logger.Logger.Err(err).Msg("Failed to parse input: ")
		return err
	}

	// Format song
	formattedSong, err := song.Format(pg.builder.SongFormatter)
	if err != nil {
		pg.status = StatusError
		logger.Logger.Err(err).Msg("Failed to format song: ")
		return err
	}

	// Write formatted song
	if err := WriteStringToFile(formattedSong, pg.config.Folder+"/tmp/songs/leadsheet.tex"); err != nil {
		pg.status = StatusError
		logger.Logger.Err(err).Msg("Failed to write formatted song to file: ")
		return err
	}
	//defer os.Remove("latex/tmp/songs/leadsheet.tex")

	// Compile latex
	if err := pg.tex2pdf(); err != nil {
		pg.status = StatusError
		logger.Logger.Err(err).Msg("Failed to compile LateX, tex2pdf: ")
		return err
	}

	pg.status = StatusDone
	return nil
}
