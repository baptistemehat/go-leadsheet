package pipeline

type ParserType int8
type FormatterType int8

// TODO move this to dataprocessing 
const (
	InlineChordParser ParserType = iota
	ChordLine
	ChordPro

	LatexFormatter FormatterType = iota
	MarkdownFormatter
	HtmlFormatter
)

type PipelineConfiguration struct {
	Parser    ParserType
	Formatter FormatterType
}

type Stage interface {
	Run() error
}

type Pipeline[I any, O any] struct {
	Input   chan I
	Outputs chan O
}

func NewPipeline[I any, O any](config PipelineConfiguration) (Pipeline[I, O], error) {

	return Pipeline[I, O]{}, nil
}

func Connect[T any](source Output[T], sink Input[T]) error {

	outputChannel, err := source.NewOutput()
	if err != nil {
		return err
	}

	sink.SetInput(outputChannel)

	return nil
}
