package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

// TODO : add a log output in json

var Logger zerolog.Logger

func Init() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	Logger = zerolog.New(output).With().Timestamp().Logger()
}

// TODO : change colors of log

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
