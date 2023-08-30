package pdfGenerator

import (
	"testing"

	"github.com/baptistemehat/go-leadsheet/core/config"
	"github.com/baptistemehat/go-leadsheet/core/song/formatter"
	"github.com/baptistemehat/go-leadsheet/core/song/parsing"
)

func TestGeneratePdfFromBuffer(t *testing.T) {
	builder := Builder{
		Parser:    parsing.InlineChordParser{},
		Formatter: &formatter.LatexSongFormatter{},
	}

	path := "/home/baptiste/Programing/projects/github.com/baptistemehat/go-leadsheet/core/config/testResources/config.yaml"
	config, err := config.LoadConfiguration(path)

	if err != nil {
		t.Error(err)
	}

	// create pdf generator
	pdfGenerator, err := NewPdfGenerator(builder, *config)
	if err != nil {
		t.Errorf("%s", err)
	}

	err = pdfGenerator.GeneratePdfFromBuffer(`
Title: Shallow
Composer: Lady Gaga & Bradley Cooper
Capo: 0
Key: G

{Intro}
[Em7] [D/F#] [G] [G]
[C] [C] [G] [D]
[Em7] [D/F#] [G] [G]

{Verse}
Tell me somethin', girl
Are you happy in this modern world?

Or do you need more?
Is there somethin' else you're searchin' for?
	
	`)

	if err != nil {
		t.Errorf("%s", err)
	}
}
