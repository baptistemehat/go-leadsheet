package pdfGenerator

import (
	"testing"

	"github.com/baptistemehat/go-leadsheet/core/song/formatter"
	"github.com/baptistemehat/go-leadsheet/core/song/parser"
)

func TestGeneratePdfFromBuffer(t *testing.T) {
	builder := Builder{
		Parser:    parser.InlineChordParser{},
		Formatter: &formatter.LatexSongFormatter{},
	}

	// create pdf generator
	pdfGenerator, err := NewPdfGenerator(builder)
	if err != nil {
		t.Errorf("%s", err)
	}

	err = pdfGenerator.GeneratePdfFromBuffer(`

Title: Hotel California
Composer: Eagles
Capo: 0
Key: Bm

{Verse}

[Am] On a dark desert highway,[E7] cool wind in my hair
[G] Warm smell of colitas [D] rising up through the air
[F] Up ahead in the distance,[C] I saw a shimmering light
[Dm] My head grew heavy and my sight grew dim,[E7] I had to stop for the night
{Chorus}

[F] Welcome to the Hotel Califo[C]rnia.
Such a [E7]lovely place, (such a lovely place), such a [Am]lovely face
[F]Plenty of room at the Hotel Cali[C]fornia
Any [Dm]time of year, (any time of year) You can [E7]find it here

	`)

	if err != nil {
		t.Errorf("%s", err)
	}
}
