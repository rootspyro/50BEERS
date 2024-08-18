package log

import (
	"fmt"
	"time"

	"github.com/rootspyro/50BEERS/config"
)

type logTypes struct {
	Info    string
	Debug   string
	Warning string
	Error   string
}

var LogTypes logTypes = logTypes{
	Info:    "INFO",
	Debug:   "DEBUG",
	Warning: "WARNING",
	Error:   "ERROR",
}

func Info(data string) {
	message(data, LogTypes.Info)
}

func Debug(data string) {
	message(data, LogTypes.Debug)
}

func Warning(data string) {
	message(data, LogTypes.Warning)
}

func Error(data string) {
	message(data, LogTypes.Error)
}

func message(data string, logType string) {
	typeMessage := ""

	switch logType {
	case  LogTypes.Info:
		typeMessage = config.Colors.Blue + LogTypes.Info + config.Colors.Reset
		break

	case LogTypes.Debug:
		typeMessage = config.Colors.Green + LogTypes.Debug + config.Colors.Reset
		break

	case LogTypes.Warning:
		typeMessage = config.Colors.Yellow + LogTypes.Warning + config.Colors.Reset
		break

	case LogTypes.Error:
		typeMessage = config.Colors.Red + LogTypes.Error + config.Colors.Reset
		break

	default:
		typeMessage = "LOG"
		break
	}

	fmt.Printf(
		"\n %s | %s | %s\n",
		time.Now().Local(),
		typeMessage,
		data,
	)
}
