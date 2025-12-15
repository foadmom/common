package logger

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// // Configuration for logging
// type Config struct {
// 	// Enable console logging
// 	ConsoleLoggingEnabled bool

// 	// EncodeLogsAsJson makes the log framework log JSON
// 	EncodeLogsAsJson bool
// 	// FileLoggingEnabled makes the framework log to a file
// 	// the fields below can be skipped if this value is false!
// 	FileLoggingEnabled bool
// 	// Directory to log to to when filelogging is enabled
// 	Directory string
// 	// Filename is the name of the logfile which will be placed inside the directory
// 	Filename string
// 	// MaxSize the max size in MB of the logfile before it's rolled
// 	MaxSize int
// 	// MaxBackups the max number of rolled files to keep
// 	MaxBackups int
// 	// MaxAge the max age in days to keep a logfile
// 	MaxAge int
// }

type Zer0Lgger struct {
}

var __Logger *Zer0Lgger = &Zer0Lgger{}

func GetInstance() *Zer0Lgger {
	// var _instance CLogger = __Logger
	// return _instance.(*Zer0Lgger)
	return __Logger
}

var _logger zerolog.Logger

// Configure sets up the logging framework
// In production, the container logs will be collected and file logging should be disabled. However,
// during development it's nicer to see logs as text and optionally write to a file when debugging
// problems in the containerized pipeline
//
// The output log file will be located at /var/log/service-xyz/service-xyz.log and
// will be rolled according to configuration set.
func (lOgger *Zer0Lgger) Configure(config Config) {
	var writers []io.Writer

	if config.ConsoleLoggingEnabled {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	}
	if config.FileLoggingEnabled {
		writers = append(writers, newRollingFile(config))
	}
	mw := io.MultiWriter(writers...)
	// mw.timestampFormat = "2006-01-02 15:04:05.000"
	_logger = zerolog.New(mw).With().Caller().Logger()
	// _logger := zerolog.New(mw).With().Timestamp().Logger()

	_logger.Info().
		Bool("fileLogging", config.FileLoggingEnabled).
		Bool("jsonLogOutput", config.EncodeLogsAsJson).
		Str("logDirectory", config.Directory).
		Str("fileName", config.Filename).
		Int("maxSizeMB", config.MaxSize).
		Int("maxBackups", config.MaxBackups).
		Int("maxAgeInDays", config.MaxAge).
		Msg("logging configured")

}

func (lOgger *Zer0Lgger) Printf(level int, format string, v ...interface{}) {
	if level <= LogLevel {
		var _msg string = fmt.Sprintf(format, v...)
		lOgger.Print(level, _msg)
	}

}

func (lOgger *Zer0Lgger) Print(level int, msg string) {
	if level <= LogLevel {
		var _timeStamp string = time.Now().Format("2006-01-02T15:04:05.000")
		var _msg string = _timeStamp + "> " + msg
		switch level {
		case Trace:
			_logger.Trace().Msg(_msg)
		case Debug:
			_logger.Debug().Msg(_msg)
		case Info:
			_logger.Info().Msg(_msg)
		case Warning:
			_logger.Warn().Msg(_msg)
		case Error:
			_logger.Error().Msg(_msg)
		case Fatal:
			_logger.Fatal().Msg(_msg)
		case Panic:
			_logger.Panic().Msg(_msg)
		default:
			_logger.Info().Msg(_msg)
		}
	}
}

// ============================================================================
func newRollingFile(config Config) io.Writer {
	var _err error
	var _iow io.Writer
	if _err = os.MkdirAll(config.Directory, 0744); _err != nil {
		log.Error().Err(_err).Str("path", config.Directory).Msg("can't create log directory")
		return nil
	} else {
		_file := config.Directory + config.Filename
		_iow, _err = os.OpenFile(_file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	}
	return _iow
}

// func Init(serviceName string) {
// 	// f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
// 	// if err != nil {
// 	// 	log.Fatalf("error opening file: %v", err)
// 	// }

// 	// log.SetOutput(f)
// 	// log.SetFlags(log.Lshortfile)

// 	var _configuration Config = Config{true, false, true, "./", serviceName + ".log", 0, 14, 1}
// 	Configure(_configuration)
// }

// func Instance() *zerolog.Logger {
// 	return __ZeroLogger
// }
