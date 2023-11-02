package dataprocessing

import (
	"github.com/baptistemehat/go-leadsheet/core/datamodel/song"
	"github.com/baptistemehat/go-leadsheet/core/dataprocessing/formatting/latexformatting"
	"github.com/baptistemehat/go-leadsheet/core/dataprocessing/parsing/inlinechordparsing"
	"github.com/baptistemehat/go-leadsheet/core/pipeline"
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

type Configuration struct {
	Input   string
	Outputs []string
}

type LeadsheetProcessor struct {
	status   Status
	pipeline pipeline.Pipeline[string, string]
}

func NewLeadsheetProcessor() *LeadsheetProcessor {
	return &LeadsheetProcessor{
		status: StatusNotStarted,
	}
}

func (l *LeadsheetProcessor) Configure(config Configuration) error {

	// Create Stages from configuration
	// TODO have a ParserFactory that creates the appropriate parser
	leadsheetParser := inlinechordparsing.InlineChordParser{}
	parsingStage := pipeline.NewProcessingStage[string, song.Song](leadsheetParser.Parse)

	// dataSaverStage := pipeline.NewSinkStage[string]()

	songFormatter := latexformatting.LatexSongFormatter{}
	formattingStage := pipeline.NewProcessingStage[song.Song, string](songFormatter.FormatSong)

	// pdfGenerator := pdfgenerator.NewPdfGenerator()

	pipeline.Connect[song.Song](parsingStage, formattingStage)
	// pipeline.Connect[song.Song](formattingStage)
	// connect pdfGenerqtors to pipeline outputs

	return nil
}

func (l *LeadsheetProcessor) GeneratePdf(leadsheet string) error {
	l.pipeline.Input <- leadsheet
	return nil
}

// Status returns pdf generation status
func (l *LeadsheetProcessor) Status() Status {
	return l.status
}

// TODO thisis tempo. Will only work for 1 output
// Output returns path to the generated pdf file
func (l *LeadsheetProcessor) Output() string {
	return ""
}
