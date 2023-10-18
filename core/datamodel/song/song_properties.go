package song

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/baptistemehat/go-leadsheet/core/common/logger"
)

// TODO : replace by an actual key-value pair ? (string, string)
type SongProperties struct {
	Title    string `json:"title"`
	Composer string `json:"composer"`
	Capo     int    `json:"capo"`
	Key      string `json:"key"`
}

// Formatter
type SongPropertiesFormatter interface {
	FormatSongProperties(*SongProperties) (string, error)
}

// Format
func (sp *SongProperties) Format(f SongPropertiesFormatter) (string, error) {
	return f.FormatSongProperties(sp)
}

// SetProperty
// temporary implementation
func (sp *SongProperties) SetProperty(name string, value interface{}) error {
	nameToLower := strings.ToLower(name)

	logger.Logger.Debug().Msgf("received new token: %s, %s", name, value)

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

		capoString, ok := value.(string)
		if !ok {
			return fmt.Errorf("value '%s' is not suitable for capo", value)
		}

		if capo, err := strconv.Atoi(capoString); err != nil {
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
