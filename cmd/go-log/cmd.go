package golog

import (
	"errors"
	"io"
	"log"
	"os"
)

const INFO_PREFIX = "INFO: "
const WARNING_PREFIX = "WARNING: "
const ERROR_PREFIX = "ERROR: "

const LOG_FILE = "log.log"

type Logger struct {
	Info   *log.Logger
	Warnig *log.Logger
	Error  *log.Logger
}

type Config struct {
	InfoOperation   Operation
	WarnigOperation Operation
	ErrorOperation  Operation
}

type Operation struct {
	Prefix string
	Flag   int
}

func GetDefaultInfo() Operation {
	return Operation{
		INFO_PREFIX,
		log.Ldate | log.Ltime,
	}
}

func GetDefaultWarning() Operation {
	return Operation{
		WARNING_PREFIX,
		log.Ldate | log.Ltime | log.Lshortfile,
	}
}

func GetDefaultError() Operation {
	return Operation{
		ERROR_PREFIX,
		log.Ldate | log.Ltime | log.Lshortfile,
	}
}

func GetDefaultConfig() Config {
	return Config{
		GetDefaultInfo(),
		GetDefaultWarning(),
		GetDefaultError(),
	}
}

func Init(out io.Writer, config *Config) (*Logger, error) {
	if config == nil {
		return nil, errors.New("configuration is required")
	}

	return &Logger{
		log.New(out, config.InfoOperation.Prefix, config.InfoOperation.Flag),
		log.New(out, config.WarnigOperation.Prefix, config.WarnigOperation.Flag),
		log.New(out, config.ErrorOperation.Prefix, config.ErrorOperation.Flag),
	}, nil

}

func DefaultConsole() (*Logger, error) {
	defaultCfg := GetDefaultConfig()
	return Init(GetConsoleWritter(), &defaultCfg)
}

func DefaultFile() (*Logger, error) {
	defaultCfg := GetDefaultConfig()

	outWritter, err := GetFileWritter(LOG_FILE)
	if err != nil {
		return nil, err
	}

	return Init(outWritter, &defaultCfg)
}

func GetConsoleWritter() io.Writer {
	return os.Stderr
}

func GetFileWritter(fileName string) (io.Writer, error) {
	return os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
}
