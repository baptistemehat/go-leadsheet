package song

import (
	"fmt"
	"strings"
)

type SongProperties struct {
	Title    string `json:"title"`
	Composer string `json:"composer"`
	Capo     int    `json:"capo"`
	Key      string `json:"key"`
}

func (sp *SongProperties) Format(f SongFormatter) (string, error) {
	return f.FormatSongProperties(sp)
}

func (sp *SongProperties) SetProperty(name string, value interface{}) error {
	nameToLower := strings.ToLower(name)

	switch nameToLower {
	case "title":
		if title, ok := value.(string); !ok {
			return fmt.Errorf("value '%s' is not suitable for title", value)
		} else {
			sp.Title = title
		}

	case "composer":
		if composer, ok := value.(string); !ok {
			return fmt.Errorf("value '%s' is not suitable for composer", value)
		} else {
			sp.Composer = composer
		}

	case "capo":
		if capo, ok := value.(int); !ok {
			return fmt.Errorf("value '%s' is not suitable for capo", value)
		} else {
			sp.Capo = capo
		}
	case "key":
		if key, ok := value.(string); !ok {
			return fmt.Errorf("value '%s' is not suitable for title", value)
		} else {
			sp.Key = key
		}

	default:
		return fmt.Errorf("property name not found: '%s'", name)
	}
	return nil
}

type Song struct {
	Properties SongProperties `json:"properties"`
	Sections   []Section      `json:"sections"`
}

func NewSong() Song {
	return Song{
		Properties: SongProperties{},
		Sections:   []Section{},
	}
}

func (song *Song) AddSection(section Section) {
	song.Sections = append(song.Sections, section)
}

func (song *Song) SetProperties(sp SongProperties) {
	song.Properties = sp
}

func (song *Song) Format(f SongFormatter) (string, error) {
	return f.FormatSong(song)
}
