package pipeline

import "fmt"

type ProcessingStage[I any, O any] struct {
    Sink[I]
	process func(I) (O, error)
    Source[O]
}

func NewProcessingStage[I any, O any](process func(I) (O, error)) *ProcessingStage[I, O] {
	return &ProcessingStage[I, O]{
        Sink: NewSink[I](),
		process: process,
        Source: NewSource[O](),
	}
}

func (processingStage *ProcessingStage[I, O]) Run() error {

	if processingStage.inputChannel == nil {
		return fmt.Errorf("No input channel was set.")
	}

    if len(processingStage.outputChannels) == 0 {
        return fmt.Errorf("No output channel was created")
    }

	go func() {

		stop := false

		for !stop {
			inputData := <-*processingStage.inputChannel

			outputData, err := processingStage.process(inputData)
			if err != nil {
				println("Error")
			}

			processingStage.Send(outputData)

			stop = true
		}
	}()

	return nil
}
