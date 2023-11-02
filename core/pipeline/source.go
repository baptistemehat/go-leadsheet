package pipeline

import "fmt"

type Output[O any] interface {
    NewOutput() (*chan O, error)
    Send(O)
}
// Source
type Source[T any] struct {
	outputChannels []chan T
}

func (source Source[T]) NewOutput() (*chan T, error) {
	source.outputChannels = append(source.outputChannels, make(chan T, 2))
	return &source.outputChannels[len(source.outputChannels)], nil
}

func (source Source[T]) Send(data T) {
	for _, outputChannel := range source.outputChannels {
		outputChannel <- data
	}
}

func NewSource[O any]() Source[O] {
	return Source[O]{
		outputChannels: []chan O{},
	}
}

type SourceStage[O any] struct {
	Source[O]
	process func() (O, error)
}

func NewSourceStage[O any](process func() (O, error)) *SourceStage[O] {
	return &SourceStage[O]{
		Source:  NewSource[O](),
		process: process,
	}
}

func (sourceStage *SourceStage[O]) Run() error {

	if len(sourceStage.outputChannels) == 0 {
		return fmt.Errorf("No output channel was created")
	}

	go func() {

		stop := false

		for !stop {
			outputData, err := sourceStage.process()
			if err != nil {
				println("Error")
			}

			sourceStage.Send(outputData)

			stop = true
		}
	}()

	return nil
}
