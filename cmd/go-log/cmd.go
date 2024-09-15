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

const DEFAULT_LOG_FILE = "log.log"

type Logger struct {
	Info   *log.Logger
	Warnig *log.Logger
	Error  *log.Logger
}

type Config struct {
	OutWriter       io.Writer
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

func Init(config *Config) (*Logger, error) {
	if config == nil {
		return nil, errors.New("configuration is required")
	}

	return &Logger{
		log.New(config.OutWriter, config.InfoOperation.Prefix, config.InfoOperation.Flag),
		log.New(config.OutWriter, config.WarnigOperation.Prefix, config.WarnigOperation.Flag),
		log.New(config.OutWriter, config.ErrorOperation.Prefix, config.ErrorOperation.Flag),
	}, nil

}

func DefaultConsole() (*Logger, error) {
	defaultCfg := Config{
		GetConsoleWriter(),
		GetDefaultInfo(),
		GetDefaultWarning(),
		GetDefaultError(),
	}

	return Init(&defaultCfg)
}

func GetConsoleWriter() io.Writer {
	return os.Stderr
}

func DefaultFile() (*Logger, error) {
	outWriter, err := GetFileWriter(DEFAULT_LOG_FILE)

	if err != nil {
		return nil, err
	}

	defaultCfg := Config{
		outWriter,
		GetDefaultInfo(),
		GetDefaultWarning(),
		GetDefaultError(),
	}

	return Init(&defaultCfg)
}

func GetFileWriter(fileName string) (io.Writer, error) {
	return os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
}
