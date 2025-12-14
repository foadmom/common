package logger

// Configuration for logging
type Config struct {
	// Enable console logging
	ConsoleLoggingEnabled bool

	// EncodeLogsAsJson makes the log framework log JSON
	EncodeLogsAsJson bool
	// FileLoggingEnabled makes the framework log to a file
	// the fields below can be skipped if this value is false!
	FileLoggingEnabled bool
	// Directory to log to to when filelogging is enabled
	Directory string
	// Filename is the name of the logfile which will be placed inside the directory
	Filename string
	// MaxSize the max size in MB of the logfile before it's rolled
	MaxSize int
	// MaxBackups the max number of rolled files to keep
	MaxBackups int
	// MaxAge the max age in days to keep a logfile
	MaxAge int
}

const (
	Off = iota
	// LevelPanicValue is the value used for the panic level field.
	Panic
	// LevelFatalValue is the value used for the fatal level field.
	Fatal
	// LevelErrorValue is the value used for the error level field.
	Error
	// LevelWarnValue is the value used for the warn level field.
	Warning
	// LevelInfoValue is the value used for the info level field.
	Info
	// LevelDebugValue is the value used for the debug level field.
	Debug
	// LevelTraceValue is the value used for the trace level field.
	Trace
)

type Logger struct {
}

type CLogger interface {
	// Instance() *Logger

	Configure(config Config)
	Printf(level int, format string, v ...interface{})
	Print(level int, msg string)
}

// var ActualLogger *CLogger

var LogLevel int = Debug

func SetLogLevel(level int) {
	LogLevel = level
}

// Configure sets up the logging framework
//
// In production, the container logs will be collected and file logging should be disabled. However,
// during development it's nicer to see logs as text and optionally write to a file when debugging
// problems in the containerized pipeline
//
// The output log file will be located at /var/log/service-xyz/service-xyz.log and
// will be rolled according to configuration set.
// func Instance() *CLogger {
// 	if ActualLogger == nil {
// 		// _Logger = &Zer0Lgger{}
// 		ActualLogger = getInstance()
// 	}
// 	return ActualLogger
// }

// func (logger *Logger) Configure(config Config) {
// 	logger.Configure(config)
// }

// func (logger *Logger) Printf(level int, format string, v ...interface{}) {
// 	logger.Printf(level, format, v...)
// }

// func (logger *Logger) Print(level int, msg string) {
// 	logger.Printf(level, msg)
// }
