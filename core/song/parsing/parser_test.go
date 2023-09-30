package parsing

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestParseInlineChord(t *testing.T) {

	p := InlineChordParser{}

	song, err := p.Parse(`

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
		t.Fatal(err)
	}

	// Print result
	u, err := json.MarshalIndent(song, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(u))
}

func TestParseeChordLine(t *testing.T) {
	p := InlineChordParser{}

	song, err := p.Parse(`

Title: Hotel California
Composer: Eagles
Capo: 0

[Verse]

Am                        E7
 On a dark desert highway, cool wind in my hair
G                     D
 Warm smell of colitas rising up through the air
F                         C
 Up ahead in the distance, I saw a shimmering light
Dm                                        E7
 My head grew heavy and my sight grew dim, I had to stop for the night


[Chorus]

F                           C
 Welcome to the Hotel California.
       E7                                          Am
Such a lovely place, (such a lovely place), such a lovely face
F                               C
Plenty of room at the Hotel California
    Dm                                       E7
Any time of year, (any time of year) You can find it here

	`)

	if err != nil {
		t.Fatal(err)
	}

	// Print result
	u, err := json.Marshal(song)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(u))
}
