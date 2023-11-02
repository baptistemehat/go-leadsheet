package pipeline

import "fmt"

type Input[I any] interface {
    SetInput(*chan I)
}

// Sink
type Sink[I any] struct {
	inputChannel *chan I
}

func (sink *Sink[I]) SetInput(newInputChannel *chan I) {
	sink.inputChannel = newInputChannel
}

func NewSink[I any]() Sink[I] {
    return Sink[I]{
        inputChannel: nil,
    }
}

type SinkStage[I any] struct {
    Sink[I]
    process func(I) error
}

func NewSinkStage[I any](process func(I) error) *SinkStage[I] {
	return &SinkStage[I]{
		Sink:    NewSink[I](),
		process: process,
	}
}

func (sink *SinkStage[I]) Run() error {

    if sink.inputChannel == nil {
       return fmt.Errorf("No input channel was set.") 
    }

    go func() {

        stop := false 
        
        for !stop {
            inputData := <- *sink.inputChannel

            err := sink.process(inputData) 
            if err != nil {
                println("Error")
            } 

            stop = true
        }
    } ()

    return nil
}
