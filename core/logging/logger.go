package logging

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
)

var f *os.File
var logger zerolog.Logger

func Init() {
	setLogFile()
	go interval()
	logger = zerolog.New(f).With().Timestamp().Logger()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger.Info().Msg("hello world")
	logger.Info().Msg("hello world!")
	log := make(map[string]interface{})
	log["TimeStamp"] = time.Now()
	logger.Info().Interface("Properties", log).Msg("hello world!")
}

func Debug(properties interface{}) {
	logger.Debug().Interface("Properties", properties)
}

func Info(properties interface{}) {
	logger.Info().Interface("Properties", properties).Msg("hello world")
}

func Warning(properties interface{}) {
	logger.Warn().Interface("Properties", properties)
}

func Error(properties interface{}) {
	logger.Error().Interface("Properties", properties)
}

func Fatal(properties interface{}) {
	logger.Fatal().Interface("Properties", properties)
}

func setLogFile() {
	currentTime := time.Now()
	fileName := os.TempDir() + string(os.PathSeparator) + "log-" + currentTime.Format("2006-01-02#15-04-05") + ".json"
	fmt.Println(fileName)
	var err error
	if f, err = os.Create(fileName); err != nil {
		fmt.Println("can't obtain log file", err)
	}
}

func interval() {
	intervalValue := 10 * time.Hour
	for range time.Tick(intervalValue) {
		fmt.Println("interval triggered")
		setLogFile()
	}
}
