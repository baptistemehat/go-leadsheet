package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var Logger zerolog.Logger

func Init() {
	zerolog.SetGlobalLevel(zerolog.TraceLevel)
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	Logger = zerolog.New(output).With().Timestamp().Logger()
}

// TODO : change colors of log

// func loggerMain() {
// 	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
// 	output.FormatLevel = func(i interface{}) string {
// 		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
// 	}
// 	output.FormatMessage = func(i interface{}) string {
// 		return fmt.Sprintf("***%s****", i)
// 	}
// 	output.FormatFieldName = func(i interface{}) string {
// 		return fmt.Sprintf("%s:", i)
// 	}
// 	output.FormatFieldValue = func(i interface{}) string {
// 		return strings.ToUpper(fmt.Sprintf("%s", i))
// 	}

// 	log := zerolog.New(output).With().Timestamp().Logger()

// 	log.Debug().Str("foo", "bar").Msg("Hello world")
// 	log.Info().Str("foo", "bar").Msg("Hello world")
// 	log.Warn().Str("foo", "bar").Msg("Hello world")
// 	log.Error().Str("foo", "bar").Msg("Hello world")
// 	log.Fatal().Str("foo", "bar").Msg("Hello world")
// 	log.Panic().Str("foo", "bar").Msg("Hello world")

// }
